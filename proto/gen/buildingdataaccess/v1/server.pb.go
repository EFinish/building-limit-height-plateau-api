// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: buildingdataaccess/v1/server.proto

package buildingdata

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	structpb "google.golang.org/protobuf/types/known/structpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateBuildingdataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TODO
	BuildingLimits *RequestBuildingLimit `protobuf:"bytes,1,opt,name=building_limits,json=buildingLimits,proto3" json:"building_limits,omitempty"`
	HeightPlateaus *RequestHeightPlateau `protobuf:"bytes,2,opt,name=height_plateaus,json=heightPlateaus,proto3" json:"height_plateaus,omitempty"`
}

func (x *CreateBuildingdataRequest) Reset() {
	*x = CreateBuildingdataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildingdataaccess_v1_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBuildingdataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBuildingdataRequest) ProtoMessage() {}

func (x *CreateBuildingdataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_buildingdataaccess_v1_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBuildingdataRequest.ProtoReflect.Descriptor instead.
func (*CreateBuildingdataRequest) Descriptor() ([]byte, []int) {
	return file_buildingdataaccess_v1_server_proto_rawDescGZIP(), []int{0}
}

func (x *CreateBuildingdataRequest) GetBuildingLimits() *RequestBuildingLimit {
	if x != nil {
		return x.BuildingLimits
	}
	return nil
}

func (x *CreateBuildingdataRequest) GetHeightPlateaus() *RequestHeightPlateau {
	if x != nil {
		return x.HeightPlateaus
	}
	return nil
}

type CreateBuildingdataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateBuildingdataResponse) Reset() {
	*x = CreateBuildingdataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildingdataaccess_v1_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateBuildingdataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateBuildingdataResponse) ProtoMessage() {}

func (x *CreateBuildingdataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_buildingdataaccess_v1_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateBuildingdataResponse.ProtoReflect.Descriptor instead.
func (*CreateBuildingdataResponse) Descriptor() ([]byte, []int) {
	return file_buildingdataaccess_v1_server_proto_rawDescGZIP(), []int{1}
}

type RequestBuildingLimit struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string             `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Features []*RequestFeatures `protobuf:"bytes,2,rep,name=features,proto3" json:"features,omitempty"`
}

func (x *RequestBuildingLimit) Reset() {
	*x = RequestBuildingLimit{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildingdataaccess_v1_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestBuildingLimit) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestBuildingLimit) ProtoMessage() {}

func (x *RequestBuildingLimit) ProtoReflect() protoreflect.Message {
	mi := &file_buildingdataaccess_v1_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestBuildingLimit.ProtoReflect.Descriptor instead.
func (*RequestBuildingLimit) Descriptor() ([]byte, []int) {
	return file_buildingdataaccess_v1_server_proto_rawDescGZIP(), []int{2}
}

func (x *RequestBuildingLimit) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RequestBuildingLimit) GetFeatures() []*RequestFeatures {
	if x != nil {
		return x.Features
	}
	return nil
}

type RequestHeightPlateau struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type     string             `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Features []*RequestFeatures `protobuf:"bytes,2,rep,name=features,proto3" json:"features,omitempty"`
}

func (x *RequestHeightPlateau) Reset() {
	*x = RequestHeightPlateau{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildingdataaccess_v1_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestHeightPlateau) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestHeightPlateau) ProtoMessage() {}

func (x *RequestHeightPlateau) ProtoReflect() protoreflect.Message {
	mi := &file_buildingdataaccess_v1_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestHeightPlateau.ProtoReflect.Descriptor instead.
func (*RequestHeightPlateau) Descriptor() ([]byte, []int) {
	return file_buildingdataaccess_v1_server_proto_rawDescGZIP(), []int{3}
}

func (x *RequestHeightPlateau) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RequestHeightPlateau) GetFeatures() []*RequestFeatures {
	if x != nil {
		return x.Features
	}
	return nil
}

