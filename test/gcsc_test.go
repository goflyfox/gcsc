package test

import (
	"fmt"
	"github.com/goflyfox/gcsc/client"
	"github.com/goflyfox/gcsc/constant"
	"github.com/goflyfox/gcsc/task"
	_ "github.com/goflyfox/gcsc/task"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcron"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"testing"
	"time"
)

func TestTmpPath(t *testing.T) {
	fmt.Println(gfile.TempDir() + gfile.Separator + "configClient")
}

func TestTask(t *testing.T) {
	cronCheckVersion := g.Config().GetString(constant.ParamCronCheckVersion, "* * * * * *")
	gcron.Add(cronCheckVersion, task.CheckVersion)

	time.Sleep(10 * time.Second)
}

func TestInit(t *testing.T) {
	time.Sleep(time.Second * 10)
}

func TestGet(t *testing.T) {
	//task.InitConfigData()
	testValue := client.Value("test")
	glog.Info("[Test]Value:" + testValue)
	if testValue != "value_test" {
		t.Error("[Test]Value get ne value_test")
	}

	testCode := client.Code("test")
	glog.Info("[Test]Code:" + testCode)
	if testCode != "code_test" {
		t.Error("[Test]Code get ne code_test")
	}

	testValue = client.ValueByProject("test", "test")
	glog.Info("[Test]ValueByProject:" + testValue)
	if testValue != "value_test" {
		t.Error("[Test]ValueByProject get ne value_test")
	}

	testCode = client.CodeByProject("test", "test")
	glog.Info("[Test]CodeByProject:" + testCode)
	if testCode != "code_test" {
		t.Error("[Test]CodeByProject get ne code_test")
	}

	testValue = client.ValueByCode("code_test", "system")
	glog.Info("[Test]ValueByCode:" + testValue)
	if testValue != "value_test" {
		t.Error("[Test]ValueByCode get ne value_test")
	}

	testValue = client.ValueByCodeProject("test", "code_test", "system")
	glog.Info("[Test]ValueByCodeProject:" + testValue)
	if testValue != "value_test" {
		t.Error("[Test]ValueByCodeProject get ne value_test")
	}
}

func TestList(t *testing.T) {
	//task.InitConfigData()
	testValue := client.Value("test")
	glog.Info("[Test]testValue:" + testValue)
	if testValue != "value_test" {
		t.Error("[Test]testValue get ne value_test")
	}

	testCode := client.Code("test")
	glog.Info("[Test]testCode:" + testCode)
	if testCode != "code_test" {
		t.Error("[Test]testCode get ne code_test")
	}

	testValue = client.ValueByProject("test", "test")
	glog.Info("[Test]testValue:" + testValue)
	if testValue != "value_test" {
		t.Error("[Test]testValue get ne value_test")
	}

	testCode = client.CodeByProject("test", "test")
	glog.Info("[Test]testCode:" + testCode)
	if testCode != "code_test" {
		t.Error("[Test]testCode get ne code_test")
	}
}
