// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.2
// source: api/proto/order.proto

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

const (
	OrderService_CreateOrder_FullMethodName     = "/order.OrderService/CreateOrder"
	OrderService_ReadOrder_FullMethodName       = "/order.OrderService/ReadOrder"
	OrderService_UpdateOrder_FullMethodName     = "/order.OrderService/UpdateOrder"
	OrderService_DeleteOrder_FullMethodName     = "/order.OrderService/DeleteOrder"
	OrderService_CreateOrderItem_FullMethodName = "/order.OrderService/CreateOrderItem"
	OrderService_ReadOrderItem_FullMethodName   = "/order.OrderService/ReadOrderItem"
	OrderService_UpdateOrderItem_FullMethodName = "/order.OrderService/UpdateOrderItem"
	OrderService_DeleteOrderItem_FullMethodName = "/order.OrderService/DeleteOrderItem"
	OrderService_DecomposeOrder_FullMethodName  = "/order.OrderService/DecomposeOrder"
)

// OrderServiceClient is the client API for OrderService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderServiceClient interface {
	CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error)
	ReadOrder(ctx context.Context, in *ReadOrderRequest, opts ...grpc.CallOption) (*ReadOrderResponse, error)
	UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error)
	DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderResponse, error)
	CreateOrderItem(ctx context.Context, in *CreateOrderItemRequest, opts ...grpc.CallOption) (*CreateOrderItemResponse, error)
	ReadOrderItem(ctx context.Context, in *ReadOrderItemRequest, opts ...grpc.CallOption) (*ReadOrderItemResponse, error)
	UpdateOrderItem(ctx context.Context, in *UpdateOrderItemRequest, opts ...grpc.CallOption) (*UpdateOrderItemResponse, error)
	DeleteOrderItem(ctx context.Context, in *DeleteOrderItemRequest, opts ...grpc.CallOption) (*DeleteOrderItemResponse, error)
	DecomposeOrder(ctx context.Context, in *DecomposeOrderRequest, opts ...grpc.CallOption) (*DecomposeOrderResponse, error)
}

type orderServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderServiceClient(cc grpc.ClientConnInterface) OrderServiceClient {
	return &orderServiceClient{cc}
}

