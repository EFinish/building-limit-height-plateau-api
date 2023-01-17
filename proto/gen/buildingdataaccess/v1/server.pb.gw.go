// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: buildingdataaccess/v1/server.proto

/*
Package buildingdata is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package buildingdata

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_BuildingdataAccess_CreateBuildingdata_0(ctx context.Context, marshaler runtime.Marshaler, client BuildingdataAccessClient, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateBuildingdataRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := client.CreateBuildingdata(ctx, &protoReq, grpc.Header(&metadata.HeaderMD), grpc.Trailer(&metadata.TrailerMD))
	return msg, metadata, err

}

func local_request_BuildingdataAccess_CreateBuildingdata_0(ctx context.Context, marshaler runtime.Marshaler, server BuildingdataAccessServer, req *http.Request, pathParams map[string]string) (proto.Message, runtime.ServerMetadata, error) {
	var protoReq CreateBuildingdataRequest
	var metadata runtime.ServerMetadata

	newReader, berr := utilities.IOReaderFactory(req.Body)
	if berr != nil {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", berr)
	}
	if err := marshaler.NewDecoder(newReader()).Decode(&protoReq); err != nil && err != io.EOF {
		return nil, metadata, status.Errorf(codes.InvalidArgument, "%v", err)
	}

	msg, err := server.CreateBuildingdata(ctx, &protoReq)
	return msg, metadata, err

}

// RegisterBuildingdataAccessHandlerServer registers the http handlers for service BuildingdataAccess to "mux".
// UnaryRPC     :call BuildingdataAccessServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterBuildingdataAccessHandlerFromEndpoint instead.
func RegisterBuildingdataAccessHandlerServer(ctx context.Context, mux *runtime.ServeMux, server BuildingdataAccessServer) error {

	mux.Handle("POST", pattern_BuildingdataAccess_CreateBuildingdata_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		var stream runtime.ServerTransportStream
		ctx = grpc.NewContextWithServerTransportStream(ctx, &stream)
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateIncomingContext(ctx, mux, req, "/buildingdataaccess.v1.BuildingdataAccess/CreateBuildingdata", runtime.WithHTTPPathPattern("/v1/buildingdata"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := local_request_BuildingdataAccess_CreateBuildingdata_0(annotatedContext, inboundMarshaler, server, req, pathParams)
		md.HeaderMD, md.TrailerMD = metadata.Join(md.HeaderMD, stream.Header()), metadata.Join(md.TrailerMD, stream.Trailer())
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BuildingdataAccess_CreateBuildingdata_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

// RegisterBuildingdataAccessHandlerFromEndpoint is same as RegisterBuildingdataAccessHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterBuildingdataAccessHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterBuildingdataAccessHandler(ctx, mux, conn)
}

// RegisterBuildingdataAccessHandler registers the http handlers for service BuildingdataAccess to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterBuildingdataAccessHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterBuildingdataAccessHandlerClient(ctx, mux, NewBuildingdataAccessClient(conn))
}

// RegisterBuildingdataAccessHandlerClient registers the http handlers for service BuildingdataAccess
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "BuildingdataAccessClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "BuildingdataAccessClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "BuildingdataAccessClient" to call the correct interceptors.
func RegisterBuildingdataAccessHandlerClient(ctx context.Context, mux *runtime.ServeMux, client BuildingdataAccessClient) error {

	mux.Handle("POST", pattern_BuildingdataAccess_CreateBuildingdata_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		var err error
		var annotatedContext context.Context
		annotatedContext, err = runtime.AnnotateContext(ctx, mux, req, "/buildingdataaccess.v1.BuildingdataAccess/CreateBuildingdata", runtime.WithHTTPPathPattern("/v1/buildingdata"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_BuildingdataAccess_CreateBuildingdata_0(annotatedContext, inboundMarshaler, client, req, pathParams)
		annotatedContext = runtime.NewServerMetadataContext(annotatedContext, md)
		if err != nil {
			runtime.HTTPError(annotatedContext, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_BuildingdataAccess_CreateBuildingdata_0(annotatedContext, mux, outboundMarshaler, w, req, resp, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_BuildingdataAccess_CreateBuildingdata_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"v1", "buildingdata"}, ""))
)

var (
	forward_BuildingdataAccess_CreateBuildingdata_0 = runtime.ForwardResponseMessage
)
