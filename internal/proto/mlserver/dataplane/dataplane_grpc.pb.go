// Copied from https://github.com/SeldonIO/MLServer/blob/1.2.4/proto/dataplane.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.5
// source: dataplane.proto

package dataplane

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	GRPCInferenceService_ServerLive_FullMethodName            = "/inference.GRPCInferenceService/ServerLive"
	GRPCInferenceService_ServerReady_FullMethodName           = "/inference.GRPCInferenceService/ServerReady"
	GRPCInferenceService_ModelReady_FullMethodName            = "/inference.GRPCInferenceService/ModelReady"
	GRPCInferenceService_ServerMetadata_FullMethodName        = "/inference.GRPCInferenceService/ServerMetadata"
	GRPCInferenceService_ModelMetadata_FullMethodName         = "/inference.GRPCInferenceService/ModelMetadata"
	GRPCInferenceService_ModelInfer_FullMethodName            = "/inference.GRPCInferenceService/ModelInfer"
	GRPCInferenceService_RepositoryIndex_FullMethodName       = "/inference.GRPCInferenceService/RepositoryIndex"
	GRPCInferenceService_RepositoryModelLoad_FullMethodName   = "/inference.GRPCInferenceService/RepositoryModelLoad"
	GRPCInferenceService_RepositoryModelUnload_FullMethodName = "/inference.GRPCInferenceService/RepositoryModelUnload"
)

// GRPCInferenceServiceClient is the client API for GRPCInferenceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Inference Server GRPC endpoints.
type GRPCInferenceServiceClient interface {
	// Check liveness of the inference server.
	ServerLive(ctx context.Context, in *ServerLiveRequest, opts ...grpc.CallOption) (*ServerLiveResponse, error)
	// Check readiness of the inference server.
	ServerReady(ctx context.Context, in *ServerReadyRequest, opts ...grpc.CallOption) (*ServerReadyResponse, error)
	// Check readiness of a model in the inference server.
	ModelReady(ctx context.Context, in *ModelReadyRequest, opts ...grpc.CallOption) (*ModelReadyResponse, error)
	// Get server metadata.
	ServerMetadata(ctx context.Context, in *ServerMetadataRequest, opts ...grpc.CallOption) (*ServerMetadataResponse, error)
	// Get model metadata.
	ModelMetadata(ctx context.Context, in *ModelMetadataRequest, opts ...grpc.CallOption) (*ModelMetadataResponse, error)
	// Perform inference using a specific model.
	ModelInfer(ctx context.Context, in *ModelInferRequest, opts ...grpc.CallOption) (*ModelInferResponse, error)
	// Get the index of model repository contents.
	RepositoryIndex(ctx context.Context, in *RepositoryIndexRequest, opts ...grpc.CallOption) (*RepositoryIndexResponse, error)
	// Load or reload a model from a repository.
	RepositoryModelLoad(ctx context.Context, in *RepositoryModelLoadRequest, opts ...grpc.CallOption) (*RepositoryModelLoadResponse, error)
	// Unload a model.
	RepositoryModelUnload(ctx context.Context, in *RepositoryModelUnloadRequest, opts ...grpc.CallOption) (*RepositoryModelUnloadResponse, error)
}

type gRPCInferenceServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGRPCInferenceServiceClient(cc grpc.ClientConnInterface) GRPCInferenceServiceClient {
	return &gRPCInferenceServiceClient{cc}
}

