package utils

import "github.com/gogf/gf/os/gtime"

func GetNow() string {
	return gtime.Datetime()
}
