// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: petstore/v1/pet.proto

package v1

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

// PetServiceClient is the client API for PetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PetServiceClient interface {
	UploadImage(ctx context.Context, in *UploadImageRequest, opts ...grpc.CallOption) (*Response, error)
	CreatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Response, error)
	UpdatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Response, error)
	DeletePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Response, error)
	FindPet(ctx context.Context, in *FindPetRequest, opts ...grpc.CallOption) (*FindPetResponse, error)
}

type petServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPetServiceClient(cc grpc.ClientConnInterface) PetServiceClient {
	return &petServiceClient{cc}
}

func (c *petServiceClient) UploadImage(ctx context.Context, in *UploadImageRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/petstore.v1.PetService/UploadImage", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petServiceClient) CreatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/petstore.v1.PetService/CreatePet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petServiceClient) UpdatePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/petstore.v1.PetService/UpdatePet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petServiceClient) DeletePet(ctx context.Context, in *Pet, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/petstore.v1.PetService/DeletePet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petServiceClient) FindPet(ctx context.Context, in *FindPetRequest, opts ...grpc.CallOption) (*FindPetResponse, error) {
	out := new(FindPetResponse)
	err := c.cc.Invoke(ctx, "/petstore.v1.PetService/FindPet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetServiceServer is the server API for PetService service.
// All implementations must embed UnimplementedPetServiceServer
// for forward compatibility
type PetServiceServer interface {
	UploadImage(context.Context, *UploadImageRequest) (*Response, error)
	CreatePet(context.Context, *Pet) (*Response, error)
	UpdatePet(context.Context, *Pet) (*Response, error)
	DeletePet(context.Context, *Pet) (*Response, error)
	FindPet(context.Context, *FindPetRequest) (*FindPetResponse, error)
	mustEmbedUnimplementedPetServiceServer()
}

// UnimplementedPetServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPetServiceServer struct {
}

func (UnimplementedPetServiceServer) UploadImage(context.Context, *UploadImageRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadImage not implemented")
}
func (UnimplementedPetServiceServer) CreatePet(context.Context, *Pet) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePet not implemented")
}
func (UnimplementedPetServiceServer) UpdatePet(context.Context, *Pet) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePet not implemented")
}
func (UnimplementedPetServiceServer) DeletePet(context.Context, *Pet) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePet not implemented")
}
func (UnimplementedPetServiceServer) FindPet(context.Context, *FindPetRequest) (*FindPetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPet not implemented")
}
func (UnimplementedPetServiceServer) mustEmbedUnimplementedPetServiceServer() {}

// UnsafePetServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PetServiceServer will
// result in compilation errors.
type UnsafePetServiceServer interface {
	mustEmbedUnimplementedPetServiceServer()
}

func RegisterPetServiceServer(s grpc.ServiceRegistrar, srv PetServiceServer) {
	s.RegisterService(&PetService_ServiceDesc, srv)
}

func _PetService_UploadImage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServiceServer).UploadImage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.v1.PetService/UploadImage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServiceServer).UploadImage(ctx, req.(*UploadImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetService_CreatePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServiceServer).CreatePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.v1.PetService/CreatePet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServiceServer).CreatePet(ctx, req.(*Pet))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetService_UpdatePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServiceServer).UpdatePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.v1.PetService/UpdatePet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServiceServer).UpdatePet(ctx, req.(*Pet))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetService_DeletePet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pet)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServiceServer).DeletePet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.v1.PetService/DeletePet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServiceServer).DeletePet(ctx, req.(*Pet))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetService_FindPet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetServiceServer).FindPet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/petstore.v1.PetService/FindPet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetServiceServer).FindPet(ctx, req.(*FindPetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PetService_ServiceDesc is the grpc.ServiceDesc for PetService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PetService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "petstore.v1.PetService",
	HandlerType: (*PetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadImage",
			Handler:    _PetService_UploadImage_Handler,
		},
		{
			MethodName: "CreatePet",
			Handler:    _PetService_CreatePet_Handler,
		},
		{
			MethodName: "UpdatePet",
			Handler:    _PetService_UpdatePet_Handler,
		},
		{
			MethodName: "DeletePet",
			Handler:    _PetService_DeletePet_Handler,
		},
		{
			MethodName: "FindPet",
			Handler:    _PetService_FindPet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "petstore/v1/pet.proto",
}