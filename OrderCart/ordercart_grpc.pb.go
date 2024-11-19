// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.19.6
// source: OrderCart/ordercart.proto

package __

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
	OrderCartService_AddProductToCart_FullMethodName         = "/ordercart.OrderCartService/AddProductToCart"
	OrderCartService_GetCartItems_FullMethodName             = "/ordercart.OrderCartService/GetCartItems"
	OrderCartService_IncrementProductQuantity_FullMethodName = "/ordercart.OrderCartService/IncrementProductQuantity"
	OrderCartService_DecrementProductQuantity_FullMethodName = "/ordercart.OrderCartService/DecrementProductQuantity"
	OrderCartService_RemoveProductFromCart_FullMethodName    = "/ordercart.OrderCartService/RemoveProductFromCart"
	OrderCartService_ClearCart_FullMethodName                = "/ordercart.OrderCartService/ClearCart"
	OrderCartService_PlaceOrderByRestID_FullMethodName       = "/ordercart.OrderCartService/PlaceOrderByRestID"
	OrderCartService_GetOrderDetailsAll_FullMethodName       = "/ordercart.OrderCartService/GetOrderDetailsAll"
	OrderCartService_GetOrderDetailsByID_FullMethodName      = "/ordercart.OrderCartService/GetOrderDetailsByID"
)

// OrderCartServiceClient is the client API for OrderCartService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Define the OrderCartService
type OrderCartServiceClient interface {
	// Cart Operations
	AddProductToCart(ctx context.Context, in *AddProductToCartRequest, opts ...grpc.CallOption) (*AddProductToCartResponse, error)
	GetCartItems(ctx context.Context, in *GetCartItemsRequest, opts ...grpc.CallOption) (*GetCartItemsResponse, error)
	IncrementProductQuantity(ctx context.Context, in *IncrementProductQuantityRequest, opts ...grpc.CallOption) (*IncrementProductQuantityResponse, error)
	DecrementProductQuantity(ctx context.Context, in *DecrementProductQuantityRequest, opts ...grpc.CallOption) (*DecrementProductQuantityResponse, error)
	RemoveProductFromCart(ctx context.Context, in *RemoveProductFromCartRequest, opts ...grpc.CallOption) (*RemoveProductFromCartResponse, error)
	ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*ClearCartResponse, error)
	// Order Operations
	PlaceOrderByRestID(ctx context.Context, in *PlaceOrderByRestIDRequest, opts ...grpc.CallOption) (*PlaceOrderByRestIDResponse, error)
	GetOrderDetailsAll(ctx context.Context, in *GetOrderDetailsAllRequest, opts ...grpc.CallOption) (*GetOrderDetailsAllResponse, error)
	GetOrderDetailsByID(ctx context.Context, in *GetOrderDetailsByIDRequest, opts ...grpc.CallOption) (*GetOrderDetailsByIDResponse, error)
}

type orderCartServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderCartServiceClient(cc grpc.ClientConnInterface) OrderCartServiceClient {
	return &orderCartServiceClient{cc}
}

