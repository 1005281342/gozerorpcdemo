package logic

import (
	"context"

	"github.com/1005281342/gozerorpcdemo/gozerorpcdemo"
	"github.com/1005281342/gozerorpcdemo/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *gozerorpcdemo.Request) (*gozerorpcdemo.Response, error) {
	return &gozerorpcdemo.Response{Pong: "OK. " + in.Ping}, nil
}
