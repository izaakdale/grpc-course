// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.4
// source: calculator/proto/calculator.proto

package grpc_course

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

// CalculatorServiceClient is the client API for CalculatorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CalculatorServiceClient interface {
	RequestAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_RequestAverageClient, error)
	GetMax(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_GetMaxClient, error)
	ManyPrimes(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (CalculatorService_ManyPrimesClient, error)
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error)
	Sqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error)
}

type calculatorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCalculatorServiceClient(cc grpc.ClientConnInterface) CalculatorServiceClient {
	return &calculatorServiceClient{cc}
}

func (c *calculatorServiceClient) RequestAverage(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_RequestAverageClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[0], "/calculator.CalculatorService/RequestAverage", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceRequestAverageClient{stream}
	return x, nil
}

type CalculatorService_RequestAverageClient interface {
	Send(*AvgRequest) error
	CloseAndRecv() (*AvgResponse, error)
	grpc.ClientStream
}

type calculatorServiceRequestAverageClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceRequestAverageClient) Send(m *AvgRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceRequestAverageClient) CloseAndRecv() (*AvgResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(AvgResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) GetMax(ctx context.Context, opts ...grpc.CallOption) (CalculatorService_GetMaxClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[1], "/calculator.CalculatorService/GetMax", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceGetMaxClient{stream}
	return x, nil
}

type CalculatorService_GetMaxClient interface {
	Send(*MaxRequest) error
	Recv() (*MaxReponse, error)
	grpc.ClientStream
}

type calculatorServiceGetMaxClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceGetMaxClient) Send(m *MaxRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *calculatorServiceGetMaxClient) Recv() (*MaxReponse, error) {
	m := new(MaxReponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) ManyPrimes(ctx context.Context, in *PrimeRequest, opts ...grpc.CallOption) (CalculatorService_ManyPrimesClient, error) {
	stream, err := c.cc.NewStream(ctx, &CalculatorService_ServiceDesc.Streams[2], "/calculator.CalculatorService/ManyPrimes", opts...)
	if err != nil {
		return nil, err
	}
	x := &calculatorServiceManyPrimesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CalculatorService_ManyPrimesClient interface {
	Recv() (*PrimeResponse, error)
	grpc.ClientStream
}

type calculatorServiceManyPrimesClient struct {
	grpc.ClientStream
}

func (x *calculatorServiceManyPrimesClient) Recv() (*PrimeResponse, error) {
	m := new(PrimeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *calculatorServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumResponse, error) {
	out := new(SumResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *calculatorServiceClient) Sqrt(ctx context.Context, in *SqrtRequest, opts ...grpc.CallOption) (*SqrtResponse, error) {
	out := new(SqrtResponse)
	err := c.cc.Invoke(ctx, "/calculator.CalculatorService/Sqrt", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CalculatorServiceServer is the server API for CalculatorService service.
// All implementations must embed UnimplementedCalculatorServiceServer
// for forward compatibility
type CalculatorServiceServer interface {
	RequestAverage(CalculatorService_RequestAverageServer) error
	GetMax(CalculatorService_GetMaxServer) error
	ManyPrimes(*PrimeRequest, CalculatorService_ManyPrimesServer) error
	Sum(context.Context, *SumRequest) (*SumResponse, error)
	Sqrt(context.Context, *SqrtRequest) (*SqrtResponse, error)
	mustEmbedUnimplementedCalculatorServiceServer()
}

// UnimplementedCalculatorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCalculatorServiceServer struct {
}

func (UnimplementedCalculatorServiceServer) RequestAverage(CalculatorService_RequestAverageServer) error {
	return status.Errorf(codes.Unimplemented, "method RequestAverage not implemented")
}
func (UnimplementedCalculatorServiceServer) GetMax(CalculatorService_GetMaxServer) error {
	return status.Errorf(codes.Unimplemented, "method GetMax not implemented")
}
func (UnimplementedCalculatorServiceServer) ManyPrimes(*PrimeRequest, CalculatorService_ManyPrimesServer) error {
	return status.Errorf(codes.Unimplemented, "method ManyPrimes not implemented")
}
func (UnimplementedCalculatorServiceServer) Sum(context.Context, *SumRequest) (*SumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}
func (UnimplementedCalculatorServiceServer) Sqrt(context.Context, *SqrtRequest) (*SqrtResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sqrt not implemented")
}
func (UnimplementedCalculatorServiceServer) mustEmbedUnimplementedCalculatorServiceServer() {}

// UnsafeCalculatorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CalculatorServiceServer will
// result in compilation errors.
type UnsafeCalculatorServiceServer interface {
	mustEmbedUnimplementedCalculatorServiceServer()
}

func RegisterCalculatorServiceServer(s grpc.ServiceRegistrar, srv CalculatorServiceServer) {
	s.RegisterService(&CalculatorService_ServiceDesc, srv)
}

func _CalculatorService_RequestAverage_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).RequestAverage(&calculatorServiceRequestAverageServer{stream})
}

type CalculatorService_RequestAverageServer interface {
	SendAndClose(*AvgResponse) error
	Recv() (*AvgRequest, error)
	grpc.ServerStream
}

type calculatorServiceRequestAverageServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceRequestAverageServer) SendAndClose(m *AvgResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceRequestAverageServer) Recv() (*AvgRequest, error) {
	m := new(AvgRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_GetMax_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(CalculatorServiceServer).GetMax(&calculatorServiceGetMaxServer{stream})
}

type CalculatorService_GetMaxServer interface {
	Send(*MaxReponse) error
	Recv() (*MaxRequest, error)
	grpc.ServerStream
}

type calculatorServiceGetMaxServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceGetMaxServer) Send(m *MaxReponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *calculatorServiceGetMaxServer) Recv() (*MaxRequest, error) {
	m := new(MaxRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _CalculatorService_ManyPrimes_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(PrimeRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CalculatorServiceServer).ManyPrimes(m, &calculatorServiceManyPrimesServer{stream})
}

type CalculatorService_ManyPrimesServer interface {
	Send(*PrimeResponse) error
	grpc.ServerStream
}

type calculatorServiceManyPrimesServer struct {
	grpc.ServerStream
}

func (x *calculatorServiceManyPrimesServer) Send(m *PrimeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _CalculatorService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CalculatorService_Sqrt_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SqrtRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CalculatorServiceServer).Sqrt(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/calculator.CalculatorService/Sqrt",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CalculatorServiceServer).Sqrt(ctx, req.(*SqrtRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CalculatorService_ServiceDesc is the grpc.ServiceDesc for CalculatorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CalculatorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "calculator.CalculatorService",
	HandlerType: (*CalculatorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _CalculatorService_Sum_Handler,
		},
		{
			MethodName: "Sqrt",
			Handler:    _CalculatorService_Sqrt_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestAverage",
			Handler:       _CalculatorService_RequestAverage_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetMax",
			Handler:       _CalculatorService_GetMax_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "ManyPrimes",
			Handler:       _CalculatorService_ManyPrimes_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "calculator/proto/calculator.proto",
}
