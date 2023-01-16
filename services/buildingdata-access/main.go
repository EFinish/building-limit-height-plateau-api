package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	protoOut "github.com/EFinish/building-limit-height-plateau-api/proto/gen/buildingdataaccess/v1"
	"github.com/EFinish/building-limit-height-plateau-api/services/buildingdata-access/config"
	logger "github.com/EFinish/building-limit-height-plateau-api/utilities/logger"
)

type (
	buildingdataAccess struct {
		conf                         *config.Config
		logger                       logger.CustomLogger
		buildingLimitCollection      *mongo.Collection
		heightPlateauCollection      *mongo.Collection
		splitBuildingLimitCollection *mongo.Collection
	}
	BuildingdataAccessServiceServer struct {
		protoOut.UnimplementedBuildingdataAccessServer
	}
)

func NewBAServer() *BuildingdataAccessServiceServer {
	return &BuildingdataAccessServiceServer{}
}

var ba *buildingdataAccess

func main() {
	conf := config.GetConfig()
	logger := logger.InitLogger("buildingdata-access")
	ba = &buildingdataAccess{
		logger: logger,
		conf:   conf,
	}

	bgContext := context.Background()

	ctx, cancel := context.WithTimeout(bgContext, 10*time.Second)
	defer cancel()

	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI(ba.conf.BuildingdataDbURL))
	if err != nil {
		ba.logger.Errorf("unable to connect to building access database: %w", err)

		panic(err)
	}
	// TODO implement?
	// dbDisconnect := func() {
	// 	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// 	defer cancel()
	// 	if err := dbClient.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }

	ctx, cancel = context.WithTimeout(bgContext, 2*time.Second)
	defer cancel()

	ba.logger.Infof("pinging buildingdata database")

	err = dbClient.Ping(ctx, readpref.Primary())
	if err != nil {
		ba.logger.Errorf("unable to ping to building access database: %w", err)

		panic(err)
	}

	ba.logger.Infof("connected to mongo DB for buildingdata database")

	ba.buildingLimitCollection = dbClient.Database("buildingdata").Collection("buildinglimit")
	ba.heightPlateauCollection = dbClient.Database("buildingdata").Collection("heightplateau")
	ba.splitBuildingLimitCollection = dbClient.Database("buildingdata").Collection("splitbuilding")

	// start gRPC server for BAAccess
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		ba.logger.Fatalf("Failed to listen: %w", err)
	}

	grpcServer := grpc.NewServer()

	protoOut.RegisterBuildingdataAccessServer(grpcServer, &BuildingdataAccessServiceServer{})
	ba.logger.Infof("Serving gRPC on 0.0.0.0:8080")

	go func() {
		ba.logger.Fatalf("failed to serve grpc server: %w", grpcServer.Serve(lis))
	}()

	// create apiGW server
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8080",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	gwmux := runtime.NewServeMux()
	err = protoOut.RegisterBuildingdataAccessHandler(context.Background(), gwmux, conn)
	if err != nil {
		ba.logger.Fatalf("Failed to register gateway: %w", err)
	}

	ba.logger.Infof("API GW server serving on port %s", ba.conf.GrpcPort)

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", ba.conf.GrpcPort),
		Handler: gwmux,
	}

	ba.logger.Fatalf("failed to listen and serve API GW server: %w", gwServer.ListenAndServe())
}
