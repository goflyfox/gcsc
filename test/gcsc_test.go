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

// 临时目录
func TestTmpPath(t *testing.T) {
	glog.Info("[Test]####TestTmpPath start...")
	fmt.Println(gfile.TempDir() + gfile.Separator + "configClient")
}

// 测试初始化
func TestInit(t *testing.T) {
	glog.Info("[Test]####TestInit start...")
	time.Sleep(time.Second * 2)
}

// 测试版本不一致重新获取
func TestTask(t *testing.T) {
	glog.Info("[Test]####TestTask start...")
	cronCheckVersion := g.Config().GetString(constant.ParamCronCheckVersion, "* * * * * *")
	gcron.Add(cronCheckVersion, task.CheckVersion)

	time.Sleep(3 * time.Second)
}

// 重新初始化
func TestInitFunc(t *testing.T) {
	glog.Info("[Test]####TestInitFunc start...")
	filePath := task.GetDataPath() + "test.txt"
	if gfile.Exists(filePath) {
		gfile.Remove(filePath)
	}
	task.InitConfigData()
}

// 测试Value和Code获取接口
func TestGet(t *testing.T) {
	glog.Info("[Test]####TestGet start...")
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

// 测试列表获取接口
func TestList(t *testing.T) {
	glog.Info("[Test]####TestList start...")
	testList := client.List("system")
	glog.Info("[Test]List:", testList)
	if len(testList) <= 0 {
		t.Error("[Test]List lt zero")
	}

	flag := true
	for _, bean := range testList {
		if bean.Key == "test" {
			glog.Info("[Test]List test:", bean.Value)
			flag = false
			break
		}
	}

	if flag {
		t.Error("[Test]List test key is not exist")
	}

	testList = client.ListByProject("test", "system")
	glog.Info("[Test]ListByProject:", testList)
	if len(testList) <= 0 {
		t.Error("[Test]ListByProject lt zero")
	}

	for _, bean := range testList {
		if bean.Key == "test" {
			glog.Info("[Test]List test:", bean.Value)
			flag = false
			break
		}
	}

	if flag {
		t.Error("[Test]List test key is not exist")
	}
}
