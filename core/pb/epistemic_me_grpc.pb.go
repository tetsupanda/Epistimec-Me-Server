// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v5.28.2
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
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	EpistemicMeService_CreateBelief_FullMethodName        = "/EpistemicMeService/CreateBelief"
	EpistemicMeService_ListBeliefs_FullMethodName         = "/EpistemicMeService/ListBeliefs"
	EpistemicMeService_CreateDialectic_FullMethodName     = "/EpistemicMeService/CreateDialectic"
	EpistemicMeService_ListDialectics_FullMethodName      = "/EpistemicMeService/ListDialectics"
	EpistemicMeService_UpdateDialectic_FullMethodName     = "/EpistemicMeService/UpdateDialectic"
	EpistemicMeService_GetBeliefSystem_FullMethodName     = "/EpistemicMeService/GetBeliefSystem"
	EpistemicMeService_UpdateKeyValueStore_FullMethodName = "/EpistemicMeService/UpdateKeyValueStore"
	EpistemicMeService_CreateSelfModel_FullMethodName     = "/EpistemicMeService/CreateSelfModel"
	EpistemicMeService_GetSelfModel_FullMethodName        = "/EpistemicMeService/GetSelfModel"
	EpistemicMeService_AddPhilosophy_FullMethodName       = "/EpistemicMeService/AddPhilosophy"
	EpistemicMeService_CreateDeveloper_FullMethodName     = "/EpistemicMeService/CreateDeveloper"
	EpistemicMeService_CreateUser_FullMethodName          = "/EpistemicMeService/CreateUser"
	EpistemicMeService_GetDeveloper_FullMethodName        = "/EpistemicMeService/GetDeveloper"
)

// EpistemicMeServiceClient is the client API for EpistemicMeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EpistemicMeServiceClient interface {
	CreateBelief(ctx context.Context, in *CreateBeliefRequest, opts ...grpc.CallOption) (*CreateBeliefResponse, error)
	ListBeliefs(ctx context.Context, in *ListBeliefsRequest, opts ...grpc.CallOption) (*ListBeliefsResponse, error)
	CreateDialectic(ctx context.Context, in *CreateDialecticRequest, opts ...grpc.CallOption) (*CreateDialecticResponse, error)
	ListDialectics(ctx context.Context, in *ListDialecticsRequest, opts ...grpc.CallOption) (*ListDialecticsResponse, error)
	UpdateDialectic(ctx context.Context, in *UpdateDialecticRequest, opts ...grpc.CallOption) (*UpdateDialecticResponse, error)
	GetBeliefSystem(ctx context.Context, in *GetBeliefSystemRequest, opts ...grpc.CallOption) (*GetBeliefSystemResponse, error)
	UpdateKeyValueStore(ctx context.Context, in *UpdateKeyValueStoreRequest, opts ...grpc.CallOption) (*UpdateKeyValueStoreResponse, error)
	CreateSelfModel(ctx context.Context, in *CreateSelfModelRequest, opts ...grpc.CallOption) (*CreateSelfModelResponse, error)
	GetSelfModel(ctx context.Context, in *GetSelfModelRequest, opts ...grpc.CallOption) (*GetSelfModelResponse, error)
	AddPhilosophy(ctx context.Context, in *AddPhilosophyRequest, opts ...grpc.CallOption) (*AddPhilosophyResponse, error)
	CreateDeveloper(ctx context.Context, in *CreateDeveloperRequest, opts ...grpc.CallOption) (*CreateDeveloperResponse, error)
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	GetDeveloper(ctx context.Context, in *GetDeveloperRequest, opts ...grpc.CallOption) (*GetDeveloperResponse, error)
}

type epistemicMeServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEpistemicMeServiceClient(cc grpc.ClientConnInterface) EpistemicMeServiceClient {
	return &epistemicMeServiceClient{cc}
}

