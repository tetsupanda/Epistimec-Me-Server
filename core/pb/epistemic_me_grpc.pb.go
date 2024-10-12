// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: proto/epistemic_me.proto

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

// EpistemicMeServiceClient is the client API for EpistemicMeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EpistemicMeServiceClient interface {
	CreateBelief(ctx context.Context, in *CreateBeliefRequest, opts ...grpc.CallOption) (*CreateBeliefResponse, error)
	ListBeliefs(ctx context.Context, in *ListBeliefsRequest, opts ...grpc.CallOption) (*ListBeliefsResponse, error)
	CreateDialectic(ctx context.Context, in *CreateDialecticRequest, opts ...grpc.CallOption) (*CreateDialecticResponse, error)
	ListDialectics(ctx context.Context, in *ListDialecticsRequest, opts ...grpc.CallOption) (*ListDialecticsResponse, error)
	UpdateDialectic(ctx context.Context, in *UpdateDialecticRequest, opts ...grpc.CallOption) (*UpdateDialecticResponse, error)
}

type epistemicMeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEpistemicMeServiceClient(cc grpc.ClientConnInterface) EpistemicMeServiceClient {
	return &epistemicMeServiceClient{cc}
}

func (c *epistemicMeServiceClient) CreateBelief(ctx context.Context, in *CreateBeliefRequest, opts ...grpc.CallOption) (*CreateBeliefResponse, error) {
	out := new(CreateBeliefResponse)
	err := c.cc.Invoke(ctx, "/EpistemicMeService/CreateBelief", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *epistemicMeServiceClient) ListBeliefs(ctx context.Context, in *ListBeliefsRequest, opts ...grpc.CallOption) (*ListBeliefsResponse, error) {
	out := new(ListBeliefsResponse)
	err := c.cc.Invoke(ctx, "/EpistemicMeService/ListBeliefs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *epistemicMeServiceClient) CreateDialectic(ctx context.Context, in *CreateDialecticRequest, opts ...grpc.CallOption) (*CreateDialecticResponse, error) {
	out := new(CreateDialecticResponse)
	err := c.cc.Invoke(ctx, "/EpistemicMeService/CreateDialectic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *epistemicMeServiceClient) ListDialectics(ctx context.Context, in *ListDialecticsRequest, opts ...grpc.CallOption) (*ListDialecticsResponse, error) {
	out := new(ListDialecticsResponse)
	err := c.cc.Invoke(ctx, "/EpistemicMeService/ListDialectics", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *epistemicMeServiceClient) UpdateDialectic(ctx context.Context, in *UpdateDialecticRequest, opts ...grpc.CallOption) (*UpdateDialecticResponse, error) {
	out := new(UpdateDialecticResponse)
	err := c.cc.Invoke(ctx, "/EpistemicMeService/UpdateDialectic", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EpistemicMeServiceServer is the server API for EpistemicMeService service.
// All implementations must embed UnimplementedEpistemicMeServiceServer
// for forward compatibility
type EpistemicMeServiceServer interface {
	CreateBelief(context.Context, *CreateBeliefRequest) (*CreateBeliefResponse, error)
	ListBeliefs(context.Context, *ListBeliefsRequest) (*ListBeliefsResponse, error)
	CreateDialectic(context.Context, *CreateDialecticRequest) (*CreateDialecticResponse, error)
	ListDialectics(context.Context, *ListDialecticsRequest) (*ListDialecticsResponse, error)
	UpdateDialectic(context.Context, *UpdateDialecticRequest) (*UpdateDialecticResponse, error)
	mustEmbedUnimplementedEpistemicMeServiceServer()
}

// UnimplementedEpistemicMeServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEpistemicMeServiceServer struct {
}

func (UnimplementedEpistemicMeServiceServer) CreateBelief(context.Context, *CreateBeliefRequest) (*CreateBeliefResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBelief not implemented")
}
func (UnimplementedEpistemicMeServiceServer) ListBeliefs(context.Context, *ListBeliefsRequest) (*ListBeliefsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBeliefs not implemented")
}
func (UnimplementedEpistemicMeServiceServer) CreateDialectic(context.Context, *CreateDialecticRequest) (*CreateDialecticResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDialectic not implemented")
}
func (UnimplementedEpistemicMeServiceServer) ListDialectics(context.Context, *ListDialecticsRequest) (*ListDialecticsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDialectics not implemented")
}
func (UnimplementedEpistemicMeServiceServer) UpdateDialectic(context.Context, *UpdateDialecticRequest) (*UpdateDialecticResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateDialectic not implemented")
}
func (UnimplementedEpistemicMeServiceServer) mustEmbedUnimplementedEpistemicMeServiceServer() {}

// UnsafeEpistemicMeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EpistemicMeServiceServer will
// result in compilation errors.
type UnsafeEpistemicMeServiceServer interface {
	mustEmbedUnimplementedEpistemicMeServiceServer()
}

func RegisterEpistemicMeServiceServer(s grpc.ServiceRegistrar, srv EpistemicMeServiceServer) {
	s.RegisterService(&EpistemicMeService_ServiceDesc, srv)
}

func _EpistemicMeService_CreateBelief_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBeliefRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpistemicMeServiceServer).CreateBelief(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EpistemicMeService/CreateBelief",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpistemicMeServiceServer).CreateBelief(ctx, req.(*CreateBeliefRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpistemicMeService_ListBeliefs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBeliefsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpistemicMeServiceServer).ListBeliefs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EpistemicMeService/ListBeliefs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpistemicMeServiceServer).ListBeliefs(ctx, req.(*ListBeliefsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpistemicMeService_CreateDialectic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDialecticRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpistemicMeServiceServer).CreateDialectic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EpistemicMeService/CreateDialectic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpistemicMeServiceServer).CreateDialectic(ctx, req.(*CreateDialecticRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpistemicMeService_ListDialectics_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDialecticsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpistemicMeServiceServer).ListDialectics(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EpistemicMeService/ListDialectics",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpistemicMeServiceServer).ListDialectics(ctx, req.(*ListDialecticsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpistemicMeService_UpdateDialectic_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateDialecticRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpistemicMeServiceServer).UpdateDialectic(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/EpistemicMeService/UpdateDialectic",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpistemicMeServiceServer).UpdateDialectic(ctx, req.(*UpdateDialecticRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EpistemicMeService_ServiceDesc is the grpc.ServiceDesc for EpistemicMeService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EpistemicMeService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "EpistemicMeService",
	HandlerType: (*EpistemicMeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBelief",
			Handler:    _EpistemicMeService_CreateBelief_Handler,
		},
		{
			MethodName: "ListBeliefs",
			Handler:    _EpistemicMeService_ListBeliefs_Handler,
		},
		{
			MethodName: "CreateDialectic",
			Handler:    _EpistemicMeService_CreateDialectic_Handler,
		},
		{
			MethodName: "ListDialectics",
			Handler:    _EpistemicMeService_ListDialectics_Handler,
		},
		{
			MethodName: "UpdateDialectic",
			Handler:    _EpistemicMeService_UpdateDialectic_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/epistemic_me.proto",
}
