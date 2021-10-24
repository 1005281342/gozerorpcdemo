package logic

import (
	"context"
	"time"

	"github.com/1005281342/gozerorpcdemo/gozerorpcdemo"
	"github.com/1005281342/gozerorpcdemo/internal/svc"
	"github.com/1005281342/log"
	"github.com/tal-tech/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type ELKLogger struct {
	log.Logger
}

func (e *ELKLogger) WithDuration(d time.Duration) logx.Logger {
	e.Logger.WithDuration(d)
	return e
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		//Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *gozerorpcdemo.Request) (*gozerorpcdemo.Response, error) {
	log.SetTraceID(l.ctx, "jesontest go-zero demo")
	var (
		logger         logx.Logger
		loggerELK, err = log.WithContext(l.ctx,
			log.WithAddress("127.0.0.1:5000"),
			log.WithAppName("gozerorpcdemo"),
			log.WithFuncName("Ping"),
		)
	)
	if err != nil {
		logger = logx.WithContext(l.ctx)
		logger.Errorf("err: %+v", err)
	} else {
		logger = &ELKLogger{Logger: loggerELK}
	}

	var out = &gozerorpcdemo.Response{Pong: "OK. " + in.Ping}
	logger.Infof("out: %+v", out)

	return out, nil
}
