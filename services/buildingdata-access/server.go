package main

import (
	"context"

	protoOut "github.com/EFinish/building-limit-height-plateau-api/proto/gen/buildingdataaccess/v1"
	"google.golang.org/protobuf/types/known/structpb"
)

func (s *BuildingdataAccessServiceServer) CreateBuildingdata(ctx context.Context, req *protoOut.CreateBuildingdataRequest) (*protoOut.CreateBuildingdataResponse, error) {
	// coordinates := req.BuildingLimits.Features[0].Geometry.Coordinates
	ba.logger.Infof("potato")

	return &protoOut.CreateBuildingdataResponse{}, nil
}

func parseCoordinatesFromFeature(ctx context.Context, feature *protoOut.RequestFeatures) ([]protoOut.Coordinate, error) {
	coordinatesData := feature.Geometry.Coordinates

	var foundCoordinates []protoOut.Coordinate

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
									// TODO throw warning
								}

								if longitude == 0 {
									longitude = fourthLayerValue.GetNumberValue()
								} else if latitude == 0 {
									latitude = fourthLayerValue.GetNumberValue()
								}
							}
						}

						if latitude != 0 && longitude != 0 {
							foundCoordinates = append(foundCoordinates, protoOut.Coordinate{Latitude: latitude, Longitude: longitude})
						}
					}
				}
			}
		}
	}

	return foundCoordinates, nil
}
