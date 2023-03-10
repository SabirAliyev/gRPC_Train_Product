// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: productmgmt/productmgmt.proto

package go_productmgmt_grpc

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

// ProductManagementClient is the client API for ProductManagement service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProductManagementClient interface {
	CreateProduct(ctx context.Context, in *NewProduct, opts ...grpc.CallOption) (*Product, error)
	GetProduct(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Product, error)
}

type productManagementClient struct {
	cc grpc.ClientConnInterface
}

func NewProductManagementClient(cc grpc.ClientConnInterface) ProductManagementClient {
	return &productManagementClient{cc}
}

func (c *productManagementClient) CreateProduct(ctx context.Context, in *NewProduct, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/productmgmt.ProductManagement/CreateProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productManagementClient) GetProduct(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Product, error) {
	out := new(Product)
	err := c.cc.Invoke(ctx, "/productmgmt.ProductManagement/GetProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProductManagementServer is the server API for ProductManagement service.
// All implementations must embed UnimplementedProductManagementServer
// for forward compatibility
type ProductManagementServer interface {
	CreateProduct(context.Context, *NewProduct) (*Product, error)
	GetProduct(context.Context, *Id) (*Product, error)
	mustEmbedUnimplementedProductManagementServer()
}

// UnimplementedProductManagementServer must be embedded to have forward compatible implementations.
type UnimplementedProductManagementServer struct {
}

func (UnimplementedProductManagementServer) CreateProduct(context.Context, *NewProduct) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateProduct not implemented")
}
func (UnimplementedProductManagementServer) GetProduct(context.Context, *Id) (*Product, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduct not implemented")
}
func (UnimplementedProductManagementServer) mustEmbedUnimplementedProductManagementServer() {}

// UnsafeProductManagementServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProductManagementServer will
// result in compilation errors.
type UnsafeProductManagementServer interface {
	mustEmbedUnimplementedProductManagementServer()
}

func RegisterProductManagementServer(s grpc.ServiceRegistrar, srv ProductManagementServer) {
	s.RegisterService(&ProductManagement_ServiceDesc, srv)
}

func _ProductManagement_CreateProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewProduct)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManagementServer).CreateProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productmgmt.ProductManagement/CreateProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManagementServer).CreateProduct(ctx, req.(*NewProduct))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProductManagement_GetProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Id)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProductManagementServer).GetProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/productmgmt.ProductManagement/GetProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProductManagementServer).GetProduct(ctx, req.(*Id))
	}
	return interceptor(ctx, in, info, handler)
}

// ProductManagement_ServiceDesc is the grpc.ServiceDesc for ProductManagement service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProductManagement_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "productmgmt.ProductManagement",
	HandlerType: (*ProductManagementServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateProduct",
			Handler:    _ProductManagement_CreateProduct_Handler,
		},
		{
			MethodName: "GetProduct",
			Handler:    _ProductManagement_GetProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "productmgmt/productmgmt.proto",
}
