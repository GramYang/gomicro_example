// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: auth.proto

package mu_micro_book_srv_auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Service service

type Service interface {
	MakeAccessToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	DelUserAccessToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetCachedAccessToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
}

type service struct {
	c    client.Client
	name string
}

func NewService(name string, c client.Client) Service {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "mu.micro.book.srv.auth"
	}
	return &service{
		c:    c,
		name: name,
	}
}

func (c *service) MakeAccessToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Service.MakeAccessToken", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) DelUserAccessToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Service.DelUserAccessToken", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *service) GetCachedAccessToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Service.GetCachedAccessToken", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Service service

type ServiceHandler interface {
	MakeAccessToken(context.Context, *Request, *Response) error
	DelUserAccessToken(context.Context, *Request, *Response) error
	GetCachedAccessToken(context.Context, *Request, *Response) error
}

func RegisterServiceHandler(s server.Server, hdlr ServiceHandler, opts ...server.HandlerOption) error {
	type service interface {
		MakeAccessToken(ctx context.Context, in *Request, out *Response) error
		DelUserAccessToken(ctx context.Context, in *Request, out *Response) error
		GetCachedAccessToken(ctx context.Context, in *Request, out *Response) error
	}
	type Service struct {
		service
	}
	h := &serviceHandler{hdlr}
	return s.Handle(s.NewHandler(&Service{h}, opts...))
}

type serviceHandler struct {
	ServiceHandler
}

func (h *serviceHandler) MakeAccessToken(ctx context.Context, in *Request, out *Response) error {
	return h.ServiceHandler.MakeAccessToken(ctx, in, out)
}

func (h *serviceHandler) DelUserAccessToken(ctx context.Context, in *Request, out *Response) error {
	return h.ServiceHandler.DelUserAccessToken(ctx, in, out)
}

func (h *serviceHandler) GetCachedAccessToken(ctx context.Context, in *Request, out *Response) error {
	return h.ServiceHandler.GetCachedAccessToken(ctx, in, out)
}
