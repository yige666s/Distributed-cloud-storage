// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: user.proto

package user

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

// Client API for UserService service

type UserService interface {
	SignUp(ctx context.Context, in *ReqSignup, opts ...client.CallOption) (*RespSignup, error)
}

type userService struct {
	c    client.Client
	name string
}

func NewUserService(name string, c client.Client) UserService {
	return &userService{
		c:    c,
		name: name,
	}
}

func (c *userService) SignUp(ctx context.Context, in *ReqSignup, opts ...client.CallOption) (*RespSignup, error) {
	req := c.c.NewRequest(c.name, "UserService.SignUp", in)
	out := new(RespSignup)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UserService service

type UserServiceHandler interface {
	SignUp(context.Context, *ReqSignup, *RespSignup) error
}

func RegisterUserServiceHandler(s server.Server, hdlr UserServiceHandler, opts ...server.HandlerOption) error {
	type userService interface {
		SignUp(ctx context.Context, in *ReqSignup, out *RespSignup) error
	}
	type UserService struct {
		userService
	}
	h := &userServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&UserService{h}, opts...))
}

type userServiceHandler struct {
	UserServiceHandler
}

func (h *userServiceHandler) SignUp(ctx context.Context, in *ReqSignup, out *RespSignup) error {
	return h.UserServiceHandler.SignUp(ctx, in, out)
}
