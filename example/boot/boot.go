package boot

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/glog"
)

func init() {
	glog.Info("init...")
	g.Config().SetPath("example")
	g.Config().SetFileName("config.toml")
}
