// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.6.1
// source: car/v1/car_gen.proto

package pbcar

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	api "github.com/pigfall/demo-kratos-curdboy/api"
	structpb "google.golang.org/protobuf/types/known/structpb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationCarCreate = "/car.v1.Car/Create"
const OperationCarQuery = "/car.v1.Car/Query"

type CarHTTPServer interface {
	Create(context.Context, *structpb.Struct) (*CarCreateResponse, error)
	Query(context.Context, *api.QueryRequest) (*CarQueryResponse, error)
}

func RegisterCarHTTPServer(s *http.Server, srv CarHTTPServer) {
	r := s.Route("/")
	r.POST("/car", _Car_Create1_HTTP_Handler(srv))
	r.GET("/car", _Car_Query1_HTTP_Handler(srv))
}

func _Car_Create1_HTTP_Handler(srv CarHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in structpb.Struct
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCarCreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*structpb.Struct))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CarCreateResponse)
		return ctx.Result(200, reply)
	}
}

func _Car_Query1_HTTP_Handler(srv CarHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in api.QueryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationCarQuery)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Query(ctx, req.(*api.QueryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CarQueryResponse)
		return ctx.Result(200, reply)
	}
}

type CarHTTPClient interface {
	Create(ctx context.Context, req *structpb.Struct, opts ...http.CallOption) (rsp *CarCreateResponse, err error)
	Query(ctx context.Context, req *api.QueryRequest, opts ...http.CallOption) (rsp *CarQueryResponse, err error)
}

type CarHTTPClientImpl struct {
	cc *http.Client
}

func NewCarHTTPClient(client *http.Client) CarHTTPClient {
	return &CarHTTPClientImpl{client}
}

func (c *CarHTTPClientImpl) Create(ctx context.Context, in *structpb.Struct, opts ...http.CallOption) (*CarCreateResponse, error) {
	var out CarCreateResponse
	pattern := "/car"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationCarCreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *CarHTTPClientImpl) Query(ctx context.Context, in *api.QueryRequest, opts ...http.CallOption) (*CarQueryResponse, error) {
	var out CarQueryResponse
	pattern := "/car"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationCarQuery))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}