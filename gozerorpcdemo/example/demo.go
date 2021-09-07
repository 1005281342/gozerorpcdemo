package main

import (
	"context"
	"github.com/1005281342/gozerorpcdemo/gozerorpcdemo"
	"github.com/1005281342/gozerorpcdemo/gozerorpcdemoclient"
	"github.com/tal-tech/go-zero/zrpc"
	"log"
)

func main() {
	var zcli, _ = zrpc.NewClientWithTarget("127.0.0.1:8080")
	var cli = gozerorpcdemoclient.NewGozerorpcdemo(zcli)
	var rsp, err = cli.Ping(context.Background(), &gozerorpcdemo.Request{Ping: "Ping"})
	if err != nil {
		log.Fatalf("ping failed: %s", err.Error())
	}
	log.Printf("rsp %+v \n", rsp)
}
