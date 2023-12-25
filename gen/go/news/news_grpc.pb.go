// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.14.0
// source: news/news.proto

package newsv1

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
	News_NewsSave_FullMethodName   = "/news.News/NewsSave"
	News_NewsEdit_FullMethodName   = "/news.News/NewsEdit"
	News_NewsDelete_FullMethodName = "/news.News/NewsDelete"
)

// NewsClient is the client API for News service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NewsClient interface {
	NewsSave(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error)
	NewsEdit(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error)
	NewsDelete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
}

type newsClient struct {
	cc grpc.ClientConnInterface
}

func NewNewsClient(cc grpc.ClientConnInterface) NewsClient {
	return &newsClient{cc}
}

func (c *newsClient) NewsSave(ctx context.Context, in *SaveRequest, opts ...grpc.CallOption) (*SaveResponse, error) {
	out := new(SaveResponse)
	err := c.cc.Invoke(ctx, News_NewsSave_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) NewsEdit(ctx context.Context, in *EditRequest, opts ...grpc.CallOption) (*EditResponse, error) {
	out := new(EditResponse)
	err := c.cc.Invoke(ctx, News_NewsEdit_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *newsClient) NewsDelete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, News_NewsDelete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NewsServer is the server API for News service.
// All implementations must embed UnimplementedNewsServer
// for forward compatibility
type NewsServer interface {
	NewsSave(context.Context, *SaveRequest) (*SaveResponse, error)
	NewsEdit(context.Context, *EditRequest) (*EditResponse, error)
	NewsDelete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	mustEmbedUnimplementedNewsServer()
}

// UnimplementedNewsServer must be embedded to have forward compatible implementations.
type UnimplementedNewsServer struct {
}

func (UnimplementedNewsServer) NewsSave(context.Context, *SaveRequest) (*SaveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewsSave not implemented")
}
func (UnimplementedNewsServer) NewsEdit(context.Context, *EditRequest) (*EditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewsEdit not implemented")
}
func (UnimplementedNewsServer) NewsDelete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewsDelete not implemented")
}
func (UnimplementedNewsServer) mustEmbedUnimplementedNewsServer() {}

// UnsafeNewsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NewsServer will
// result in compilation errors.
type UnsafeNewsServer interface {
	mustEmbedUnimplementedNewsServer()
}

func RegisterNewsServer(s grpc.ServiceRegistrar, srv NewsServer) {
	s.RegisterService(&News_ServiceDesc, srv)
}

func _News_NewsSave_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).NewsSave(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_NewsSave_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).NewsSave(ctx, req.(*SaveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_NewsEdit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).NewsEdit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_NewsEdit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).NewsEdit(ctx, req.(*EditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _News_NewsDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NewsServer).NewsDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: News_NewsDelete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NewsServer).NewsDelete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// News_ServiceDesc is the grpc.ServiceDesc for News service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var News_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "news.News",
	HandlerType: (*NewsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewsSave",
			Handler:    _News_NewsSave_Handler,
		},
		{
			MethodName: "NewsEdit",
			Handler:    _News_NewsEdit_Handler,
		},
		{
			MethodName: "NewsDelete",
			Handler:    _News_NewsDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "news/news.proto",
}
