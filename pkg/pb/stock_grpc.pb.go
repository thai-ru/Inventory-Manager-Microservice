// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// StockLevelsClient is the client API for StockLevels service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StockLevelsClient interface {
	AddStockLevel(ctx context.Context, in *AddStockLevelRequest, opts ...grpc.CallOption) (*AddStockLevelResponse, error)
	GetStockLevel(ctx context.Context, in *GetStockLevelRequest, opts ...grpc.CallOption) (*GetStockLevelResponse, error)
	UpdateStockLevel(ctx context.Context, in *UpdateStockLevelRequest, opts ...grpc.CallOption) (*GetStockLevelResponse, error)
}

type stockLevelsClient struct {
	cc grpc.ClientConnInterface
}

func NewStockLevelsClient(cc grpc.ClientConnInterface) StockLevelsClient {
	return &stockLevelsClient{cc}
}

func (c *stockLevelsClient) AddStockLevel(ctx context.Context, in *AddStockLevelRequest, opts ...grpc.CallOption) (*AddStockLevelResponse, error) {
	out := new(AddStockLevelResponse)
	err := c.cc.Invoke(ctx, "/inventory.StockLevels/AddStockLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockLevelsClient) GetStockLevel(ctx context.Context, in *GetStockLevelRequest, opts ...grpc.CallOption) (*GetStockLevelResponse, error) {
	out := new(GetStockLevelResponse)
	err := c.cc.Invoke(ctx, "/inventory.StockLevels/GetStockLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *stockLevelsClient) UpdateStockLevel(ctx context.Context, in *UpdateStockLevelRequest, opts ...grpc.CallOption) (*GetStockLevelResponse, error) {
	out := new(GetStockLevelResponse)
	err := c.cc.Invoke(ctx, "/inventory.StockLevels/UpdateStockLevel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StockLevelsServer is the server API for StockLevels service.
// All implementations must embed UnimplementedStockLevelsServer
// for forward compatibility
type StockLevelsServer interface {
	AddStockLevel(context.Context, *AddStockLevelRequest) (*AddStockLevelResponse, error)
	GetStockLevel(context.Context, *GetStockLevelRequest) (*GetStockLevelResponse, error)
	UpdateStockLevel(context.Context, *UpdateStockLevelRequest) (*GetStockLevelResponse, error)
	mustEmbedUnimplementedStockLevelsServer()
}

// UnimplementedStockLevelsServer must be embedded to have forward compatible implementations.
type UnimplementedStockLevelsServer struct {
}

func (UnimplementedStockLevelsServer) AddStockLevel(context.Context, *AddStockLevelRequest) (*AddStockLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddStockLevel not implemented")
}
func (UnimplementedStockLevelsServer) GetStockLevel(context.Context, *GetStockLevelRequest) (*GetStockLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStockLevel not implemented")
}
func (UnimplementedStockLevelsServer) UpdateStockLevel(context.Context, *UpdateStockLevelRequest) (*GetStockLevelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStockLevel not implemented")
}
func (UnimplementedStockLevelsServer) mustEmbedUnimplementedStockLevelsServer() {}

// UnsafeStockLevelsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StockLevelsServer will
// result in compilation errors.
type UnsafeStockLevelsServer interface {
	mustEmbedUnimplementedStockLevelsServer()
}

func RegisterStockLevelsServer(s grpc.ServiceRegistrar, srv StockLevelsServer) {
	s.RegisterService(&StockLevels_ServiceDesc, srv)
}

func _StockLevels_AddStockLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddStockLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockLevelsServer).AddStockLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.StockLevels/AddStockLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockLevelsServer).AddStockLevel(ctx, req.(*AddStockLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockLevels_GetStockLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStockLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockLevelsServer).GetStockLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.StockLevels/GetStockLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockLevelsServer).GetStockLevel(ctx, req.(*GetStockLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StockLevels_UpdateStockLevel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateStockLevelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StockLevelsServer).UpdateStockLevel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/inventory.StockLevels/UpdateStockLevel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StockLevelsServer).UpdateStockLevel(ctx, req.(*UpdateStockLevelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StockLevels_ServiceDesc is the grpc.ServiceDesc for StockLevels service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StockLevels_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "inventory.StockLevels",
	HandlerType: (*StockLevelsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddStockLevel",
			Handler:    _StockLevels_AddStockLevel_Handler,
		},
		{
			MethodName: "GetStockLevel",
			Handler:    _StockLevels_GetStockLevel_Handler,
		},
		{
			MethodName: "UpdateStockLevel",
			Handler:    _StockLevels_UpdateStockLevel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "stock.proto",
}
