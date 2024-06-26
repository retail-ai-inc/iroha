// protoc --go_out=. --go-grpc_out=. decrypt.proto

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.27.1
// source: decrypt.proto

package decrypt

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	Decrypt_Decrypt_FullMethodName = "/decrypt.Decrypt/Decrypt"
)

// DecryptClient is the client API for Decrypt service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DecryptClient interface {
	Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error)
}

type decryptClient struct {
	cc grpc.ClientConnInterface
}

func NewDecryptClient(cc grpc.ClientConnInterface) DecryptClient {
	return &decryptClient{cc}
}

func (c *decryptClient) Decrypt(ctx context.Context, in *DecryptRequest, opts ...grpc.CallOption) (*DecryptResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DecryptResponse)
	err := c.cc.Invoke(ctx, Decrypt_Decrypt_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DecryptServer is the server API for Decrypt service.
// All implementations must embed UnimplementedDecryptServer
// for forward compatibility
type DecryptServer interface {
	Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error)
	mustEmbedUnimplementedDecryptServer()
}

// UnimplementedDecryptServer must be embedded to have forward compatible implementations.
type UnimplementedDecryptServer struct {
}

func (UnimplementedDecryptServer) Decrypt(context.Context, *DecryptRequest) (*DecryptResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Decrypt not implemented")
}
func (UnimplementedDecryptServer) mustEmbedUnimplementedDecryptServer() {}

// UnsafeDecryptServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DecryptServer will
// result in compilation errors.
type UnsafeDecryptServer interface {
	mustEmbedUnimplementedDecryptServer()
}

func RegisterDecryptServer(s grpc.ServiceRegistrar, srv DecryptServer) {
	s.RegisterService(&Decrypt_ServiceDesc, srv)
}

func _Decrypt_Decrypt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DecryptRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DecryptServer).Decrypt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Decrypt_Decrypt_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DecryptServer).Decrypt(ctx, req.(*DecryptRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Decrypt_ServiceDesc is the grpc.ServiceDesc for Decrypt service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Decrypt_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "decrypt.Decrypt",
	HandlerType: (*DecryptServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Decrypt",
			Handler:    _Decrypt_Decrypt_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "decrypt.proto",
}