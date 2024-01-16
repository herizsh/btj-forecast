// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: weather.proto

package pb2

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

// WeatherServiceClient is the client API for WeatherService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WeatherServiceClient interface {
	StreamForecasts(ctx context.Context, in *ForecastRequest, opts ...grpc.CallOption) (WeatherService_StreamForecastsClient, error)
}

type weatherServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWeatherServiceClient(cc grpc.ClientConnInterface) WeatherServiceClient {
	return &weatherServiceClient{cc}
}

func (c *weatherServiceClient) StreamForecasts(ctx context.Context, in *ForecastRequest, opts ...grpc.CallOption) (WeatherService_StreamForecastsClient, error) {
	stream, err := c.cc.NewStream(ctx, &WeatherService_ServiceDesc.Streams[0], "/weather.WeatherService/StreamForecasts", opts...)
	if err != nil {
		return nil, err
	}
	x := &weatherServiceStreamForecastsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WeatherService_StreamForecastsClient interface {
	Recv() (*ForecastResponse, error)
	grpc.ClientStream
}

type weatherServiceStreamForecastsClient struct {
	grpc.ClientStream
}

func (x *weatherServiceStreamForecastsClient) Recv() (*ForecastResponse, error) {
	m := new(ForecastResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WeatherServiceServer is the server API for WeatherService service.
// All implementations must embed UnimplementedWeatherServiceServer
// for forward compatibility
type WeatherServiceServer interface {
	StreamForecasts(*ForecastRequest, WeatherService_StreamForecastsServer) error
	mustEmbedUnimplementedWeatherServiceServer()
}

// UnimplementedWeatherServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWeatherServiceServer struct {
}

func (UnimplementedWeatherServiceServer) StreamForecasts(*ForecastRequest, WeatherService_StreamForecastsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamForecasts not implemented")
}
func (UnimplementedWeatherServiceServer) mustEmbedUnimplementedWeatherServiceServer() {}

// UnsafeWeatherServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WeatherServiceServer will
// result in compilation errors.
type UnsafeWeatherServiceServer interface {
	mustEmbedUnimplementedWeatherServiceServer()
}

func RegisterWeatherServiceServer(s grpc.ServiceRegistrar, srv WeatherServiceServer) {
	s.RegisterService(&WeatherService_ServiceDesc, srv)
}

func _WeatherService_StreamForecasts_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ForecastRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WeatherServiceServer).StreamForecasts(m, &weatherServiceStreamForecastsServer{stream})
}

type WeatherService_StreamForecastsServer interface {
	Send(*ForecastResponse) error
	grpc.ServerStream
}

type weatherServiceStreamForecastsServer struct {
	grpc.ServerStream
}

func (x *weatherServiceStreamForecastsServer) Send(m *ForecastResponse) error {
	return x.ServerStream.SendMsg(m)
}

// WeatherService_ServiceDesc is the grpc.ServiceDesc for WeatherService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WeatherService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "weather.WeatherService",
	HandlerType: (*WeatherServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamForecasts",
			Handler:       _WeatherService_StreamForecasts_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "weather.proto",
}