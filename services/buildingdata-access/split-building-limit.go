package main

import (
	"context"

	protoOut "github.com/EFinish/building-limit-height-plateau-api/proto/gen/buildingdataaccess/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func insertSplitBuildingLimit(ctx context.Context, splitBuildingLimit protoOut.SplitBuildingLimit) (*protoOut.SplitBuildingLimit, error) {
	ba.logger.Debugf("inserting new split building limit record")

	newSplitBuildingLimit := bson.M{
		"polygon_coordinates":                    splitBuildingLimit.PolygonCoordinates,
		"corresponding_height_plateau_elevation": splitBuildingLimit.CorrespondingHeightPlateauElevation,
	}
	vehicle, err := ba.splitBuildingLimitCollection.InsertOne(ctx, newSplitBuildingLimit)

	if err != nil {
		return nil, err
	}

	insertedId := vehicle.InsertedID.(primitive.ObjectID).Hex()
	splitBuildingLimit.Id = &insertedId

	return &splitBuildingLimit, nil
}
