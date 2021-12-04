// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc

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

// CustomerServiceClient is the client API for CustomerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CustomerServiceClient interface {
	AddNewCustomer(ctx context.Context, in *AddNewCustomerRequest, opts ...grpc.CallOption) (*AddNewCustomerResponses, error)
	RemoveCustomer(ctx context.Context, in *RemoveCustomerRequest, opts ...grpc.CallOption) (*RemoveCustomerResponse, error)
	UpdateCustomerData(ctx context.Context, in *UpdateCustomerDataRequest, opts ...grpc.CallOption) (*UpdateCustomerDataResponses, error)
	UpdateCustomerProfilePicture(ctx context.Context, in *UpdateCustomerProfilePictureRequest, opts ...grpc.CallOption) (*UpdateCustomerProfilePictureResponses, error)
	// Customer contacts
	AddCustomerContact(ctx context.Context, in *AddCustomerContactRequest, opts ...grpc.CallOption) (*AddCustomerContactResponses, error)
	RemoveCustomerContact(ctx context.Context, in *RemoveCustomerContactRequest, opts ...grpc.CallOption) (*RemoveCustomerContactResponses, error)
	UpdateCustomerContact(ctx context.Context, in *UpdateCustomerContactRequest, opts ...grpc.CallOption) (*UpdateCustomerContactResponses, error)
}

type customerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCustomerServiceClient(cc grpc.ClientConnInterface) CustomerServiceClient {
	return &customerServiceClient{cc}
}

