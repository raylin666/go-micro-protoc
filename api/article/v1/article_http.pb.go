// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.5.0
// - protoc             v3.15.8
// source: article/v1/article.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationArticleAdd = "/article.v1.Article/Add"
const OperationArticleCategoryList = "/article.v1.Article/CategoryList"
const OperationArticleDelete = "/article.v1.Article/Delete"
const OperationArticleForceDelete = "/article.v1.Article/ForceDelete"
const OperationArticleInfo = "/article.v1.Article/Info"
const OperationArticleList = "/article.v1.Article/List"
const OperationArticleUpdate = "/article.v1.Article/Update"
const OperationArticleUpdateField = "/article.v1.Article/UpdateField"

type ArticleHTTPServer interface {
	Add(context.Context, *AddRequest) (*AddResponse, error)
	CategoryList(context.Context, *emptypb.Empty) (*CategoryListResponse, error)
	Delete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	ForceDelete(context.Context, *DeleteRequest) (*emptypb.Empty, error)
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	List(context.Context, *ListRequest) (*ListResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	UpdateField(context.Context, *UpdateFieldRequest) (*emptypb.Empty, error)
}

func RegisterArticleHTTPServer(s *http.Server, srv ArticleHTTPServer) {
	r := s.Route("/")
	r.GET("/article/list", _Article_List0_HTTP_Handler(srv))
	r.GET("/article/info/{id}", _Article_Info0_HTTP_Handler(srv))
	r.PUT("/article/add", _Article_Add0_HTTP_Handler(srv))
	r.POST("/article/update/{id}", _Article_Update0_HTTP_Handler(srv))
	r.DELETE("/article/delete/{id}", _Article_Delete0_HTTP_Handler(srv))
	r.DELETE("/article/force_delete/{id}", _Article_ForceDelete0_HTTP_Handler(srv))
	r.POST("/article/update_field/{id}/{field}", _Article_UpdateField0_HTTP_Handler(srv))
	r.GET("/article/category/list", _Article_CategoryList0_HTTP_Handler(srv))
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

func _Article_Add0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AddRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleAdd)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Add(ctx, req.(*AddRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AddResponse)
		return ctx.Result(200, reply)
	}
}

func _Article_Update0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleUpdate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*UpdateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateResponse)
		return ctx.Result(200, reply)
	}
}

func _Article_Delete0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*DeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Article_ForceDelete0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleForceDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ForceDelete(ctx, req.(*DeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Article_UpdateField0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateFieldRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleUpdateField)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateField(ctx, req.(*UpdateFieldRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _Article_CategoryList0_HTTP_Handler(srv ArticleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in emptypb.Empty
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationArticleCategoryList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CategoryList(ctx, req.(*emptypb.Empty))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CategoryListResponse)
		return ctx.Result(200, reply)
	}
}

type ArticleHTTPClient interface {
	Add(ctx context.Context, req *AddRequest, opts ...http.CallOption) (rsp *AddResponse, err error)
	CategoryList(ctx context.Context, req *emptypb.Empty, opts ...http.CallOption) (rsp *CategoryListResponse, err error)
	Delete(ctx context.Context, req *DeleteRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	ForceDelete(ctx context.Context, req *DeleteRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Info(ctx context.Context, req *InfoRequest, opts ...http.CallOption) (rsp *InfoResponse, err error)
	List(ctx context.Context, req *ListRequest, opts ...http.CallOption) (rsp *ListResponse, err error)
	Update(ctx context.Context, req *UpdateRequest, opts ...http.CallOption) (rsp *UpdateResponse, err error)
	UpdateField(ctx context.Context, req *UpdateFieldRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type ArticleHTTPClientImpl struct {
	cc *http.Client
}

func NewArticleHTTPClient(client *http.Client) ArticleHTTPClient {
	return &ArticleHTTPClientImpl{client}
}

func (c *ArticleHTTPClientImpl) Add(ctx context.Context, in *AddRequest, opts ...http.CallOption) (*AddResponse, error) {
	var out AddResponse
	pattern := "/article/add"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationArticleAdd))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ArticleHTTPClientImpl) CategoryList(ctx context.Context, in *emptypb.Empty, opts ...http.CallOption) (*CategoryListResponse, error) {
	var out CategoryListResponse
	pattern := "/article/category/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationArticleCategoryList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ArticleHTTPClientImpl) Delete(ctx context.Context, in *DeleteRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/article/delete/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationArticleDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ArticleHTTPClientImpl) ForceDelete(ctx context.Context, in *DeleteRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/article/force_delete/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationArticleForceDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ArticleHTTPClientImpl) Info(ctx context.Context, in *InfoRequest, opts ...http.CallOption) (*InfoResponse, error) {
	var out InfoResponse
	pattern := "/article/info/{id}"
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
	pattern := "/article/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationArticleList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ArticleHTTPClientImpl) Update(ctx context.Context, in *UpdateRequest, opts ...http.CallOption) (*UpdateResponse, error) {
	var out UpdateResponse
	pattern := "/article/update/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationArticleUpdate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ArticleHTTPClientImpl) UpdateField(ctx context.Context, in *UpdateFieldRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/article/update_field/{id}/{field}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationArticleUpdateField))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
