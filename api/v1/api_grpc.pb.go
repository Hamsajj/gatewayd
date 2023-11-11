// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/v1/api.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GatewayDAdminAPIService_Version_FullMethodName         = "/api.v1.GatewayDAdminAPIService/Version"
	GatewayDAdminAPIService_GetGlobalConfig_FullMethodName = "/api.v1.GatewayDAdminAPIService/GetGlobalConfig"
	GatewayDAdminAPIService_GetPluginConfig_FullMethodName = "/api.v1.GatewayDAdminAPIService/GetPluginConfig"
	GatewayDAdminAPIService_GetPlugins_FullMethodName      = "/api.v1.GatewayDAdminAPIService/GetPlugins"
	GatewayDAdminAPIService_GetPools_FullMethodName        = "/api.v1.GatewayDAdminAPIService/GetPools"
	GatewayDAdminAPIService_GetProxies_FullMethodName      = "/api.v1.GatewayDAdminAPIService/GetProxies"
	GatewayDAdminAPIService_GetServers_FullMethodName      = "/api.v1.GatewayDAdminAPIService/GetServers"
)

// GatewayDAdminAPIServiceClient is the client API for GatewayDAdminAPIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GatewayDAdminAPIServiceClient interface {
	// Version returns the version of the GatewayD.
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	// GetGlobalConfig returns the global configuration of the GatewayD.
	GetGlobalConfig(ctx context.Context, in *Group, opts ...grpc.CallOption) (*structpb.Struct, error)
	// GetPluginConfig returns the configuration of the specified plugin.
	GetPluginConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*structpb.Struct, error)
	// GetPlugins returns the list of plugins installed on the GatewayD.
	GetPlugins(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PluginConfigs, error)
	// GetPools returns the list of pools configured on the GatewayD.
	GetPools(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*structpb.Struct, error)
	// GetProxies returns the list of proxies configured on the GatewayD.
	GetProxies(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*structpb.Struct, error)
	// GetServers returns the list of servers configured on the GatewayD.
	GetServers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*structpb.Struct, error)
}

type gatewayDAdminAPIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGatewayDAdminAPIServiceClient(cc grpc.ClientConnInterface) GatewayDAdminAPIServiceClient {
	return &gatewayDAdminAPIServiceClient{cc}
}

func (c *gatewayDAdminAPIServiceClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, GatewayDAdminAPIService_Version_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDAdminAPIServiceClient) GetGlobalConfig(ctx context.Context, in *Group, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, GatewayDAdminAPIService_GetGlobalConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDAdminAPIServiceClient) GetPluginConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, GatewayDAdminAPIService_GetPluginConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDAdminAPIServiceClient) GetPlugins(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*PluginConfigs, error) {
	out := new(PluginConfigs)
	err := c.cc.Invoke(ctx, GatewayDAdminAPIService_GetPlugins_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDAdminAPIServiceClient) GetPools(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, GatewayDAdminAPIService_GetPools_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDAdminAPIServiceClient) GetProxies(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, GatewayDAdminAPIService_GetProxies_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gatewayDAdminAPIServiceClient) GetServers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*structpb.Struct, error) {
	out := new(structpb.Struct)
	err := c.cc.Invoke(ctx, GatewayDAdminAPIService_GetServers_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GatewayDAdminAPIServiceServer is the server API for GatewayDAdminAPIService service.
// All implementations must embed UnimplementedGatewayDAdminAPIServiceServer
// for forward compatibility
type GatewayDAdminAPIServiceServer interface {
	// Version returns the version of the GatewayD.
	Version(context.Context, *emptypb.Empty) (*VersionResponse, error)
	// GetGlobalConfig returns the global configuration of the GatewayD.
	GetGlobalConfig(context.Context, *Group) (*structpb.Struct, error)
	// GetPluginConfig returns the configuration of the specified plugin.
	GetPluginConfig(context.Context, *emptypb.Empty) (*structpb.Struct, error)
	// GetPlugins returns the list of plugins installed on the GatewayD.
	GetPlugins(context.Context, *emptypb.Empty) (*PluginConfigs, error)
	// GetPools returns the list of pools configured on the GatewayD.
	GetPools(context.Context, *emptypb.Empty) (*structpb.Struct, error)
	// GetProxies returns the list of proxies configured on the GatewayD.
	GetProxies(context.Context, *emptypb.Empty) (*structpb.Struct, error)
	// GetServers returns the list of servers configured on the GatewayD.
	GetServers(context.Context, *emptypb.Empty) (*structpb.Struct, error)
	mustEmbedUnimplementedGatewayDAdminAPIServiceServer()
}

// UnimplementedGatewayDAdminAPIServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGatewayDAdminAPIServiceServer struct {
}

func (UnimplementedGatewayDAdminAPIServiceServer) Version(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedGatewayDAdminAPIServiceServer) GetGlobalConfig(context.Context, *Group) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGlobalConfig not implemented")
}
func (UnimplementedGatewayDAdminAPIServiceServer) GetPluginConfig(context.Context, *emptypb.Empty) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPluginConfig not implemented")
}
func (UnimplementedGatewayDAdminAPIServiceServer) GetPlugins(context.Context, *emptypb.Empty) (*PluginConfigs, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlugins not implemented")
}
func (UnimplementedGatewayDAdminAPIServiceServer) GetPools(context.Context, *emptypb.Empty) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPools not implemented")
}
func (UnimplementedGatewayDAdminAPIServiceServer) GetProxies(context.Context, *emptypb.Empty) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProxies not implemented")
}
func (UnimplementedGatewayDAdminAPIServiceServer) GetServers(context.Context, *emptypb.Empty) (*structpb.Struct, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServers not implemented")
}
func (UnimplementedGatewayDAdminAPIServiceServer) mustEmbedUnimplementedGatewayDAdminAPIServiceServer() {
}

// UnsafeGatewayDAdminAPIServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GatewayDAdminAPIServiceServer will
// result in compilation errors.
type UnsafeGatewayDAdminAPIServiceServer interface {
	mustEmbedUnimplementedGatewayDAdminAPIServiceServer()
}

func RegisterGatewayDAdminAPIServiceServer(s grpc.ServiceRegistrar, srv GatewayDAdminAPIServiceServer) {
	s.RegisterService(&GatewayDAdminAPIService_ServiceDesc, srv)
}

func _GatewayDAdminAPIService_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDAdminAPIServiceServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayDAdminAPIService_Version_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDAdminAPIServiceServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDAdminAPIService_GetGlobalConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Group)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDAdminAPIServiceServer).GetGlobalConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayDAdminAPIService_GetGlobalConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDAdminAPIServiceServer).GetGlobalConfig(ctx, req.(*Group))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDAdminAPIService_GetPluginConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDAdminAPIServiceServer).GetPluginConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayDAdminAPIService_GetPluginConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDAdminAPIServiceServer).GetPluginConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDAdminAPIService_GetPlugins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDAdminAPIServiceServer).GetPlugins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayDAdminAPIService_GetPlugins_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDAdminAPIServiceServer).GetPlugins(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDAdminAPIService_GetPools_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDAdminAPIServiceServer).GetPools(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayDAdminAPIService_GetPools_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDAdminAPIServiceServer).GetPools(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDAdminAPIService_GetProxies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDAdminAPIServiceServer).GetProxies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayDAdminAPIService_GetProxies_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDAdminAPIServiceServer).GetProxies(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GatewayDAdminAPIService_GetServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GatewayDAdminAPIServiceServer).GetServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GatewayDAdminAPIService_GetServers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GatewayDAdminAPIServiceServer).GetServers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// GatewayDAdminAPIService_ServiceDesc is the grpc.ServiceDesc for GatewayDAdminAPIService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GatewayDAdminAPIService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.v1.GatewayDAdminAPIService",
	HandlerType: (*GatewayDAdminAPIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _GatewayDAdminAPIService_Version_Handler,
		},
		{
			MethodName: "GetGlobalConfig",
			Handler:    _GatewayDAdminAPIService_GetGlobalConfig_Handler,
		},
		{
			MethodName: "GetPluginConfig",
			Handler:    _GatewayDAdminAPIService_GetPluginConfig_Handler,
		},
		{
			MethodName: "GetPlugins",
			Handler:    _GatewayDAdminAPIService_GetPlugins_Handler,
		},
		{
			MethodName: "GetPools",
			Handler:    _GatewayDAdminAPIService_GetPools_Handler,
		},
		{
			MethodName: "GetProxies",
			Handler:    _GatewayDAdminAPIService_GetProxies_Handler,
		},
		{
			MethodName: "GetServers",
			Handler:    _GatewayDAdminAPIService_GetServers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/v1/api.proto",
}