func (c *customerServiceClient) AddNewCustomer(ctx context.Context, in *AddNewCustomerRequest, opts ...grpc.CallOption) (*AddNewCustomerResponses, error) {
	out := new(AddNewCustomerResponses)
	err := c.cc.Invoke(ctx, "/connectors.CustomerService/AddNewCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) RemoveCustomer(ctx context.Context, in *RemoveCustomerRequest, opts ...grpc.CallOption) (*RemoveCustomerResponse, error) {
	out := new(RemoveCustomerResponse)
	err := c.cc.Invoke(ctx, "/connectors.CustomerService/RemoveCustomer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) UpdateCustomerData(ctx context.Context, in *UpdateCustomerDataRequest, opts ...grpc.CallOption) (*UpdateCustomerDataResponses, error) {
	out := new(UpdateCustomerDataResponses)
	err := c.cc.Invoke(ctx, "/connectors.CustomerService/UpdateCustomerData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) UpdateCustomerProfilePicture(ctx context.Context, in *UpdateCustomerProfilePictureRequest, opts ...grpc.CallOption) (*UpdateCustomerProfilePictureResponses, error) {
	out := new(UpdateCustomerProfilePictureResponses)
	err := c.cc.Invoke(ctx, "/connectors.CustomerService/UpdateCustomerProfilePicture", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) AddCustomerContact(ctx context.Context, in *AddCustomerContactRequest, opts ...grpc.CallOption) (*AddCustomerContactResponses, error) {
	out := new(AddCustomerContactResponses)
	err := c.cc.Invoke(ctx, "/connectors.CustomerService/AddCustomerContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) RemoveCustomerContact(ctx context.Context, in *RemoveCustomerContactRequest, opts ...grpc.CallOption) (*RemoveCustomerContactResponses, error) {
	out := new(RemoveCustomerContactResponses)
	err := c.cc.Invoke(ctx, "/connectors.CustomerService/RemoveCustomerContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *customerServiceClient) UpdateCustomerContact(ctx context.Context, in *UpdateCustomerContactRequest, opts ...grpc.CallOption) (*UpdateCustomerContactResponses, error) {
	out := new(UpdateCustomerContactResponses)
	err := c.cc.Invoke(ctx, "/connectors.CustomerService/UpdateCustomerContact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CustomerServiceServer is the server API for CustomerService service.
// All implementations must embed UnimplementedCustomerServiceServer
// for forward compatibility
type CustomerServiceServer interface {
	AddNewCustomer(context.Context, *AddNewCustomerRequest) (*AddNewCustomerResponses, error)
	RemoveCustomer(context.Context, *RemoveCustomerRequest) (*RemoveCustomerResponse, error)
	UpdateCustomerData(context.Context, *UpdateCustomerDataRequest) (*UpdateCustomerDataResponses, error)
	UpdateCustomerProfilePicture(context.Context, *UpdateCustomerProfilePictureRequest) (*UpdateCustomerProfilePictureResponses, error)
	// Customer contacts
	AddCustomerContact(context.Context, *AddCustomerContactRequest) (*AddCustomerContactResponses, error)
	RemoveCustomerContact(context.Context, *RemoveCustomerContactRequest) (*RemoveCustomerContactResponses, error)
	UpdateCustomerContact(context.Context, *UpdateCustomerContactRequest) (*UpdateCustomerContactResponses, error)
	mustEmbedUnimplementedCustomerServiceServer()
}

// UnimplementedCustomerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCustomerServiceServer struct {
}

func (UnimplementedCustomerServiceServer) AddNewCustomer(context.Context, *AddNewCustomerRequest) (*AddNewCustomerResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNewCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) RemoveCustomer(context.Context, *RemoveCustomerRequest) (*RemoveCustomerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCustomer not implemented")
}
func (UnimplementedCustomerServiceServer) UpdateCustomerData(context.Context, *UpdateCustomerDataRequest) (*UpdateCustomerDataResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomerData not implemented")
}
func (UnimplementedCustomerServiceServer) UpdateCustomerProfilePicture(context.Context, *UpdateCustomerProfilePictureRequest) (*UpdateCustomerProfilePictureResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomerProfilePicture not implemented")
}
func (UnimplementedCustomerServiceServer) AddCustomerContact(context.Context, *AddCustomerContactRequest) (*AddCustomerContactResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddCustomerContact not implemented")
}
func (UnimplementedCustomerServiceServer) RemoveCustomerContact(context.Context, *RemoveCustomerContactRequest) (*RemoveCustomerContactResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveCustomerContact not implemented")
}
func (UnimplementedCustomerServiceServer) UpdateCustomerContact(context.Context, *UpdateCustomerContactRequest) (*UpdateCustomerContactResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCustomerContact not implemented")
}
func (UnimplementedCustomerServiceServer) mustEmbedUnimplementedCustomerServiceServer() {}

// UnsafeCustomerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CustomerServiceServer will
// result in compilation errors.
type UnsafeCustomerServiceServer interface {
	mustEmbedUnimplementedCustomerServiceServer()
}

func RegisterCustomerServiceServer(s grpc.ServiceRegistrar, srv CustomerServiceServer) {
	s.RegisterService(&CustomerService_ServiceDesc, srv)
}

func _CustomerService_AddNewCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNewCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).AddNewCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.CustomerService/AddNewCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).AddNewCustomer(ctx, req.(*AddNewCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_RemoveCustomer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCustomerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).RemoveCustomer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.CustomerService/RemoveCustomer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).RemoveCustomer(ctx, req.(*RemoveCustomerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_UpdateCustomerData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustomerDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).UpdateCustomerData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.CustomerService/UpdateCustomerData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).UpdateCustomerData(ctx, req.(*UpdateCustomerDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_UpdateCustomerProfilePicture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustomerProfilePictureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).UpdateCustomerProfilePicture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.CustomerService/UpdateCustomerProfilePicture",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).UpdateCustomerProfilePicture(ctx, req.(*UpdateCustomerProfilePictureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_AddCustomerContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddCustomerContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).AddCustomerContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.CustomerService/AddCustomerContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).AddCustomerContact(ctx, req.(*AddCustomerContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_RemoveCustomerContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveCustomerContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).RemoveCustomerContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.CustomerService/RemoveCustomerContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).RemoveCustomerContact(ctx, req.(*RemoveCustomerContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CustomerService_UpdateCustomerContact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateCustomerContactRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CustomerServiceServer).UpdateCustomerContact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/connectors.CustomerService/UpdateCustomerContact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CustomerServiceServer).UpdateCustomerContact(ctx, req.(*UpdateCustomerContactRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CustomerService_ServiceDesc is the grpc.ServiceDesc for CustomerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CustomerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "connectors.CustomerService",
	HandlerType: (*CustomerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddNewCustomer",
			Handler:    _CustomerService_AddNewCustomer_Handler,
		},
		{
			MethodName: "RemoveCustomer",
			Handler:    _CustomerService_RemoveCustomer_Handler,
		},
		{
			MethodName: "UpdateCustomerData",
			Handler:    _CustomerService_UpdateCustomerData_Handler,
		},
		{
			MethodName: "UpdateCustomerProfilePicture",
			Handler:    _CustomerService_UpdateCustomerProfilePicture_Handler,
		},
		{
			MethodName: "AddCustomerContact",
			Handler:    _CustomerService_AddCustomerContact_Handler,
		},
		{
			MethodName: "RemoveCustomerContact",
			Handler:    _CustomerService_RemoveCustomerContact_Handler,
		},
		{
			MethodName: "UpdateCustomerContact",
			Handler:    _CustomerService_UpdateCustomerContact_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "customer.proto",
}
