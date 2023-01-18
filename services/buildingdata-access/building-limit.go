package main

import (
	"context"

	protoOut "github.com/EFinish/building-limit-height-plateau-api/proto/gen/buildingdataaccess/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func insertBuildingLimit(ctx context.Context, buildingLimit protoOut.BuildingLimit) (*protoOut.BuildingLimit, error) {
	ba.logger.Debugf("inserting new building limit record")

	newBuildingLimit := bson.M{
		"polygon_coordinates": buildingLimit.PolygonCoordinates,
	}
	vehicle, err := ba.buildingLimitCollection.InsertOne(ctx, newBuildingLimit)

	if err != nil {
		return nil, err
	}

	insertedId := vehicle.InsertedID.(primitive.ObjectID).Hex()
	buildingLimit.Id = &insertedId

	return &buildingLimit, nil
}
