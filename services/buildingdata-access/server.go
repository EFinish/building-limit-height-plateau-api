package main

import (
	"context"

	protoOut "github.com/EFinish/building-limit-height-plateau-api/proto/gen/buildingdataaccess/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

type Coordinate struct {
	number1 float64
	number2 float64
}

func (s *BuildingdataAccessServiceServer) CreateBuildingdata(ctx context.Context, req *protoOut.CreateBuildingdataRequest) (*protoOut.CreateBuildingdataResponse, error) {
	// coordinates := req.BuildingLimits.Features[0].Geometry.Coordinates

	return &protoOut.CreateBuildingdataResponse{}, nil
}

func parseCoordinatesFromFeature(ctx context.Context, feature *protoOut.RequestFeatures) ([]Coordinate, error) {
	coordinatesData := feature.Geometry.Coordinates

	var foundCoordinates []Coordinate

	if _, isFirstValueList := coordinatesData.Kind.(*structpb.Value_ListValue); isFirstValueList {
		ba.logger.Debugf("first layer is a list")
		for _, secondLayerValue := range coordinatesData.GetListValue().Values {
			if _, isSecondValueList := secondLayerValue.Kind.(*structpb.Value_ListValue); isSecondValueList {
				ba.logger.Debugf("second layer is a list")
				for _, thirdLayerValue := range secondLayerValue.GetListValue().Values {
					if _, isThirdValueList := thirdLayerValue.Kind.(*structpb.Value_ListValue); isThirdValueList {
						var number1 float64
						var number2 float64

						ba.logger.Debugf("third layer is a list")

						for _, fourthLayerValue := range thirdLayerValue.GetListValue().Values {
							if _, isFourthValueNumber := fourthLayerValue.Kind.(*structpb.Value_NumberValue); isFourthValueNumber {
								ba.logger.Debugf("fourth layer is a number")

								if fourthLayerValue.GetNumberValue() == 0 {
									// TODO throw warning
								}

								if number1 == 0 {
									number1 = fourthLayerValue.GetNumberValue()
								} else if number2 == 0 {
									number2 = fourthLayerValue.GetNumberValue()
								}
							}
						}

						if number1 != 0 && number2 != 0 {
							foundCoordinates = append(foundCoordinates, Coordinate{number1: number1, number2: number2})
						}
					}
				}
			}
		}
	}

	return foundCoordinates, nil
}
