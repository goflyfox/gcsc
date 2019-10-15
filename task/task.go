package task

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/glog"
)

func init() {
	glog.Info("task init...")

	cronCheckVersion := g.Config().GetString("config.cron-check-version", "0 * * * * *")
	gcron.Add(cronCheckVersion, checkVersion)

	glog.Info("task finish")
}