type RequestFeatures struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type       string             `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Properties *RequestProperties `protobuf:"bytes,2,opt,name=properties,proto3" json:"properties,omitempty"`
	Geometry   *RequestGeometry   `protobuf:"bytes,3,opt,name=geometry,proto3" json:"geometry,omitempty"`
}

func (x *RequestFeatures) Reset() {
	*x = RequestFeatures{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildingdataaccess_v1_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestFeatures) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestFeatures) ProtoMessage() {}

func (x *RequestFeatures) ProtoReflect() protoreflect.Message {
	mi := &file_buildingdataaccess_v1_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestFeatures.ProtoReflect.Descriptor instead.
func (*RequestFeatures) Descriptor() ([]byte, []int) {
	return file_buildingdataaccess_v1_server_proto_rawDescGZIP(), []int{4}
}

func (x *RequestFeatures) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RequestFeatures) GetProperties() *RequestProperties {
	if x != nil {
		return x.Properties
	}
	return nil
}

func (x *RequestFeatures) GetGeometry() *RequestGeometry {
	if x != nil {
		return x.Geometry
	}
	return nil
}

type RequestProperties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Elevation float64 `protobuf:"fixed64,1,opt,name=elevation,proto3" json:"elevation,omitempty"`
}

func (x *RequestProperties) Reset() {
	*x = RequestProperties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildingdataaccess_v1_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestProperties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestProperties) ProtoMessage() {}

func (x *RequestProperties) ProtoReflect() protoreflect.Message {
	mi := &file_buildingdataaccess_v1_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestProperties.ProtoReflect.Descriptor instead.
func (*RequestProperties) Descriptor() ([]byte, []int) {
	return file_buildingdataaccess_v1_server_proto_rawDescGZIP(), []int{5}
}

func (x *RequestProperties) GetElevation() float64 {
	if x != nil {
		return x.Elevation
	}
	return 0
}

type RequestGeometry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string                `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Coordinates []*structpb.ListValue `protobuf:"bytes,2,rep,name=coordinates,proto3" json:"coordinates,omitempty"`
}

func (x *RequestGeometry) Reset() {
	*x = RequestGeometry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buildingdataaccess_v1_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestGeometry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestGeometry) ProtoMessage() {}

func (x *RequestGeometry) ProtoReflect() protoreflect.Message {
	mi := &file_buildingdataaccess_v1_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestGeometry.ProtoReflect.Descriptor instead.
func (*RequestGeometry) Descriptor() ([]byte, []int) {
	return file_buildingdataaccess_v1_server_proto_rawDescGZIP(), []int{6}
}

func (x *RequestGeometry) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *RequestGeometry) GetCoordinates() []*structpb.ListValue {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

var File_buildingdataaccess_v1_server_proto protoreflect.FileDescriptor

var file_buildingdataaccess_v1_server_proto_rawDesc = []byte{
	0x0a, 0x22, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x64, 0x61, 0x74, 0x61, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x64, 0x61,
	0x74, 0x61, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x1a, 0x28, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x69, 0x6e, 0x67, 0x64, 0x61, 0x74, 0x61, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc7, 0x01, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x69, 0x6e, 0x67, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x54, 0x0a, 0x0f, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x6c, 0x69, 0x6d, 0x69,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x69, 0x6e, 0x67, 0x64, 0x61, 0x74, 0x61, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x52, 0x0e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x4c,
	0x69, 0x6d, 0x69, 0x74, 0x73, 0x12, 0x54, 0x0a, 0x0f, 0x68, 0x65, 0x69, 0x67, 0x68, 0x74, 0x5f,
	0x70, 0x6c, 0x61, 0x74, 0x65, 0x61, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b,
	0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x64, 0x61, 0x74, 0x61, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x48, 0x65,
	0x69, 0x67, 0x68, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x61, 0x75, 0x52, 0x0e, 0x68, 0x65, 0x69,
	0x67, 0x68, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x61, 0x75, 0x73, 0x22, 0x1c, 0x0a, 0x1a, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x64, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x6e, 0x0a, 0x14, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x4c, 0x69, 0x6d, 0x69,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69,
	0x6e, 0x67, 0x64, 0x61, 0x74, 0x61, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52,
	0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0x6e, 0x0a, 0x14, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x48, 0x65, 0x69, 0x67, 0x68, 0x74, 0x50, 0x6c, 0x61, 0x74, 0x65, 0x61,
	0x75, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x42, 0x0a, 0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69,
	0x6e, 0x67, 0x64, 0x61, 0x74, 0x61, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x52,
	0x08, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x22, 0xb3, 0x01, 0x0a, 0x0f, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70,
	0x65, 0x12, 0x48, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67,
	0x64, 0x61, 0x74, 0x61, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x0a, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x42, 0x0a, 0x08, 0x67,
	0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x64, 0x61, 0x74, 0x61, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6f,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x52, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x22,
	0x31, 0x0a, 0x11, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09, 0x65, 0x6c, 0x65, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x22, 0x63, 0x0a, 0x0f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x47, 0x65, 0x6f,
	0x6d, 0x65, 0x74, 0x72, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x3c, 0x0a, 0x0b, 0x63, 0x6f, 0x6f,
	0x72, 0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72,
	0x64, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x32, 0xad, 0x01, 0x0a, 0x12, 0x42, 0x75, 0x69, 0x6c,
	0x64, 0x69, 0x6e, 0x67, 0x64, 0x61, 0x74, 0x61, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x96,
	0x01, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e,
	0x67, 0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67,
	0x64, 0x61, 0x74, 0x61, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x69,
	0x6e, 0x67, 0x64, 0x61, 0x74, 0x61, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x42, 0x75, 0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x64, 0x61,
	0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x15, 0x3a, 0x01, 0x2a, 0x22, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x64,
	0x69, 0x6e, 0x67, 0x64, 0x61, 0x74, 0x61, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x6f, 0x2f, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x3b, 0x62, 0x75,
	0x69, 0x6c, 0x64, 0x69, 0x6e, 0x67, 0x64, 0x61, 0x74, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_buildingdataaccess_v1_server_proto_rawDescOnce sync.Once
	file_buildingdataaccess_v1_server_proto_rawDescData = file_buildingdataaccess_v1_server_proto_rawDesc
)

func file_buildingdataaccess_v1_server_proto_rawDescGZIP() []byte {
	file_buildingdataaccess_v1_server_proto_rawDescOnce.Do(func() {
		file_buildingdataaccess_v1_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_buildingdataaccess_v1_server_proto_rawDescData)
	})
	return file_buildingdataaccess_v1_server_proto_rawDescData
}

