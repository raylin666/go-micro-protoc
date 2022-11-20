// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.15.8
// source: basic/v1/short_link.proto

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

const OperationShortLinkGenerateShortUrl = "/basic.v1.ShortLink/GenerateShortUrl"
const OperationShortLinkTransformLongUrl = "/basic.v1.ShortLink/TransformLongUrl"

type ShortLinkHTTPServer interface {
	GenerateShortUrl(context.Context, *GenerateShortUrlRequest) (*GenerateShortUrlResponse, error)
	TransformLongUrl(context.Context, *TransformLongUrlRequest) (*TransformLongUrlResponse, error)
}

func RegisterShortLinkHTTPServer(s *http.Server, srv ShortLinkHTTPServer) {
	r := s.Route("/")
	r.POST("/short_url/generate", _ShortLink_GenerateShortUrl0_HTTP_Handler(srv))
	r.POST("/short_url/transform_long_url", _ShortLink_TransformLongUrl0_HTTP_Handler(srv))
}

func _ShortLink_GenerateShortUrl0_HTTP_Handler(srv ShortLinkHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GenerateShortUrlRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShortLinkGenerateShortUrl)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GenerateShortUrl(ctx, req.(*GenerateShortUrlRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GenerateShortUrlResponse)
		return ctx.Result(200, reply)
	}
}

func _ShortLink_TransformLongUrl0_HTTP_Handler(srv ShortLinkHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in TransformLongUrlRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationShortLinkTransformLongUrl)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.TransformLongUrl(ctx, req.(*TransformLongUrlRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*TransformLongUrlResponse)
		return ctx.Result(200, reply)
	}
}

type ShortLinkHTTPClient interface {
	GenerateShortUrl(ctx context.Context, req *GenerateShortUrlRequest, opts ...http.CallOption) (rsp *GenerateShortUrlResponse, err error)
	TransformLongUrl(ctx context.Context, req *TransformLongUrlRequest, opts ...http.CallOption) (rsp *TransformLongUrlResponse, err error)
}

type ShortLinkHTTPClientImpl struct {
	cc *http.Client
}

func NewShortLinkHTTPClient(client *http.Client) ShortLinkHTTPClient {
	return &ShortLinkHTTPClientImpl{client}
}

func (c *ShortLinkHTTPClientImpl) GenerateShortUrl(ctx context.Context, in *GenerateShortUrlRequest, opts ...http.CallOption) (*GenerateShortUrlResponse, error) {
	var out GenerateShortUrlResponse
	pattern := "/short_url/generate"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShortLinkGenerateShortUrl))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ShortLinkHTTPClientImpl) TransformLongUrl(ctx context.Context, in *TransformLongUrlRequest, opts ...http.CallOption) (*TransformLongUrlResponse, error) {
	var out TransformLongUrlResponse
	pattern := "/short_url/transform_long_url"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationShortLinkTransformLongUrl))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
