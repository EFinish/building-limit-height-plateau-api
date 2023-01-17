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
	// ba.logger.Infof("got request!!! %+v", req.BuildingLimits.Features[0].Geometry.Coordinates)

	coordinates := req.BuildingLimits.Features[0].Geometry.Coordinates

	ba.logger.Infof("got request22!!! %+v", coordinates.GetKind())

	// if _, ok := wow.Kind.(*structpb.Value_ListValue); ok {
	// 	fmt.Println("Value is of type lsit")

	// }

	// coordinates, err := getCoordinateValues(ctx, coordinates.GetListValue().Values)

	// if err != nil {
	// 	// TODO
	// }

	// ba.logger.Infof("got request4444!!! %+v", coordinates)

	var foundCoordinates []Coordinate

	if _, isFirstValueList := coordinates.Kind.(*structpb.Value_ListValue); isFirstValueList {
		ba.logger.Infof("first layer is a list %+v", coordinates)
		for _, secondLayerValue := range coordinates.GetListValue().Values {
			if _, isSecondValueList := secondLayerValue.Kind.(*structpb.Value_ListValue); isSecondValueList {
				ba.logger.Infof("second layer is a list %+v", secondLayerValue)
				for _, thirdLayerValue := range secondLayerValue.GetListValue().Values {
					if _, isThirdValueList := thirdLayerValue.Kind.(*structpb.Value_ListValue); isThirdValueList {
						var number1 float64
						var number2 float64

						ba.logger.Infof("third layer is a list %+v", thirdLayerValue)

						for _, fourthLayerValue := range thirdLayerValue.GetListValue().Values {
							if _, isFourthValueNumber := fourthLayerValue.Kind.(*structpb.Value_NumberValue); isFourthValueNumber {
								ba.logger.Infof("fourth layer is a number! %+v", fourthLayerValue.GetNumberValue())
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

	ba.logger.Infof("found coordinates: %+v", foundCoordinates)

	// switch v := wow.GetListValue().Values[0].GetListValue().Values[0].GetListValue().Values[0].GetKind().(type) {
	// case *structpb.Value_StringValue:
	// 	fmt.Println("StringValue:", v.StringValue)
	// case *structpb.Value_NumberValue:
	// 	fmt.Println("NumberValue:", v.NumberValue)
	// 	// Handle other cases
	// case *structpb.Value_StructValue:
	// 	fmt.Println("um:", v.StructValue)
	// case *structpb.Value_ListValue:
	// 	fmt.Println("um3: ", v.ListValue.Values[0].GetListValue().Values[0].GetListValue().Values)
	// default:
	// 	// TODO throw error
	// 	fmt.Println("wow ", v)
	// }

	// fmt.Println("um3: ", wow.GetListValue().Values[0].GetListValue().Values[0].GetListValue().Values)

	// var potato Potato
	// unmarshal(wow, &potato)
	// ba.logger.Infof("got REQUEST333!!! %+v", potato)

	// for array, i := range wow[0] {

	// }

	return &protoOut.CreateBuildingdataResponse{}, nil
}

// func getCoordinateValues(ctx context.Context, values []*structpb.Value) ([]Coordinate, error) {
// 	var coordinates []Coordinate

// 	for _, value := range values {
// 		var coordinate Coordinate

// 		switch v := value.GetKind().(type) {
// 		case *structpb.Value_NumberValue:
// 			// if v.NumberValue == 0 {
// 			// 	// TODO write warning log
// 			// }
// 			// if coordinate.number1 == 0 {
// 			// 	coordinate.number1 = v.NumberValue
// 			// } else if coordinate.number2 == 0 {
// 			// 	coordinate.number2 = v.NumberValue
// 			// }
// 		case *structpb.Value_ListValue:
// 			var listCoordinates []Coordinate

// 			newCoordinates, err := getCoordinateValues(ctx, value.GetListValue().Values)

// 			if err != nil {
// 				return nil, fmt.Errorf("TODO3: %w", err)
// 			}

// 			// for _, newCoordinate := range newCoordinates {
// 			// 	coordinates = append(coordinates, newCoordinate)
// 			// }

// 			listCoordinates = append(listCoordinates, newCoordinates...)

// 			return listCoordinates, nil
// 		default:
// 			return nil, errors.New("TODO")
// 		}

// 		coordinates = append(coordinates, coordinate)
// 	}
// 	fmt.Println("maggot")

// 	return coordinates, nil
// }