var file_buildingdataaccess_v1_server_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_buildingdataaccess_v1_server_proto_goTypes = []interface{}{
	(*CreateBuildingdataRequest)(nil),  // 0: buildingdataaccess.v1.CreateBuildingdataRequest
	(*CreateBuildingdataResponse)(nil), // 1: buildingdataaccess.v1.CreateBuildingdataResponse
	(*RequestBuildingLimit)(nil),       // 2: buildingdataaccess.v1.RequestBuildingLimit
	(*RequestHeightPlateau)(nil),       // 3: buildingdataaccess.v1.RequestHeightPlateau
	(*RequestFeatures)(nil),            // 4: buildingdataaccess.v1.RequestFeatures
	(*RequestProperties)(nil),          // 5: buildingdataaccess.v1.RequestProperties
	(*RequestGeometry)(nil),            // 6: buildingdataaccess.v1.RequestGeometry
	(*structpb.ListValue)(nil),         // 7: google.protobuf.ListValue
}
var file_buildingdataaccess_v1_server_proto_depIdxs = []int32{
	2, // 0: buildingdataaccess.v1.CreateBuildingdataRequest.building_limits:type_name -> buildingdataaccess.v1.RequestBuildingLimit
	3, // 1: buildingdataaccess.v1.CreateBuildingdataRequest.height_plateaus:type_name -> buildingdataaccess.v1.RequestHeightPlateau
	4, // 2: buildingdataaccess.v1.RequestBuildingLimit.features:type_name -> buildingdataaccess.v1.RequestFeatures
	4, // 3: buildingdataaccess.v1.RequestHeightPlateau.features:type_name -> buildingdataaccess.v1.RequestFeatures
	5, // 4: buildingdataaccess.v1.RequestFeatures.properties:type_name -> buildingdataaccess.v1.RequestProperties
	6, // 5: buildingdataaccess.v1.RequestFeatures.geometry:type_name -> buildingdataaccess.v1.RequestGeometry
	7, // 6: buildingdataaccess.v1.RequestGeometry.coordinates:type_name -> google.protobuf.ListValue
	0, // 7: buildingdataaccess.v1.BuildingdataAccess.CreateBuildingdata:input_type -> buildingdataaccess.v1.CreateBuildingdataRequest
	1, // 8: buildingdataaccess.v1.BuildingdataAccess.CreateBuildingdata:output_type -> buildingdataaccess.v1.CreateBuildingdataResponse
	8, // [8:9] is the sub-list for method output_type
	7, // [7:8] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_buildingdataaccess_v1_server_proto_init() }
func file_buildingdataaccess_v1_server_proto_init() {
	if File_buildingdataaccess_v1_server_proto != nil {
		return
	}
	file_buildingdataaccess_v1_buildingdata_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_buildingdataaccess_v1_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBuildingdataRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buildingdataaccess_v1_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateBuildingdataResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buildingdataaccess_v1_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestBuildingLimit); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buildingdataaccess_v1_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestHeightPlateau); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buildingdataaccess_v1_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestFeatures); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buildingdataaccess_v1_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestProperties); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buildingdataaccess_v1_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestGeometry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buildingdataaccess_v1_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_buildingdataaccess_v1_server_proto_goTypes,
		DependencyIndexes: file_buildingdataaccess_v1_server_proto_depIdxs,
		MessageInfos:      file_buildingdataaccess_v1_server_proto_msgTypes,
	}.Build()
	File_buildingdataaccess_v1_server_proto = out.File
	file_buildingdataaccess_v1_server_proto_rawDesc = nil
	file_buildingdataaccess_v1_server_proto_goTypes = nil
	file_buildingdataaccess_v1_server_proto_depIdxs = nil
}
