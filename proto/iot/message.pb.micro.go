// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: iot/message.proto

package message

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	proto1 "github.com/micro/go-micro/api/proto"
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

// Client API for IotSvc service

type IotSvcService interface {
	GetGateway(ctx context.Context, in *ThingRequest, opts ...client.CallOption) (*ThingResponse, error)
	GetThing(ctx context.Context, in *ThingRequest, opts ...client.CallOption) (*ThingResponse, error)
	ExcuteTemplate(ctx context.Context, in *DataRequest, opts ...client.CallOption) (*DataResponse, error)
	Log(ctx context.Context, in *DataRequest, opts ...client.CallOption) (*DataResponse, error)
	GetDeviceMapFail(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error)
	CountThing(ctx context.Context, in *wrappers.StringValue, opts ...client.CallOption) (*wrappers.Int64Value, error)
}

type iotSvcService struct {
	c    client.Client
	name string
}

func NewIotSvcService(name string, c client.Client) IotSvcService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "iotsvc"
	}
	return &iotSvcService{
		c:    c,
		name: name,
	}
}

func (c *iotSvcService) GetGateway(ctx context.Context, in *ThingRequest, opts ...client.CallOption) (*ThingResponse, error) {
	req := c.c.NewRequest(c.name, "IotSvc.GetGateway", in)
	out := new(ThingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotSvcService) GetThing(ctx context.Context, in *ThingRequest, opts ...client.CallOption) (*ThingResponse, error) {
	req := c.c.NewRequest(c.name, "IotSvc.GetThing", in)
	out := new(ThingResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotSvcService) ExcuteTemplate(ctx context.Context, in *DataRequest, opts ...client.CallOption) (*DataResponse, error) {
	req := c.c.NewRequest(c.name, "IotSvc.ExcuteTemplate", in)
	out := new(DataResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotSvcService) Log(ctx context.Context, in *DataRequest, opts ...client.CallOption) (*DataResponse, error) {
	req := c.c.NewRequest(c.name, "IotSvc.Log", in)
	out := new(DataResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotSvcService) GetDeviceMapFail(ctx context.Context, in *proto1.Request, opts ...client.CallOption) (*proto1.Response, error) {
	req := c.c.NewRequest(c.name, "IotSvc.GetDeviceMapFail", in)
	out := new(proto1.Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *iotSvcService) CountThing(ctx context.Context, in *wrappers.StringValue, opts ...client.CallOption) (*wrappers.Int64Value, error) {
	req := c.c.NewRequest(c.name, "IotSvc.CountThing", in)
	out := new(wrappers.Int64Value)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for IotSvc service

type IotSvcHandler interface {
	GetGateway(context.Context, *ThingRequest, *ThingResponse) error
	GetThing(context.Context, *ThingRequest, *ThingResponse) error
	ExcuteTemplate(context.Context, *DataRequest, *DataResponse) error
	Log(context.Context, *DataRequest, *DataResponse) error
	GetDeviceMapFail(context.Context, *proto1.Request, *proto1.Response) error
	CountThing(context.Context, *wrappers.StringValue, *wrappers.Int64Value) error
}

func RegisterIotSvcHandler(s server.Server, hdlr IotSvcHandler, opts ...server.HandlerOption) error {
	type iotSvc interface {
		GetGateway(ctx context.Context, in *ThingRequest, out *ThingResponse) error
		GetThing(ctx context.Context, in *ThingRequest, out *ThingResponse) error
		ExcuteTemplate(ctx context.Context, in *DataRequest, out *DataResponse) error
		Log(ctx context.Context, in *DataRequest, out *DataResponse) error
		GetDeviceMapFail(ctx context.Context, in *proto1.Request, out *proto1.Response) error
		CountThing(ctx context.Context, in *wrappers.StringValue, out *wrappers.Int64Value) error
	}
	type IotSvc struct {
		iotSvc
	}
	h := &iotSvcHandler{hdlr}
	return s.Handle(s.NewHandler(&IotSvc{h}, opts...))
}

type iotSvcHandler struct {
	IotSvcHandler
}

func (h *iotSvcHandler) GetGateway(ctx context.Context, in *ThingRequest, out *ThingResponse) error {
	return h.IotSvcHandler.GetGateway(ctx, in, out)
}

func (h *iotSvcHandler) GetThing(ctx context.Context, in *ThingRequest, out *ThingResponse) error {
	return h.IotSvcHandler.GetThing(ctx, in, out)
}

func (h *iotSvcHandler) ExcuteTemplate(ctx context.Context, in *DataRequest, out *DataResponse) error {
	return h.IotSvcHandler.ExcuteTemplate(ctx, in, out)
}

func (h *iotSvcHandler) Log(ctx context.Context, in *DataRequest, out *DataResponse) error {
	return h.IotSvcHandler.Log(ctx, in, out)
}

func (h *iotSvcHandler) GetDeviceMapFail(ctx context.Context, in *proto1.Request, out *proto1.Response) error {
	return h.IotSvcHandler.GetDeviceMapFail(ctx, in, out)
}

func (h *iotSvcHandler) CountThing(ctx context.Context, in *wrappers.StringValue, out *wrappers.Int64Value) error {
	return h.IotSvcHandler.CountThing(ctx, in, out)
}
