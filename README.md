# README

## 创建项目

`goctl rpc new gozerorpcdemo`

## go mod init

`go mod init "github.com/1005281342/gozerorpcdemo"`

## 替换桩代码包名

将桩代码包名`__`替换为`gozerorpcdemo`

## RPC调用

```go
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

//2021/09/07 22:35:19 {"@timestamp":"2021-09-07T22:35:19.283+08","level":"stat","content":"p2c - conn: 127.0.0.1:8080, load: 1027, reqs: 1"}
//2021/09/07 22:35:19 rsp pong:"OK. Ping"
```
