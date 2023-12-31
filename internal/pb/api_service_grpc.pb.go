// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.4
// source: api_service.proto

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
	App_UpdateProfile_FullMethodName                 = "/pb.App/UpdateProfile"
	App_BuildCircle_FullMethodName                   = "/pb.App/BuildCircle"
	App_ShareMood_FullMethodName                     = "/pb.App/ShareMood"
	App_GetCircleJoinRequestsWithUser_FullMethodName = "/pb.App/GetCircleJoinRequestsWithUser"
)

// AppClient is the client API for App service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AppClient interface {
	UpdateProfile(ctx context.Context, opts ...grpc.CallOption) (App_UpdateProfileClient, error)
	BuildCircle(ctx context.Context, opts ...grpc.CallOption) (App_BuildCircleClient, error)
	ShareMood(ctx context.Context, opts ...grpc.CallOption) (App_ShareMoodClient, error)
	GetCircleJoinRequestsWithUser(ctx context.Context, in *GetCircleJoinRequestsWithUserRequest, opts ...grpc.CallOption) (*CircleJoinRequestsWithUserResponse, error)
}

type appClient struct {
	cc grpc.ClientConnInterface
}

func NewAppClient(cc grpc.ClientConnInterface) AppClient {
	return &appClient{cc}
}

