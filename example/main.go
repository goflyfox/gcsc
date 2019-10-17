package main

import (
	"github.com/goflyfox/gcsc/client"
	_ "github.com/goflyfox/gcsc/example/boot"
	_ "github.com/goflyfox/gcsc/task"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

func main() {
	s := g.Server()
	// 调试路由
	s.BindHandler("/test", func(r *ghttp.Request) {
		r.Response.Write(client.Value("test"))
	})

	glog.Info("[TEST]get test value", client.Value("test"))

	s.SetPort(g.Config().GetInt("http-port"))
	g.Server().Run()
}
