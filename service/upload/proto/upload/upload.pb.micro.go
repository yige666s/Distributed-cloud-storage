// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: upload.proto

package upload

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "micro.dev/v4/service/client"
	server "micro.dev/v4/service/server"
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

// Client API for UploadEntry service

type UploadEntryService interface {
	// 获取上传入口地址
	UploadEntry(ctx context.Context, in *RespEntry, opts ...client.CallOption) (*RespEntry, error)
}

type uploadEntryService struct {
	c    client.Client
	name string
}

func NewUploadEntryService(name string, c client.Client) UploadEntryService {
	return &uploadEntryService{
		c:    c,
		name: name,
	}
}

func (c *uploadEntryService) UploadEntry(ctx context.Context, in *RespEntry, opts ...client.CallOption) (*RespEntry, error) {
	req := c.c.NewRequest(c.name, "UploadEntry.UploadEntry", in)
	out := new(RespEntry)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UploadEntry service

type UploadEntryHandler interface {
	// 获取上传入口地址
	UploadEntry(context.Context, *RespEntry, *RespEntry) error
}

func RegisterUploadEntryHandler(s server.Server, hdlr UploadEntryHandler, opts ...server.HandlerOption) error {
	type uploadEntry interface {
		UploadEntry(ctx context.Context, in *RespEntry, out *RespEntry) error
	}
	type UploadEntry struct {
		uploadEntry
	}
	h := &uploadEntryHandler{hdlr}
	return s.Handle(s.NewHandler(&UploadEntry{h}, opts...))
}

type uploadEntryHandler struct {
	UploadEntryHandler
}

func (h *uploadEntryHandler) UploadEntry(ctx context.Context, in *RespEntry, out *RespEntry) error {
	return h.UploadEntryHandler.UploadEntry(ctx, in, out)
}
