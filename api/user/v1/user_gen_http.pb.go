// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.1
// - protoc             v3.19.6
// source: user/v1/user_gen.proto

package pbuser

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

const OperationUserCreate = "/user.v1.User/Create"
const OperationUserQuery = "/user.v1.User/Query"

type UserHTTPServer interface {
	Create(context.Context, *structpb.Struct) (*UserCreateResponse, error)
	Query(context.Context, *api.QueryRequest) (*UserQueryResponse, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/user", _User_Create0_HTTP_Handler(srv))
	r.GET("/user", _User_Query0_HTTP_Handler(srv))
}

func _User_Create0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in structpb.Struct
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserCreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*structpb.Struct))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserCreateResponse)
		return ctx.Result(200, reply)
	}
}

func _User_Query0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in api.QueryRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserQuery)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Query(ctx, req.(*api.QueryRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UserQueryResponse)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	Create(ctx context.Context, req *structpb.Struct, opts ...http.CallOption) (rsp *UserCreateResponse, err error)
	Query(ctx context.Context, req *api.QueryRequest, opts ...http.CallOption) (rsp *UserQueryResponse, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) Create(ctx context.Context, in *structpb.Struct, opts ...http.CallOption) (*UserCreateResponse, error) {
	var out UserCreateResponse
	pattern := "/user"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserCreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Query(ctx context.Context, in *api.QueryRequest, opts ...http.CallOption) (*UserQueryResponse, error) {
	var out UserQueryResponse
	pattern := "/user"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserQuery))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}