func (c *gRPCInferenceServiceClient) ServerLive(ctx context.Context, in *ServerLiveRequest, opts ...grpc.CallOption) (*ServerLiveResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServerLiveResponse)
	err := c.cc.Invoke(ctx, GRPCInferenceService_ServerLive_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCInferenceServiceClient) ServerReady(ctx context.Context, in *ServerReadyRequest, opts ...grpc.CallOption) (*ServerReadyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServerReadyResponse)
	err := c.cc.Invoke(ctx, GRPCInferenceService_ServerReady_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCInferenceServiceClient) ModelReady(ctx context.Context, in *ModelReadyRequest, opts ...grpc.CallOption) (*ModelReadyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ModelReadyResponse)
	err := c.cc.Invoke(ctx, GRPCInferenceService_ModelReady_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCInferenceServiceClient) ServerMetadata(ctx context.Context, in *ServerMetadataRequest, opts ...grpc.CallOption) (*ServerMetadataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServerMetadataResponse)
	err := c.cc.Invoke(ctx, GRPCInferenceService_ServerMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCInferenceServiceClient) ModelMetadata(ctx context.Context, in *ModelMetadataRequest, opts ...grpc.CallOption) (*ModelMetadataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ModelMetadataResponse)
	err := c.cc.Invoke(ctx, GRPCInferenceService_ModelMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCInferenceServiceClient) ModelInfer(ctx context.Context, in *ModelInferRequest, opts ...grpc.CallOption) (*ModelInferResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ModelInferResponse)
	err := c.cc.Invoke(ctx, GRPCInferenceService_ModelInfer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCInferenceServiceClient) RepositoryIndex(ctx context.Context, in *RepositoryIndexRequest, opts ...grpc.CallOption) (*RepositoryIndexResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RepositoryIndexResponse)
	err := c.cc.Invoke(ctx, GRPCInferenceService_RepositoryIndex_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCInferenceServiceClient) RepositoryModelLoad(ctx context.Context, in *RepositoryModelLoadRequest, opts ...grpc.CallOption) (*RepositoryModelLoadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RepositoryModelLoadResponse)
	err := c.cc.Invoke(ctx, GRPCInferenceService_RepositoryModelLoad_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gRPCInferenceServiceClient) RepositoryModelUnload(ctx context.Context, in *RepositoryModelUnloadRequest, opts ...grpc.CallOption) (*RepositoryModelUnloadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RepositoryModelUnloadResponse)
	err := c.cc.Invoke(ctx, GRPCInferenceService_RepositoryModelUnload_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GRPCInferenceServiceServer is the server API for GRPCInferenceService service.
// All implementations must embed UnimplementedGRPCInferenceServiceServer
// for forward compatibility.
//
// Inference Server GRPC endpoints.
type GRPCInferenceServiceServer interface {
	// Check liveness of the inference server.
	ServerLive(context.Context, *ServerLiveRequest) (*ServerLiveResponse, error)
	// Check readiness of the inference server.
	ServerReady(context.Context, *ServerReadyRequest) (*ServerReadyResponse, error)
	// Check readiness of a model in the inference server.
	ModelReady(context.Context, *ModelReadyRequest) (*ModelReadyResponse, error)
	// Get server metadata.
	ServerMetadata(context.Context, *ServerMetadataRequest) (*ServerMetadataResponse, error)
	// Get model metadata.
	ModelMetadata(context.Context, *ModelMetadataRequest) (*ModelMetadataResponse, error)
	// Perform inference using a specific model.
	ModelInfer(context.Context, *ModelInferRequest) (*ModelInferResponse, error)
	// Get the index of model repository contents.
	RepositoryIndex(context.Context, *RepositoryIndexRequest) (*RepositoryIndexResponse, error)
	// Load or reload a model from a repository.
	RepositoryModelLoad(context.Context, *RepositoryModelLoadRequest) (*RepositoryModelLoadResponse, error)
	// Unload a model.
	RepositoryModelUnload(context.Context, *RepositoryModelUnloadRequest) (*RepositoryModelUnloadResponse, error)
	mustEmbedUnimplementedGRPCInferenceServiceServer()
}

// UnimplementedGRPCInferenceServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedGRPCInferenceServiceServer struct{}

func (UnimplementedGRPCInferenceServiceServer) ServerLive(context.Context, *ServerLiveRequest) (*ServerLiveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerLive not implemented")
}
func (UnimplementedGRPCInferenceServiceServer) ServerReady(context.Context, *ServerReadyRequest) (*ServerReadyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerReady not implemented")
}
func (UnimplementedGRPCInferenceServiceServer) ModelReady(context.Context, *ModelReadyRequest) (*ModelReadyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModelReady not implemented")
}
func (UnimplementedGRPCInferenceServiceServer) ServerMetadata(context.Context, *ServerMetadataRequest) (*ServerMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ServerMetadata not implemented")
}
func (UnimplementedGRPCInferenceServiceServer) ModelMetadata(context.Context, *ModelMetadataRequest) (*ModelMetadataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModelMetadata not implemented")
}
func (UnimplementedGRPCInferenceServiceServer) ModelInfer(context.Context, *ModelInferRequest) (*ModelInferResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModelInfer not implemented")
}
func (UnimplementedGRPCInferenceServiceServer) RepositoryIndex(context.Context, *RepositoryIndexRequest) (*RepositoryIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepositoryIndex not implemented")
}
func (UnimplementedGRPCInferenceServiceServer) RepositoryModelLoad(context.Context, *RepositoryModelLoadRequest) (*RepositoryModelLoadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepositoryModelLoad not implemented")
}
func (UnimplementedGRPCInferenceServiceServer) RepositoryModelUnload(context.Context, *RepositoryModelUnloadRequest) (*RepositoryModelUnloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RepositoryModelUnload not implemented")
}
func (UnimplementedGRPCInferenceServiceServer) mustEmbedUnimplementedGRPCInferenceServiceServer() {}
func (UnimplementedGRPCInferenceServiceServer) testEmbeddedByValue()                              {}

// UnsafeGRPCInferenceServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GRPCInferenceServiceServer will
// result in compilation errors.
type UnsafeGRPCInferenceServiceServer interface {
	mustEmbedUnimplementedGRPCInferenceServiceServer()
}

func RegisterGRPCInferenceServiceServer(s grpc.ServiceRegistrar, srv GRPCInferenceServiceServer) {
	// If the following call pancis, it indicates UnimplementedGRPCInferenceServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&GRPCInferenceService_ServiceDesc, srv)
}

func _GRPCInferenceService_ServerLive_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerLiveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCInferenceServiceServer).ServerLive(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GRPCInferenceService_ServerLive_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCInferenceServiceServer).ServerLive(ctx, req.(*ServerLiveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCInferenceService_ServerReady_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerReadyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCInferenceServiceServer).ServerReady(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GRPCInferenceService_ServerReady_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCInferenceServiceServer).ServerReady(ctx, req.(*ServerReadyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCInferenceService_ModelReady_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelReadyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCInferenceServiceServer).ModelReady(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GRPCInferenceService_ModelReady_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCInferenceServiceServer).ModelReady(ctx, req.(*ModelReadyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCInferenceService_ServerMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCInferenceServiceServer).ServerMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GRPCInferenceService_ServerMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCInferenceServiceServer).ServerMetadata(ctx, req.(*ServerMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCInferenceService_ModelMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCInferenceServiceServer).ModelMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GRPCInferenceService_ModelMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCInferenceServiceServer).ModelMetadata(ctx, req.(*ModelMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCInferenceService_ModelInfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModelInferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCInferenceServiceServer).ModelInfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GRPCInferenceService_ModelInfer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCInferenceServiceServer).ModelInfer(ctx, req.(*ModelInferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCInferenceService_RepositoryIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCInferenceServiceServer).RepositoryIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GRPCInferenceService_RepositoryIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCInferenceServiceServer).RepositoryIndex(ctx, req.(*RepositoryIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCInferenceService_RepositoryModelLoad_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryModelLoadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCInferenceServiceServer).RepositoryModelLoad(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GRPCInferenceService_RepositoryModelLoad_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCInferenceServiceServer).RepositoryModelLoad(ctx, req.(*RepositoryModelLoadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GRPCInferenceService_RepositoryModelUnload_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RepositoryModelUnloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GRPCInferenceServiceServer).RepositoryModelUnload(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GRPCInferenceService_RepositoryModelUnload_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GRPCInferenceServiceServer).RepositoryModelUnload(ctx, req.(*RepositoryModelUnloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GRPCInferenceService_ServiceDesc is the grpc.ServiceDesc for GRPCInferenceService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GRPCInferenceService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inference.GRPCInferenceService",
	HandlerType: (*GRPCInferenceServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ServerLive",
			Handler:    _GRPCInferenceService_ServerLive_Handler,
		},
		{
			MethodName: "ServerReady",
			Handler:    _GRPCInferenceService_ServerReady_Handler,
		},
		{
			MethodName: "ModelReady",
			Handler:    _GRPCInferenceService_ModelReady_Handler,
		},
		{
			MethodName: "ServerMetadata",
			Handler:    _GRPCInferenceService_ServerMetadata_Handler,
		},
		{
			MethodName: "ModelMetadata",
			Handler:    _GRPCInferenceService_ModelMetadata_Handler,
		},
		{
			MethodName: "ModelInfer",
			Handler:    _GRPCInferenceService_ModelInfer_Handler,
		},
		{
			MethodName: "RepositoryIndex",
			Handler:    _GRPCInferenceService_RepositoryIndex_Handler,
		},
		{
			MethodName: "RepositoryModelLoad",
			Handler:    _GRPCInferenceService_RepositoryModelLoad_Handler,
		},
		{
			MethodName: "RepositoryModelUnload",
			Handler:    _GRPCInferenceService_RepositoryModelUnload_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dataplane.proto",
}
