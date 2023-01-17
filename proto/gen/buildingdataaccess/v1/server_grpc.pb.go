// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: buildingdataaccess/v1/server.proto

package buildingdata

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BuildingdataAccessClient is the client API for BuildingdataAccess service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BuildingdataAccessClient interface {
	CreateBuildingdata(ctx context.Context, in *CreateBuildingdataRequest, opts ...grpc.CallOption) (*CreateBuildingdataResponse, error)
}

type buildingdataAccessClient struct {
	cc grpc.ClientConnInterface
}

func NewBuildingdataAccessClient(cc grpc.ClientConnInterface) BuildingdataAccessClient {
	return &buildingdataAccessClient{cc}
}

func (c *buildingdataAccessClient) CreateBuildingdata(ctx context.Context, in *CreateBuildingdataRequest, opts ...grpc.CallOption) (*CreateBuildingdataResponse, error) {
	out := new(CreateBuildingdataResponse)
	err := c.cc.Invoke(ctx, "/buildingdataaccess.v1.BuildingdataAccess/CreateBuildingdata", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BuildingdataAccessServer is the server API for BuildingdataAccess service.
// All implementations must embed UnimplementedBuildingdataAccessServer
// for forward compatibility
type BuildingdataAccessServer interface {
	CreateBuildingdata(context.Context, *CreateBuildingdataRequest) (*CreateBuildingdataResponse, error)
	mustEmbedUnimplementedBuildingdataAccessServer()
}

// UnimplementedBuildingdataAccessServer must be embedded to have forward compatible implementations.
type UnimplementedBuildingdataAccessServer struct {
}

func (UnimplementedBuildingdataAccessServer) CreateBuildingdata(context.Context, *CreateBuildingdataRequest) (*CreateBuildingdataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBuildingdata not implemented")
}
func (UnimplementedBuildingdataAccessServer) mustEmbedUnimplementedBuildingdataAccessServer() {}

// UnsafeBuildingdataAccessServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BuildingdataAccessServer will
// result in compilation errors.
type UnsafeBuildingdataAccessServer interface {
	mustEmbedUnimplementedBuildingdataAccessServer()
}

func RegisterBuildingdataAccessServer(s grpc.ServiceRegistrar, srv BuildingdataAccessServer) {
	s.RegisterService(&BuildingdataAccess_ServiceDesc, srv)
}

func _BuildingdataAccess_CreateBuildingdata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBuildingdataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuildingdataAccessServer).CreateBuildingdata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/buildingdataaccess.v1.BuildingdataAccess/CreateBuildingdata",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuildingdataAccessServer).CreateBuildingdata(ctx, req.(*CreateBuildingdataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BuildingdataAccess_ServiceDesc is the grpc.ServiceDesc for BuildingdataAccess service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BuildingdataAccess_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "buildingdataaccess.v1.BuildingdataAccess",
	HandlerType: (*BuildingdataAccessServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBuildingdata",
			Handler:    _BuildingdataAccess_CreateBuildingdata_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "buildingdataaccess/v1/server.proto",
}
