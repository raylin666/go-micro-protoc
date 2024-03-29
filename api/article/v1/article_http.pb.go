// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.15.8
// source: article/v1/article.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationArticleInfo = "/article.v1.Article/Info"
const OperationArticleList = "/article.v1.Article/List"

type ArticleHTTPServer interface {
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
}

func RegisterArticleHTTPServer(s *http.Server, srv ArticleHTTPServer) {
	r := s.Route("/")
	r.GET("/list", _Article_List0_HTTP_Handler(srv))
	r.GET("/info/{id}", _Article_Info0_HTTP_Handler(srv))
}

func _Article_List0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*ListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListResponse)
		return ctx.Result(200, reply)
	}
}

func _Article_Info0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in InfoRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Info(ctx, req.(*InfoRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*InfoResponse)
		return ctx.Result(200, reply)
	}
}

type ArticleHTTPClient interface {
	Info(ctx context.Context, req *InfoRequest, opts ...http.CallOption) (rsp *InfoResponse, err error)
	List(ctx context.Context, req *ListRequest, opts ...http.CallOption) (rsp *ListResponse, err error)
}

type ArticleHTTPClientImpl struct {
	cc *http.Client
}

func NewArticleHTTPClient(client *http.Client) ArticleHTTPClient {
	return &ArticleHTTPClientImpl{client}
}

func (c *ArticleHTTPClientImpl) Info(ctx context.Context, in *InfoRequest, opts ...http.CallOption) (*InfoResponse, error) {
	var out InfoResponse
	pattern := "/info/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationArticleInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ArticleHTTPClientImpl) List(ctx context.Context, in *ListRequest, opts ...http.CallOption) (*ListResponse, error) {
	var out ListResponse
	pattern := "/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationArticleList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
