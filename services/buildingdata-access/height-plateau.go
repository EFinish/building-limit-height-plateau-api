package main

import (
	"context"

	protoOut "github.com/EFinish/building-limit-height-plateau-api/proto/gen/buildingdataaccess/v1"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func insertHeightPlateau(ctx context.Context, heightPlateau protoOut.HeightPlateau) (*protoOut.HeightPlateau, error) {
	ba.logger.Debugf("inserting new height plateau record")

	newHeightPlateau := bson.M{
		"polygon_coordinates": heightPlateau.PolygonCoordinates,
		"elevation":           heightPlateau.Elevation,
	}
	vehicle, err := ba.heightPlateauCollection.InsertOne(ctx, newHeightPlateau)

	if err != nil {
		return nil, err
	}

	insertedId := vehicle.InsertedID.(primitive.ObjectID).Hex()
	heightPlateau.Id = &insertedId

	return &heightPlateau, nil
}
