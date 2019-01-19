// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: publish.proto

/*
Package publishen_publish_service is a generated protocol buffer package.

It is generated from these files:
	publish.proto

It has these top-level messages:
	Request
	Publish
	Location
	Attribute
	Response
*/
package publishen_publish_service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for PublishService service

type PublishService interface {
	CreatePublish(ctx context.Context, in *Publish, opts ...client.CallOption) (*Response, error)
	ListOfPublishes(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error)
	GetPublish(ctx context.Context, in *Publish, opts ...client.CallOption) (*Response, error)
	UpdatePublish(ctx context.Context, in *Publish, opts ...client.CallOption) (*Response, error)
	DeletePublish(ctx context.Context, in *Publish, opts ...client.CallOption) (*Response, error)
}

type publishService struct {
	c    client.Client
	name string
}

func NewPublishService(name string, c client.Client) PublishService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "publishen.publish.service"
	}
	return &publishService{
		c:    c,
		name: name,
	}
}

func (c *publishService) CreatePublish(ctx context.Context, in *Publish, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PublishService.CreatePublish", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishService) ListOfPublishes(ctx context.Context, in *Request, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PublishService.ListOfPublishes", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishService) GetPublish(ctx context.Context, in *Publish, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PublishService.GetPublish", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishService) UpdatePublish(ctx context.Context, in *Publish, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PublishService.UpdatePublish", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *publishService) DeletePublish(ctx context.Context, in *Publish, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "PublishService.DeletePublish", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PublishService service

type PublishServiceHandler interface {
	CreatePublish(context.Context, *Publish, *Response) error
	ListOfPublishes(context.Context, *Request, *Response) error
	GetPublish(context.Context, *Publish, *Response) error
	UpdatePublish(context.Context, *Publish, *Response) error
	DeletePublish(context.Context, *Publish, *Response) error
}

func RegisterPublishServiceHandler(s server.Server, hdlr PublishServiceHandler, opts ...server.HandlerOption) error {
	type publishService interface {
		CreatePublish(ctx context.Context, in *Publish, out *Response) error
		ListOfPublishes(ctx context.Context, in *Request, out *Response) error
		GetPublish(ctx context.Context, in *Publish, out *Response) error
		UpdatePublish(ctx context.Context, in *Publish, out *Response) error
		DeletePublish(ctx context.Context, in *Publish, out *Response) error
	}
	type PublishService struct {
		publishService
	}
	h := &publishServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PublishService{h}, opts...))
}

type publishServiceHandler struct {
	PublishServiceHandler
}

func (h *publishServiceHandler) CreatePublish(ctx context.Context, in *Publish, out *Response) error {
	return h.PublishServiceHandler.CreatePublish(ctx, in, out)
}

func (h *publishServiceHandler) ListOfPublishes(ctx context.Context, in *Request, out *Response) error {
	return h.PublishServiceHandler.ListOfPublishes(ctx, in, out)
}

func (h *publishServiceHandler) GetPublish(ctx context.Context, in *Publish, out *Response) error {
	return h.PublishServiceHandler.GetPublish(ctx, in, out)
}

func (h *publishServiceHandler) UpdatePublish(ctx context.Context, in *Publish, out *Response) error {
	return h.PublishServiceHandler.UpdatePublish(ctx, in, out)
}

func (h *publishServiceHandler) DeletePublish(ctx context.Context, in *Publish, out *Response) error {
	return h.PublishServiceHandler.DeletePublish(ctx, in, out)
}
