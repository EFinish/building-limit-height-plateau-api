package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

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

var ba *buildingdataAccess

func main() {
	conf := config.GetConfig()
	logger := logger.InitLogger("buildingdata-access")
	ba = &buildingdataAccess{
		// TODO implement logger?
		logger: logger,
		conf:   conf,
	}

	bgContext := context.Background()

	ctx, cancel := context.WithTimeout(bgContext, 10*time.Second)
	defer cancel()

	dbClient, err := mongo.Connect(ctx, options.Client().ApplyURI(ba.conf.BuildingdataDbURL))
	if err != nil {
		fmt.Printf("unable to connect to building access database: %v", err)

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

	fmt.Println("pinging buildingdata database")

	err = dbClient.Ping(ctx, readpref.Primary())
	if err != nil {
		fmt.Printf("unable to ping to building access database: %v", err)

		panic(err)
	}

	fmt.Println("connected to mongo DB for buildingdata database")

	ba.buildingLimitCollection = dbClient.Database("buildingdata").Collection("buildinglimit")
	ba.heightPlateauCollection = dbClient.Database("buildingdata").Collection("heightplateau")
	ba.splitBuildingLimitCollection = dbClient.Database("buildingdata").Collection("splitbuilding")

	fmt.Printf("Listening on port %v\n", conf.GrpcPort)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", conf.GrpcPort))

	if err != nil {
		panic(err)
	}
	defer lis.Close()

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	buildingdataAccessServer := BuildingdataAccessServiceServer{}
	protoOut.RegisterBuildingdataAccessServer(grpcServer, &buildingdataAccessServer)

	ba.logger.Infof("gRPC server starting")

	if err := grpcServer.Serve(lis); err != nil {
		ba.logger.Fatalf("failed to start gRPC server: %v", err)
	}
}