func (c *orderServiceClient) CreateOrder(ctx context.Context, in *CreateOrderRequest, opts ...grpc.CallOption) (*CreateOrderResponse, error) {
	out := new(CreateOrderResponse)
	err := c.cc.Invoke(ctx, OrderService_CreateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ReadOrder(ctx context.Context, in *ReadOrderRequest, opts ...grpc.CallOption) (*ReadOrderResponse, error) {
	out := new(ReadOrderResponse)
	err := c.cc.Invoke(ctx, OrderService_ReadOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrder(ctx context.Context, in *UpdateOrderRequest, opts ...grpc.CallOption) (*UpdateOrderResponse, error) {
	out := new(UpdateOrderResponse)
	err := c.cc.Invoke(ctx, OrderService_UpdateOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) DeleteOrder(ctx context.Context, in *DeleteOrderRequest, opts ...grpc.CallOption) (*DeleteOrderResponse, error) {
	out := new(DeleteOrderResponse)
	err := c.cc.Invoke(ctx, OrderService_DeleteOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) CreateOrderItem(ctx context.Context, in *CreateOrderItemRequest, opts ...grpc.CallOption) (*CreateOrderItemResponse, error) {
	out := new(CreateOrderItemResponse)
	err := c.cc.Invoke(ctx, OrderService_CreateOrderItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) ReadOrderItem(ctx context.Context, in *ReadOrderItemRequest, opts ...grpc.CallOption) (*ReadOrderItemResponse, error) {
	out := new(ReadOrderItemResponse)
	err := c.cc.Invoke(ctx, OrderService_ReadOrderItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) UpdateOrderItem(ctx context.Context, in *UpdateOrderItemRequest, opts ...grpc.CallOption) (*UpdateOrderItemResponse, error) {
	out := new(UpdateOrderItemResponse)
	err := c.cc.Invoke(ctx, OrderService_UpdateOrderItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) DeleteOrderItem(ctx context.Context, in *DeleteOrderItemRequest, opts ...grpc.CallOption) (*DeleteOrderItemResponse, error) {
	out := new(DeleteOrderItemResponse)
	err := c.cc.Invoke(ctx, OrderService_DeleteOrderItem_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderServiceClient) DecomposeOrder(ctx context.Context, in *DecomposeOrderRequest, opts ...grpc.CallOption) (*DecomposeOrderResponse, error) {
	out := new(DecomposeOrderResponse)
	err := c.cc.Invoke(ctx, OrderService_DecomposeOrder_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServiceServer is the server API for OrderService service.
// All implementations must embed UnimplementedOrderServiceServer
// for forward compatibility
type OrderServiceServer interface {
	CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error)
	ReadOrder(context.Context, *ReadOrderRequest) (*ReadOrderResponse, error)
	UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error)
	DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderResponse, error)
	CreateOrderItem(context.Context, *CreateOrderItemRequest) (*CreateOrderItemResponse, error)
	ReadOrderItem(context.Context, *ReadOrderItemRequest) (*ReadOrderItemResponse, error)
	UpdateOrderItem(context.Context, *UpdateOrderItemRequest) (*UpdateOrderItemResponse, error)
	DeleteOrderItem(context.Context, *DeleteOrderItemRequest) (*DeleteOrderItemResponse, error)
	DecomposeOrder(context.Context, *DecomposeOrderRequest) (*DecomposeOrderResponse, error)
	mustEmbedUnimplementedOrderServiceServer()
}

// UnimplementedOrderServiceServer must be embedded to have forward compatible implementations.
type UnimplementedOrderServiceServer struct {
}

func (UnimplementedOrderServiceServer) CreateOrder(context.Context, *CreateOrderRequest) (*CreateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrder not implemented")
}
func (UnimplementedOrderServiceServer) ReadOrder(context.Context, *ReadOrderRequest) (*ReadOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadOrder not implemented")
}
func (UnimplementedOrderServiceServer) UpdateOrder(context.Context, *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrder not implemented")
}
func (UnimplementedOrderServiceServer) DeleteOrder(context.Context, *DeleteOrderRequest) (*DeleteOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrder not implemented")
}
func (UnimplementedOrderServiceServer) CreateOrderItem(context.Context, *CreateOrderItemRequest) (*CreateOrderItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateOrderItem not implemented")
}
func (UnimplementedOrderServiceServer) ReadOrderItem(context.Context, *ReadOrderItemRequest) (*ReadOrderItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReadOrderItem not implemented")
}
func (UnimplementedOrderServiceServer) UpdateOrderItem(context.Context, *UpdateOrderItemRequest) (*UpdateOrderItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateOrderItem not implemented")
}
func (UnimplementedOrderServiceServer) DeleteOrderItem(context.Context, *DeleteOrderItemRequest) (*DeleteOrderItemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteOrderItem not implemented")
}
func (UnimplementedOrderServiceServer) DecomposeOrder(context.Context, *DecomposeOrderRequest) (*DecomposeOrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecomposeOrder not implemented")
}
func (UnimplementedOrderServiceServer) mustEmbedUnimplementedOrderServiceServer() {}

// UnsafeOrderServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderServiceServer will
// result in compilation errors.
type UnsafeOrderServiceServer interface {
	mustEmbedUnimplementedOrderServiceServer()
}

func RegisterOrderServiceServer(s grpc.ServiceRegistrar, srv OrderServiceServer) {
	s.RegisterService(&OrderService_ServiceDesc, srv)
}

func _OrderService_CreateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrder(ctx, req.(*CreateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ReadOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ReadOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_ReadOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ReadOrder(ctx, req.(*ReadOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_UpdateOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrder(ctx, req.(*UpdateOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_DeleteOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).DeleteOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_DeleteOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).DeleteOrder(ctx, req.(*DeleteOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_CreateOrderItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateOrderItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).CreateOrderItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_CreateOrderItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).CreateOrderItem(ctx, req.(*CreateOrderItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_ReadOrderItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadOrderItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).ReadOrderItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_ReadOrderItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).ReadOrderItem(ctx, req.(*ReadOrderItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_UpdateOrderItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateOrderItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).UpdateOrderItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_UpdateOrderItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).UpdateOrderItem(ctx, req.(*UpdateOrderItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_DeleteOrderItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteOrderItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).DeleteOrderItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_DeleteOrderItem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).DeleteOrderItem(ctx, req.(*DeleteOrderItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderService_DecomposeOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecomposeOrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServiceServer).DecomposeOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderService_DecomposeOrder_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServiceServer).DecomposeOrder(ctx, req.(*DecomposeOrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderService_ServiceDesc is the grpc.ServiceDesc for OrderService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "order.OrderService",
	HandlerType: (*OrderServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateOrder",
			Handler:    _OrderService_CreateOrder_Handler,
		},
		{
			MethodName: "ReadOrder",
			Handler:    _OrderService_ReadOrder_Handler,
		},
		{
			MethodName: "UpdateOrder",
			Handler:    _OrderService_UpdateOrder_Handler,
		},
		{
			MethodName: "DeleteOrder",
			Handler:    _OrderService_DeleteOrder_Handler,
		},
		{
			MethodName: "CreateOrderItem",
			Handler:    _OrderService_CreateOrderItem_Handler,
		},
		{
			MethodName: "ReadOrderItem",
			Handler:    _OrderService_ReadOrderItem_Handler,
		},
		{
			MethodName: "UpdateOrderItem",
			Handler:    _OrderService_UpdateOrderItem_Handler,
		},
		{
			MethodName: "DeleteOrderItem",
			Handler:    _OrderService_DeleteOrderItem_Handler,
		},
		{
			MethodName: "DecomposeOrder",
			Handler:    _OrderService_DecomposeOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/order.proto",
}