func (c *appClient) UpdateProfile(ctx context.Context, opts ...grpc.CallOption) (App_UpdateProfileClient, error) {
	stream, err := c.cc.NewStream(ctx, &App_ServiceDesc.Streams[0], App_UpdateProfile_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &appUpdateProfileClient{stream}
	return x, nil
}

type App_UpdateProfileClient interface {
	Send(*UpdateProfileRequest) error
	CloseAndRecv() (*UpdateProfileResponse, error)
	grpc.ClientStream
}

type appUpdateProfileClient struct {
	grpc.ClientStream
}

func (x *appUpdateProfileClient) Send(m *UpdateProfileRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *appUpdateProfileClient) CloseAndRecv() (*UpdateProfileResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(UpdateProfileResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *appClient) BuildCircle(ctx context.Context, opts ...grpc.CallOption) (App_BuildCircleClient, error) {
	stream, err := c.cc.NewStream(ctx, &App_ServiceDesc.Streams[1], App_BuildCircle_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &appBuildCircleClient{stream}
	return x, nil
}

type App_BuildCircleClient interface {
	Send(*BuildCircleRequest) error
	CloseAndRecv() (*BuildCircleResponse, error)
	grpc.ClientStream
}

type appBuildCircleClient struct {
	grpc.ClientStream
}

func (x *appBuildCircleClient) Send(m *BuildCircleRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *appBuildCircleClient) CloseAndRecv() (*BuildCircleResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BuildCircleResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *appClient) ShareMood(ctx context.Context, opts ...grpc.CallOption) (App_ShareMoodClient, error) {
	stream, err := c.cc.NewStream(ctx, &App_ServiceDesc.Streams[2], App_ShareMood_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &appShareMoodClient{stream}
	return x, nil
}

type App_ShareMoodClient interface {
	Send(*ShareMoodRequest) error
	CloseAndRecv() (*ShareMoodResponse, error)
	grpc.ClientStream
}

type appShareMoodClient struct {
	grpc.ClientStream
}

func (x *appShareMoodClient) Send(m *ShareMoodRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *appShareMoodClient) CloseAndRecv() (*ShareMoodResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ShareMoodResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *appClient) GetCircleJoinRequestsWithUser(ctx context.Context, in *GetCircleJoinRequestsWithUserRequest, opts ...grpc.CallOption) (*CircleJoinRequestsWithUserResponse, error) {
	out := new(CircleJoinRequestsWithUserResponse)
	err := c.cc.Invoke(ctx, App_GetCircleJoinRequestsWithUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AppServer is the server API for App service.
// All implementations must embed UnimplementedAppServer
// for forward compatibility
type AppServer interface {
	UpdateProfile(App_UpdateProfileServer) error
	BuildCircle(App_BuildCircleServer) error
	ShareMood(App_ShareMoodServer) error
	GetCircleJoinRequestsWithUser(context.Context, *GetCircleJoinRequestsWithUserRequest) (*CircleJoinRequestsWithUserResponse, error)
	mustEmbedUnimplementedAppServer()
}

// UnimplementedAppServer must be embedded to have forward compatible implementations.
type UnimplementedAppServer struct {
}

func (UnimplementedAppServer) UpdateProfile(App_UpdateProfileServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedAppServer) BuildCircle(App_BuildCircleServer) error {
	return status.Errorf(codes.Unimplemented, "method BuildCircle not implemented")
}
func (UnimplementedAppServer) ShareMood(App_ShareMoodServer) error {
	return status.Errorf(codes.Unimplemented, "method ShareMood not implemented")
}
func (UnimplementedAppServer) GetCircleJoinRequestsWithUser(context.Context, *GetCircleJoinRequestsWithUserRequest) (*CircleJoinRequestsWithUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCircleJoinRequestsWithUser not implemented")
}
func (UnimplementedAppServer) mustEmbedUnimplementedAppServer() {}

// UnsafeAppServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AppServer will
// result in compilation errors.
type UnsafeAppServer interface {
	mustEmbedUnimplementedAppServer()
}

func RegisterAppServer(s grpc.ServiceRegistrar, srv AppServer) {
	s.RegisterService(&App_ServiceDesc, srv)
}

func _App_UpdateProfile_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AppServer).UpdateProfile(&appUpdateProfileServer{stream})
}

type App_UpdateProfileServer interface {
	SendAndClose(*UpdateProfileResponse) error
	Recv() (*UpdateProfileRequest, error)
	grpc.ServerStream
}

type appUpdateProfileServer struct {
	grpc.ServerStream
}

func (x *appUpdateProfileServer) SendAndClose(m *UpdateProfileResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *appUpdateProfileServer) Recv() (*UpdateProfileRequest, error) {
	m := new(UpdateProfileRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _App_BuildCircle_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AppServer).BuildCircle(&appBuildCircleServer{stream})
}

type App_BuildCircleServer interface {
	SendAndClose(*BuildCircleResponse) error
	Recv() (*BuildCircleRequest, error)
	grpc.ServerStream
}

type appBuildCircleServer struct {
	grpc.ServerStream
}

func (x *appBuildCircleServer) SendAndClose(m *BuildCircleResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *appBuildCircleServer) Recv() (*BuildCircleRequest, error) {
	m := new(BuildCircleRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _App_ShareMood_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(AppServer).ShareMood(&appShareMoodServer{stream})
}

type App_ShareMoodServer interface {
	SendAndClose(*ShareMoodResponse) error
	Recv() (*ShareMoodRequest, error)
	grpc.ServerStream
}

type appShareMoodServer struct {
	grpc.ServerStream
}

func (x *appShareMoodServer) SendAndClose(m *ShareMoodResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *appShareMoodServer) Recv() (*ShareMoodRequest, error) {
	m := new(ShareMoodRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _App_GetCircleJoinRequestsWithUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCircleJoinRequestsWithUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AppServer).GetCircleJoinRequestsWithUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: App_GetCircleJoinRequestsWithUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AppServer).GetCircleJoinRequestsWithUser(ctx, req.(*GetCircleJoinRequestsWithUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// App_ServiceDesc is the grpc.ServiceDesc for App service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var App_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.App",
	HandlerType: (*AppServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCircleJoinRequestsWithUser",
			Handler:    _App_GetCircleJoinRequestsWithUser_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdateProfile",
			Handler:       _App_UpdateProfile_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BuildCircle",
			Handler:       _App_BuildCircle_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "ShareMood",
			Handler:       _App_ShareMood_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api_service.proto",
}
