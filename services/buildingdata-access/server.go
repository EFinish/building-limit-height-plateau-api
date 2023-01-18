package main

import (
	"context"
	"fmt"

	protoOut "github.com/EFinish/building-limit-height-plateau-api/proto/gen/buildingdataaccess/v1"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/golang/geo/s2"
)

func (s *BuildingdataAccessServiceServer) CreateBuildingdata(ctx context.Context, req *protoOut.CreateBuildingdataRequest) (*protoOut.CreateBuildingdataResponse, error) {
	var buildingLimitPolygons []s2.Polygon

	for _, feature := range req.BuildingLimits.Features {

		coordinates, err := getCoordinatesFromFeature(ctx, feature)

		if err != nil {
			return nil, fmt.Errorf("while parsing coordinates from a feature: %w", err)
		}

		_, err = insertBuildingLimit(ctx, protoOut.BuildingLimit{PolygonCoordinates: coordinates})

		if err != nil {
			return nil, fmt.Errorf("while saving new building limit: %w", err)
		}

		polygon := createPolygon(ctx, coordinates, 0)
		buildingLimitPolygons = append(buildingLimitPolygons, polygon)
	}

	var heightPlateauPolygons []s2.Polygon

	for _, feature := range req.HeightPlateaus.Features {

		coordinates, err := getCoordinatesFromFeature(ctx, feature)

		if err != nil {
			return nil, fmt.Errorf("while parsing coordinates from a feature: %w", err)
		}

		_, err = insertHeightPlateau(ctx, protoOut.HeightPlateau{PolygonCoordinates: coordinates, Elevation: feature.Properties.Elevation})

		if err != nil {
			return nil, fmt.Errorf("while saving new height plateau: %w", err)
		}

		polygon := createPolygon(ctx, coordinates, 0)
		heightPlateauPolygons = append(heightPlateauPolygons, polygon)
	}

	splitBuildingLimits := getSplitBuildingLimitsFromPolygons(ctx, buildingLimitPolygons, heightPlateauPolygons)

	for _, sbl := range splitBuildingLimits {
		_, err := insertSplitBuildingLimit(ctx, *sbl)

		if err != nil {
			return nil, fmt.Errorf("while inserting split building limit: %w", err)
		}
	}

	return &protoOut.CreateBuildingdataResponse{}, nil
}

func getSplitBuildingLimitsFromPolygons(ctx context.Context, buildingLimitPolygons []s2.Polygon, heightPlateauPolygons []s2.Polygon) []*protoOut.SplitBuildingLimit {
	var splitBuildingLimits []*protoOut.SplitBuildingLimit

	for _, blPolygon := range buildingLimitPolygons {
		for _, hpPolygon := range heightPlateauPolygons {
			if !blPolygon.Intersects(&hpPolygon) {
				ba.logger.Debugf("polygons do not intersect")

				continue
			}
			// TODO calculate intersection and add to array

		}
	}

	return splitBuildingLimits
}

func createPolygon(ctx context.Context, coordinates []*protoOut.Coordinate, elevation float64) s2.Polygon {
	// create points from coordinates
	var points []s2.Point

	for _, coord := range coordinates {
		point := s2.PointFromCoords(coord.Latitude, coord.Longitude, elevation)
		points = append(points, point)
	}

	// create loop from points
	loop := s2.LoopFromPoints(points)

	// create polygon from loop
	polygon := s2.PolygonFromLoops([]*s2.Loop{loop})

	return *polygon
}

func getCoordinatesFromFeature(ctx context.Context, feature *protoOut.RequestFeatures) ([]*protoOut.Coordinate, error) {
	coordinatesData := feature.Geometry.Coordinates

	var foundCoordinates []*protoOut.Coordinate

	if _, isFirstValueList := coordinatesData.Kind.(*structpb.Value_ListValue); isFirstValueList {
		ba.logger.Debugf("first layer is a list")
		for _, secondLayerValue := range coordinatesData.GetListValue().Values {
			if _, isSecondValueList := secondLayerValue.Kind.(*structpb.Value_ListValue); isSecondValueList {
				ba.logger.Debugf("second layer is a list")
				for _, thirdLayerValue := range secondLayerValue.GetListValue().Values {
					if _, isThirdValueList := thirdLayerValue.Kind.(*structpb.Value_ListValue); isThirdValueList {
						var latitude float64
						var longitude float64

						ba.logger.Debugf("third layer is a list")

						for _, fourthLayerValue := range thirdLayerValue.GetListValue().Values {
							if _, isFourthValueNumber := fourthLayerValue.Kind.(*structpb.Value_NumberValue); isFourthValueNumber {
								ba.logger.Debugf("fourth layer is a number")

								if fourthLayerValue.GetNumberValue() == 0 {
									return nil, fmt.Errorf("longitude or a latitude value of 0.90 detected")
								}

								if longitude == 0 {
									longitude = fourthLayerValue.GetNumberValue()
								} else if latitude == 0 {
									latitude = fourthLayerValue.GetNumberValue()
								}
							}
						}

						if latitude != 0 && longitude != 0 {
							foundCoordinates = append(foundCoordinates, &protoOut.Coordinate{Latitude: latitude, Longitude: longitude})
						}
					}
				}
			}
		}
	}

	return foundCoordinates, nil
}
