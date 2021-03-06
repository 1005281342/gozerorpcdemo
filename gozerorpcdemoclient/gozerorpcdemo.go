// Code generated by goctl. DO NOT EDIT!
// Source: gozerorpcdemo.proto

package gozerorpcdemoclient

import (
	"context"

	"github.com/1005281342/gozerorpcdemo/gozerorpcdemo"

	"github.com/tal-tech/go-zero/zrpc"
)

type (
	Response = gozerorpcdemo.Response
	Request  = gozerorpcdemo.Request

	Gozerorpcdemo interface {
		Ping(ctx context.Context, in *Request) (*Response, error)
	}

	defaultGozerorpcdemo struct {
		cli zrpc.Client
	}
)

func NewGozerorpcdemo(cli zrpc.Client) Gozerorpcdemo {
	return &defaultGozerorpcdemo{
		cli: cli,
	}
}

func (m *defaultGozerorpcdemo) Ping(ctx context.Context, in *Request) (*Response, error) {
	client := gozerorpcdemo.NewGozerorpcdemoClient(m.cli.Conn())
	return client.Ping(ctx, in)
}