func (c *orderCartServiceClient) AddProductToCart(ctx context.Context, in *AddProductToCartRequest, opts ...grpc.CallOption) (*AddProductToCartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddProductToCartResponse)
	err := c.cc.Invoke(ctx, OrderCartService_AddProductToCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderCartServiceClient) GetCartItems(ctx context.Context, in *GetCartItemsRequest, opts ...grpc.CallOption) (*GetCartItemsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetCartItemsResponse)
	err := c.cc.Invoke(ctx, OrderCartService_GetCartItems_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderCartServiceClient) IncrementProductQuantity(ctx context.Context, in *IncrementProductQuantityRequest, opts ...grpc.CallOption) (*IncrementProductQuantityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(IncrementProductQuantityResponse)
	err := c.cc.Invoke(ctx, OrderCartService_IncrementProductQuantity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderCartServiceClient) DecrementProductQuantity(ctx context.Context, in *DecrementProductQuantityRequest, opts ...grpc.CallOption) (*DecrementProductQuantityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DecrementProductQuantityResponse)
	err := c.cc.Invoke(ctx, OrderCartService_DecrementProductQuantity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderCartServiceClient) RemoveProductFromCart(ctx context.Context, in *RemoveProductFromCartRequest, opts ...grpc.CallOption) (*RemoveProductFromCartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RemoveProductFromCartResponse)
	err := c.cc.Invoke(ctx, OrderCartService_RemoveProductFromCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderCartServiceClient) ClearCart(ctx context.Context, in *ClearCartRequest, opts ...grpc.CallOption) (*ClearCartResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ClearCartResponse)
	err := c.cc.Invoke(ctx, OrderCartService_ClearCart_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderCartServiceClient) PlaceOrderByRestID(ctx context.Context, in *PlaceOrderByRestIDRequest, opts ...grpc.CallOption) (*PlaceOrderByRestIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PlaceOrderByRestIDResponse)
	err := c.cc.Invoke(ctx, OrderCartService_PlaceOrderByRestID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderCartServiceClient) GetOrderDetailsAll(ctx context.Context, in *GetOrderDetailsAllRequest, opts ...grpc.CallOption) (*GetOrderDetailsAllResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderDetailsAllResponse)
	err := c.cc.Invoke(ctx, OrderCartService_GetOrderDetailsAll_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *orderCartServiceClient) GetOrderDetailsByID(ctx context.Context, in *GetOrderDetailsByIDRequest, opts ...grpc.CallOption) (*GetOrderDetailsByIDResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetOrderDetailsByIDResponse)
	err := c.cc.Invoke(ctx, OrderCartService_GetOrderDetailsByID_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderCartServiceServer is the server API for OrderCartService service.
// All implementations must embed UnimplementedOrderCartServiceServer
// for forward compatibility.
//
// Define the OrderCartService
type OrderCartServiceServer interface {
	// Cart Operations
	AddProductToCart(context.Context, *AddProductToCartRequest) (*AddProductToCartResponse, error)
	GetCartItems(context.Context, *GetCartItemsRequest) (*GetCartItemsResponse, error)
	IncrementProductQuantity(context.Context, *IncrementProductQuantityRequest) (*IncrementProductQuantityResponse, error)
	DecrementProductQuantity(context.Context, *DecrementProductQuantityRequest) (*DecrementProductQuantityResponse, error)
	RemoveProductFromCart(context.Context, *RemoveProductFromCartRequest) (*RemoveProductFromCartResponse, error)
	ClearCart(context.Context, *ClearCartRequest) (*ClearCartResponse, error)
	// Order Operations
	PlaceOrderByRestID(context.Context, *PlaceOrderByRestIDRequest) (*PlaceOrderByRestIDResponse, error)
	GetOrderDetailsAll(context.Context, *GetOrderDetailsAllRequest) (*GetOrderDetailsAllResponse, error)
	GetOrderDetailsByID(context.Context, *GetOrderDetailsByIDRequest) (*GetOrderDetailsByIDResponse, error)
	mustEmbedUnimplementedOrderCartServiceServer()
}

// UnimplementedOrderCartServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedOrderCartServiceServer struct{}

func (UnimplementedOrderCartServiceServer) AddProductToCart(context.Context, *AddProductToCartRequest) (*AddProductToCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddProductToCart not implemented")
}
func (UnimplementedOrderCartServiceServer) GetCartItems(context.Context, *GetCartItemsRequest) (*GetCartItemsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCartItems not implemented")
}
func (UnimplementedOrderCartServiceServer) IncrementProductQuantity(context.Context, *IncrementProductQuantityRequest) (*IncrementProductQuantityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IncrementProductQuantity not implemented")
}
func (UnimplementedOrderCartServiceServer) DecrementProductQuantity(context.Context, *DecrementProductQuantityRequest) (*DecrementProductQuantityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DecrementProductQuantity not implemented")
}
func (UnimplementedOrderCartServiceServer) RemoveProductFromCart(context.Context, *RemoveProductFromCartRequest) (*RemoveProductFromCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveProductFromCart not implemented")
}
func (UnimplementedOrderCartServiceServer) ClearCart(context.Context, *ClearCartRequest) (*ClearCartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClearCart not implemented")
}
func (UnimplementedOrderCartServiceServer) PlaceOrderByRestID(context.Context, *PlaceOrderByRestIDRequest) (*PlaceOrderByRestIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceOrderByRestID not implemented")
}
func (UnimplementedOrderCartServiceServer) GetOrderDetailsAll(context.Context, *GetOrderDetailsAllRequest) (*GetOrderDetailsAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderDetailsAll not implemented")
}
func (UnimplementedOrderCartServiceServer) GetOrderDetailsByID(context.Context, *GetOrderDetailsByIDRequest) (*GetOrderDetailsByIDResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrderDetailsByID not implemented")
}
func (UnimplementedOrderCartServiceServer) mustEmbedUnimplementedOrderCartServiceServer() {}
func (UnimplementedOrderCartServiceServer) testEmbeddedByValue()                          {}

// UnsafeOrderCartServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderCartServiceServer will
// result in compilation errors.
type UnsafeOrderCartServiceServer interface {
	mustEmbedUnimplementedOrderCartServiceServer()
}

func RegisterOrderCartServiceServer(s grpc.ServiceRegistrar, srv OrderCartServiceServer) {
	// If the following call pancis, it indicates UnimplementedOrderCartServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&OrderCartService_ServiceDesc, srv)
}

func _OrderCartService_AddProductToCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddProductToCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderCartServiceServer).AddProductToCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderCartService_AddProductToCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderCartServiceServer).AddProductToCart(ctx, req.(*AddProductToCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderCartService_GetCartItems_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCartItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderCartServiceServer).GetCartItems(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderCartService_GetCartItems_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderCartServiceServer).GetCartItems(ctx, req.(*GetCartItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderCartService_IncrementProductQuantity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IncrementProductQuantityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderCartServiceServer).IncrementProductQuantity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderCartService_IncrementProductQuantity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderCartServiceServer).IncrementProductQuantity(ctx, req.(*IncrementProductQuantityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderCartService_DecrementProductQuantity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecrementProductQuantityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderCartServiceServer).DecrementProductQuantity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderCartService_DecrementProductQuantity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderCartServiceServer).DecrementProductQuantity(ctx, req.(*DecrementProductQuantityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderCartService_RemoveProductFromCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveProductFromCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderCartServiceServer).RemoveProductFromCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderCartService_RemoveProductFromCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderCartServiceServer).RemoveProductFromCart(ctx, req.(*RemoveProductFromCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderCartService_ClearCart_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClearCartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderCartServiceServer).ClearCart(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderCartService_ClearCart_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderCartServiceServer).ClearCart(ctx, req.(*ClearCartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderCartService_PlaceOrderByRestID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceOrderByRestIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderCartServiceServer).PlaceOrderByRestID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderCartService_PlaceOrderByRestID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderCartServiceServer).PlaceOrderByRestID(ctx, req.(*PlaceOrderByRestIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderCartService_GetOrderDetailsAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderDetailsAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderCartServiceServer).GetOrderDetailsAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderCartService_GetOrderDetailsAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderCartServiceServer).GetOrderDetailsAll(ctx, req.(*GetOrderDetailsAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _OrderCartService_GetOrderDetailsByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOrderDetailsByIDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderCartServiceServer).GetOrderDetailsByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: OrderCartService_GetOrderDetailsByID_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderCartServiceServer).GetOrderDetailsByID(ctx, req.(*GetOrderDetailsByIDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderCartService_ServiceDesc is the grpc.ServiceDesc for OrderCartService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderCartService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ordercart.OrderCartService",
	HandlerType: (*OrderCartServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddProductToCart",
			Handler:    _OrderCartService_AddProductToCart_Handler,
		},
		{
			MethodName: "GetCartItems",
			Handler:    _OrderCartService_GetCartItems_Handler,
		},
		{
			MethodName: "IncrementProductQuantity",
			Handler:    _OrderCartService_IncrementProductQuantity_Handler,
		},
		{
			MethodName: "DecrementProductQuantity",
			Handler:    _OrderCartService_DecrementProductQuantity_Handler,
		},
		{
			MethodName: "RemoveProductFromCart",
			Handler:    _OrderCartService_RemoveProductFromCart_Handler,
		},
		{
			MethodName: "ClearCart",
			Handler:    _OrderCartService_ClearCart_Handler,
		},
		{
			MethodName: "PlaceOrderByRestID",
			Handler:    _OrderCartService_PlaceOrderByRestID_Handler,
		},
		{
			MethodName: "GetOrderDetailsAll",
			Handler:    _OrderCartService_GetOrderDetailsAll_Handler,
		},
		{
			MethodName: "GetOrderDetailsByID",
			Handler:    _OrderCartService_GetOrderDetailsByID_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "OrderCart/ordercart.proto",
}
