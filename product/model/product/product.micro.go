// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: product.proto

/*
Package publishen_product_service is a generated protocol buffer package.

It is generated from these files:
	product.proto

It has these top-level messages:
	Product
	Location
	Attribute
	Response
*/
package publishen_product_service

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

// Client API for ProductService service

type ProductService interface {
	CreateProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error)
}

type productService struct {
	c    client.Client
	name string
}

func NewProductService(name string, c client.Client) ProductService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "publishen.product.service"
	}
	return &productService{
		c:    c,
		name: name,
	}
}

func (c *productService) CreateProduct(ctx context.Context, in *Product, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "ProductService.CreateProduct", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ProductService service

type ProductServiceHandler interface {
	CreateProduct(context.Context, *Product, *Response) error
}

func RegisterProductServiceHandler(s server.Server, hdlr ProductServiceHandler, opts ...server.HandlerOption) error {
	type productService interface {
		CreateProduct(ctx context.Context, in *Product, out *Response) error
	}
	type ProductService struct {
		productService
	}
	h := &productServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&ProductService{h}, opts...))
}

type productServiceHandler struct {
	ProductServiceHandler
}

func (h *productServiceHandler) CreateProduct(ctx context.Context, in *Product, out *Response) error {
	return h.ProductServiceHandler.CreateProduct(ctx, in, out)
}
