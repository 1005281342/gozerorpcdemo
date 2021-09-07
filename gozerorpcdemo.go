package main

import (
	"flag"
	"fmt"

	"github.com/1005281342/gozerorpcdemo/gozerorpcdemo"
	"github.com/1005281342/gozerorpcdemo/internal/config"
	"github.com/1005281342/gozerorpcdemo/internal/server"
	"github.com/1005281342/gozerorpcdemo/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/gozerorpcdemo.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewGozerorpcdemoServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		gozerorpcdemo.RegisterGozerorpcdemoServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
