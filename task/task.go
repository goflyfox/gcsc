package task

import (
	"github.com/goflyfox/gcsc/constant"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/glog"
)

func init() {
	glog.Info("[Task]init...")

	cronCheckVersion := g.Config().GetString(constant.ParamCronCheckVersion, "0 * * * * *")
	gcron.Add(cronCheckVersion, CheckVersion)
	// 初始化
	InitConfigData()

	glog.Info("[Task]finish.")
}
