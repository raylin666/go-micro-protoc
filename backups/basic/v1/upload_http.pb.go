// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.3
// - protoc             v3.15.8
// source: basic/v1/upload.proto

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

const OperationUploadStreamUploadFile = "/basic.v1.Upload/StreamUploadFile"
const OperationUploadUrlUploadFile = "/basic.v1.Upload/UrlUploadFile"

type UploadHTTPServer interface {
	StreamUploadFile(context.Context, *StreamUploadFileRequest) (*StreamUploadFileResponse, error)
	UrlUploadFile(context.Context, *UrlUploadFileRequest) (*UrlUploadFileResponse, error)
}

func RegisterUploadHTTPServer(s *http.Server, srv UploadHTTPServer) {
	r := s.Route("/")
	r.PUT("/upload/file/stream", _Upload_StreamUploadFile0_HTTP_Handler(srv))
	r.POST("/upload/file/url", _Upload_UrlUploadFile0_HTTP_Handler(srv))
}

func _Upload_StreamUploadFile0_HTTP_Handler(srv UploadHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in StreamUploadFileRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUploadStreamUploadFile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.StreamUploadFile(ctx, req.(*StreamUploadFileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*StreamUploadFileResponse)
		return ctx.Result(200, reply)
	}
}

func _Upload_UrlUploadFile0_HTTP_Handler(srv UploadHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UrlUploadFileRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUploadUrlUploadFile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UrlUploadFile(ctx, req.(*UrlUploadFileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UrlUploadFileResponse)
		return ctx.Result(200, reply)
	}
}

type UploadHTTPClient interface {
	StreamUploadFile(ctx context.Context, req *StreamUploadFileRequest, opts ...http.CallOption) (rsp *StreamUploadFileResponse, err error)
	UrlUploadFile(ctx context.Context, req *UrlUploadFileRequest, opts ...http.CallOption) (rsp *UrlUploadFileResponse, err error)
}

type UploadHTTPClientImpl struct {
	cc *http.Client
}

func NewUploadHTTPClient(client *http.Client) UploadHTTPClient {
	return &UploadHTTPClientImpl{client}
}

func (c *UploadHTTPClientImpl) StreamUploadFile(ctx context.Context, in *StreamUploadFileRequest, opts ...http.CallOption) (*StreamUploadFileResponse, error) {
	var out StreamUploadFileResponse
	pattern := "/upload/file/stream"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUploadStreamUploadFile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UploadHTTPClientImpl) UrlUploadFile(ctx context.Context, in *UrlUploadFileRequest, opts ...http.CallOption) (*UrlUploadFileResponse, error) {
	var out UrlUploadFileResponse
	pattern := "/upload/file/url"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUploadUrlUploadFile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