func (c *epistemicMeServiceClient) CreateBelief(ctx context.Context, in *CreateBeliefRequest, opts ...grpc.CallOption) (*CreateBeliefResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateBeliefResponse)
	err := c.cc.Invoke(ctx, EpistemicMeService_CreateBelief_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *epistemicMeServiceClient) ListBeliefs(ctx context.Context, in *ListBeliefsRequest, opts ...grpc.CallOption) (*ListBeliefsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListBeliefsResponse)
	err := c.cc.Invoke(ctx, EpistemicMeService_ListBeliefs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *epistemicMeServiceClient) CreateDialectic(ctx context.Context, in *CreateDialecticRequest, opts ...grpc.CallOption) (*CreateDialecticResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDialecticResponse)
	err := c.cc.Invoke(ctx, EpistemicMeService_CreateDialectic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *epistemicMeServiceClient) ListDialectics(ctx context.Context, in *ListDialecticsRequest, opts ...grpc.CallOption) (*ListDialecticsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListDialecticsResponse)
	err := c.cc.Invoke(ctx, EpistemicMeService_ListDialectics_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *epistemicMeServiceClient) UpdateDialectic(ctx context.Context, in *UpdateDialecticRequest, opts ...grpc.CallOption) (*UpdateDialecticResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateDialecticResponse)
	err := c.cc.Invoke(ctx, EpistemicMeService_UpdateDialectic_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *epistemicMeServiceClient) GetBeliefSystem(ctx context.Context, in *GetBeliefSystemRequest, opts ...grpc.CallOption) (*GetBeliefSystemResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBeliefSystemResponse)
	err := c.cc.Invoke(ctx, EpistemicMeService_GetBeliefSystem_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *epistemicMeServiceClient) UpdateKeyValueStore(ctx context.Context, in *UpdateKeyValueStoreRequest, opts ...grpc.CallOption) (*UpdateKeyValueStoreResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateKeyValueStoreResponse)
	err := c.cc.Invoke(ctx, EpistemicMeService_UpdateKeyValueStore_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *epistemicMeServiceClient) CreateSelfModel(ctx context.Context, in *CreateSelfModelRequest, opts ...grpc.CallOption) (*CreateSelfModelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSelfModelResponse)
	err := c.cc.Invoke(ctx, EpistemicMeService_CreateSelfModel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *epistemicMeServiceClient) GetSelfModel(ctx context.Context, in *GetSelfModelRequest, opts ...grpc.CallOption) (*GetSelfModelResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetSelfModelResponse)
	err := c.cc.Invoke(ctx, EpistemicMeService_GetSelfModel_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *epistemicMeServiceClient) AddPhilosophy(ctx context.Context, in *AddPhilosophyRequest, opts ...grpc.CallOption) (*AddPhilosophyResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddPhilosophyResponse)
	err := c.cc.Invoke(ctx, EpistemicMeService_AddPhilosophy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *epistemicMeServiceClient) CreateDeveloper(ctx context.Context, in *CreateDeveloperRequest, opts ...grpc.CallOption) (*CreateDeveloperResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDeveloperResponse)
	err := c.cc.Invoke(ctx, EpistemicMeService_CreateDeveloper_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *epistemicMeServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, EpistemicMeService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *epistemicMeServiceClient) GetDeveloper(ctx context.Context, in *GetDeveloperRequest, opts ...grpc.CallOption) (*GetDeveloperResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDeveloperResponse)
	err := c.cc.Invoke(ctx, EpistemicMeService_GetDeveloper_FullMethodName, in, out, cOpts...)
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
	GetBeliefSystem(context.Context, *GetBeliefSystemRequest) (*GetBeliefSystemResponse, error)
	UpdateKeyValueStore(context.Context, *UpdateKeyValueStoreRequest) (*UpdateKeyValueStoreResponse, error)
	CreateSelfModel(context.Context, *CreateSelfModelRequest) (*CreateSelfModelResponse, error)
	GetSelfModel(context.Context, *GetSelfModelRequest) (*GetSelfModelResponse, error)
	AddPhilosophy(context.Context, *AddPhilosophyRequest) (*AddPhilosophyResponse, error)
	CreateDeveloper(context.Context, *CreateDeveloperRequest) (*CreateDeveloperResponse, error)
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	GetDeveloper(context.Context, *GetDeveloperRequest) (*GetDeveloperResponse, error)
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
func (UnimplementedEpistemicMeServiceServer) GetBeliefSystem(context.Context, *GetBeliefSystemRequest) (*GetBeliefSystemResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBeliefSystem not implemented")
}
func (UnimplementedEpistemicMeServiceServer) UpdateKeyValueStore(context.Context, *UpdateKeyValueStoreRequest) (*UpdateKeyValueStoreResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateKeyValueStore not implemented")
}
func (UnimplementedEpistemicMeServiceServer) CreateSelfModel(context.Context, *CreateSelfModelRequest) (*CreateSelfModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSelfModel not implemented")
}
func (UnimplementedEpistemicMeServiceServer) GetSelfModel(context.Context, *GetSelfModelRequest) (*GetSelfModelResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSelfModel not implemented")
}
func (UnimplementedEpistemicMeServiceServer) AddPhilosophy(context.Context, *AddPhilosophyRequest) (*AddPhilosophyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPhilosophy not implemented")
}
func (UnimplementedEpistemicMeServiceServer) CreateDeveloper(context.Context, *CreateDeveloperRequest) (*CreateDeveloperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDeveloper not implemented")
}
func (UnimplementedEpistemicMeServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedEpistemicMeServiceServer) GetDeveloper(context.Context, *GetDeveloperRequest) (*GetDeveloperResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeveloper not implemented")
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
		FullMethod: EpistemicMeService_CreateBelief_FullMethodName,
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
		FullMethod: EpistemicMeService_ListBeliefs_FullMethodName,
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
		FullMethod: EpistemicMeService_CreateDialectic_FullMethodName,
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
		FullMethod: EpistemicMeService_ListDialectics_FullMethodName,
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
		FullMethod: EpistemicMeService_UpdateDialectic_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpistemicMeServiceServer).UpdateDialectic(ctx, req.(*UpdateDialecticRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpistemicMeService_GetBeliefSystem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBeliefSystemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpistemicMeServiceServer).GetBeliefSystem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EpistemicMeService_GetBeliefSystem_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpistemicMeServiceServer).GetBeliefSystem(ctx, req.(*GetBeliefSystemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpistemicMeService_UpdateKeyValueStore_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateKeyValueStoreRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpistemicMeServiceServer).UpdateKeyValueStore(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EpistemicMeService_UpdateKeyValueStore_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpistemicMeServiceServer).UpdateKeyValueStore(ctx, req.(*UpdateKeyValueStoreRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpistemicMeService_CreateSelfModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSelfModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpistemicMeServiceServer).CreateSelfModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EpistemicMeService_CreateSelfModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpistemicMeServiceServer).CreateSelfModel(ctx, req.(*CreateSelfModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpistemicMeService_GetSelfModel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSelfModelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpistemicMeServiceServer).GetSelfModel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EpistemicMeService_GetSelfModel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpistemicMeServiceServer).GetSelfModel(ctx, req.(*GetSelfModelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpistemicMeService_AddPhilosophy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddPhilosophyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpistemicMeServiceServer).AddPhilosophy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EpistemicMeService_AddPhilosophy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpistemicMeServiceServer).AddPhilosophy(ctx, req.(*AddPhilosophyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpistemicMeService_CreateDeveloper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDeveloperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpistemicMeServiceServer).CreateDeveloper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EpistemicMeService_CreateDeveloper_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpistemicMeServiceServer).CreateDeveloper(ctx, req.(*CreateDeveloperRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpistemicMeService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpistemicMeServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EpistemicMeService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpistemicMeServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EpistemicMeService_GetDeveloper_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeveloperRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EpistemicMeServiceServer).GetDeveloper(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: EpistemicMeService_GetDeveloper_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EpistemicMeServiceServer).GetDeveloper(ctx, req.(*GetDeveloperRequest))
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
		{
			MethodName: "GetBeliefSystem",
			Handler:    _EpistemicMeService_GetBeliefSystem_Handler,
		},
		{
			MethodName: "UpdateKeyValueStore",
			Handler:    _EpistemicMeService_UpdateKeyValueStore_Handler,
		},
		{
			MethodName: "CreateSelfModel",
			Handler:    _EpistemicMeService_CreateSelfModel_Handler,
		},
		{
			MethodName: "GetSelfModel",
			Handler:    _EpistemicMeService_GetSelfModel_Handler,
		},
		{
			MethodName: "AddPhilosophy",
			Handler:    _EpistemicMeService_AddPhilosophy_Handler,
		},
		{
			MethodName: "CreateDeveloper",
			Handler:    _EpistemicMeService_CreateDeveloper_Handler,
		},
		{
			MethodName: "CreateUser",
			Handler:    _EpistemicMeService_CreateUser_Handler,
		},
		{
			MethodName: "GetDeveloper",
			Handler:    _EpistemicMeService_GetDeveloper_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/epistemic_me.proto",
}
