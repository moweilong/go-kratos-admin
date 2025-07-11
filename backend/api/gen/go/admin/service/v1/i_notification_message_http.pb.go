// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.4
// - protoc             (unknown)
// source: admin/service/v1/i_notification_message.proto

package servicev1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
	v1 "github.com/tx7do/kratos-bootstrap/api/gen/go/pagination/v1"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	v11 "kratos-admin/api/gen/go/internal_message/service/v1"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationNotificationMessageServiceCreate = "/admin.service.v1.NotificationMessageService/Create"
const OperationNotificationMessageServiceDelete = "/admin.service.v1.NotificationMessageService/Delete"
const OperationNotificationMessageServiceGet = "/admin.service.v1.NotificationMessageService/Get"
const OperationNotificationMessageServiceList = "/admin.service.v1.NotificationMessageService/List"
const OperationNotificationMessageServiceUpdate = "/admin.service.v1.NotificationMessageService/Update"

type NotificationMessageServiceHTTPServer interface {
	// Create 创建通知消息
	Create(context.Context, *v11.CreateNotificationMessageRequest) (*emptypb.Empty, error)
	// Delete 删除通知消息
	Delete(context.Context, *v11.DeleteNotificationMessageRequest) (*emptypb.Empty, error)
	// Get 查询通知消息详情
	Get(context.Context, *v11.GetNotificationMessageRequest) (*v11.NotificationMessage, error)
	// List 查询通知消息列表
	List(context.Context, *v1.PagingRequest) (*v11.ListNotificationMessageResponse, error)
	// Update 更新通知消息
	Update(context.Context, *v11.UpdateNotificationMessageRequest) (*emptypb.Empty, error)
}

func RegisterNotificationMessageServiceHTTPServer(s *http.Server, srv NotificationMessageServiceHTTPServer) {
	r := s.Route("/")
	r.GET("/admin/v1/notifications", _NotificationMessageService_List8_HTTP_Handler(srv))
	r.GET("/admin/v1/notifications/{id}", _NotificationMessageService_Get8_HTTP_Handler(srv))
	r.POST("/admin/v1/notifications", _NotificationMessageService_Create6_HTTP_Handler(srv))
	r.PUT("/admin/v1/notifications/{data.id}", _NotificationMessageService_Update6_HTTP_Handler(srv))
	r.DELETE("/admin/v1/notifications/{id}", _NotificationMessageService_Delete6_HTTP_Handler(srv))
}

func _NotificationMessageService_List8_HTTP_Handler(srv NotificationMessageServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v1.PagingRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNotificationMessageServiceList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.List(ctx, req.(*v1.PagingRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.ListNotificationMessageResponse)
		return ctx.Result(200, reply)
	}
}

func _NotificationMessageService_Get8_HTTP_Handler(srv NotificationMessageServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.GetNotificationMessageRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNotificationMessageServiceGet)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*v11.GetNotificationMessageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*v11.NotificationMessage)
		return ctx.Result(200, reply)
	}
}

func _NotificationMessageService_Create6_HTTP_Handler(srv NotificationMessageServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.CreateNotificationMessageRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNotificationMessageServiceCreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*v11.CreateNotificationMessageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _NotificationMessageService_Update6_HTTP_Handler(srv NotificationMessageServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.UpdateNotificationMessageRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNotificationMessageServiceUpdate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*v11.UpdateNotificationMessageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

func _NotificationMessageService_Delete6_HTTP_Handler(srv NotificationMessageServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in v11.DeleteNotificationMessageRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationNotificationMessageServiceDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*v11.DeleteNotificationMessageRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*emptypb.Empty)
		return ctx.Result(200, reply)
	}
}

type NotificationMessageServiceHTTPClient interface {
	Create(ctx context.Context, req *v11.CreateNotificationMessageRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Delete(ctx context.Context, req *v11.DeleteNotificationMessageRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
	Get(ctx context.Context, req *v11.GetNotificationMessageRequest, opts ...http.CallOption) (rsp *v11.NotificationMessage, err error)
	List(ctx context.Context, req *v1.PagingRequest, opts ...http.CallOption) (rsp *v11.ListNotificationMessageResponse, err error)
	Update(ctx context.Context, req *v11.UpdateNotificationMessageRequest, opts ...http.CallOption) (rsp *emptypb.Empty, err error)
}

type NotificationMessageServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewNotificationMessageServiceHTTPClient(client *http.Client) NotificationMessageServiceHTTPClient {
	return &NotificationMessageServiceHTTPClientImpl{client}
}

func (c *NotificationMessageServiceHTTPClientImpl) Create(ctx context.Context, in *v11.CreateNotificationMessageRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/notifications"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationNotificationMessageServiceCreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *NotificationMessageServiceHTTPClientImpl) Delete(ctx context.Context, in *v11.DeleteNotificationMessageRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/notifications/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNotificationMessageServiceDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *NotificationMessageServiceHTTPClientImpl) Get(ctx context.Context, in *v11.GetNotificationMessageRequest, opts ...http.CallOption) (*v11.NotificationMessage, error) {
	var out v11.NotificationMessage
	pattern := "/admin/v1/notifications/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNotificationMessageServiceGet))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *NotificationMessageServiceHTTPClientImpl) List(ctx context.Context, in *v1.PagingRequest, opts ...http.CallOption) (*v11.ListNotificationMessageResponse, error) {
	var out v11.ListNotificationMessageResponse
	pattern := "/admin/v1/notifications"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationNotificationMessageServiceList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *NotificationMessageServiceHTTPClientImpl) Update(ctx context.Context, in *v11.UpdateNotificationMessageRequest, opts ...http.CallOption) (*emptypb.Empty, error) {
	var out emptypb.Empty
	pattern := "/admin/v1/notifications/{data.id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationNotificationMessageServiceUpdate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
