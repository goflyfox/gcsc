package client

import (
	"github.com/goflyfox/gcsc/constant"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
	"github.com/gogf/gf/text/gstr"
)

type ConfigBean struct {
	// columns START
	Name      string `json:"name" gconv:"name,omitempty"`            // 名称
	Key       string `json:"key" gconv:"key,omitempty"`              // 键
	Value     string `json:"value" gconv:"value,omitempty"`          // 值
	Code      string `json:"code" gconv:"code,omitempty"`            // 编码
	ParentKey string `json:"parentKey" gconv:"parent_key,omitempty"` //
	Sort      int    `json:"sort" gconv:"sort,omitempty"`            // 排序号
	ProjectId int    `json:"projectId" gconv:"project_id,omitempty"` // 项目ID
}

type ConfigListBean struct {
	// columns START
	Name     string       `json:"name" gconv:"name,omitempty"`
	Version  string       `json:"version" gconv:"version,omitempty"`
	ListData []ConfigBean `json:"value" gconv:"value,omitempty"`
}

func Value(key string) string {
	return ValueByProject("", key)
}

func Code(key string) string {
	return CodeByProject("", key)
}

func ValueByCode(code string, parentKey string) string {
	return ValueByCodeProject("", code, parentKey)
}

func List(parentKey string) []ConfigBean {
	return ListByProject("", parentKey)
}

func ValueByProject(projectName string, key string) string {
	list := GetCache(projectName).ListData
	for _, bean := range list {
		if bean.Key == key {
			return bean.Value
		}
	}

	return ""
}

func CodeByProject(projectName string, key string) string {
	list := GetCache(projectName).ListData
	for _, bean := range list {
		if bean.Key == key {
			return bean.Code
		}
	}

	return ""
}

func ValueByCodeProject(projectName string, code string, parentKey string) string {
	list := GetCache(projectName).ListData

	for _, bean := range list {
		if bean.ParentKey == parentKey && bean.Code == code {
			return bean.Value
		}
	}

	return ""
}

func ListByProject(projectName string, parentKey string) []ConfigBean {
	var configBeans []ConfigBean
	list := GetCache(projectName).ListData
	for _, bean := range list {
		if bean.ParentKey == parentKey {
			configBeans = append(configBeans, bean)
		}
	}

	return configBeans
}

// 获取缓存列表 缓存：projectName，ConfigListBean
func GetCache(projectName ...string) ConfigListBean {
	if projectName[0] == "" {
		projectNames := g.Config().GetString(constant.ParamProjectName)
		projectName[0] = gstr.Split(projectNames, ",")[0]
	}

	cacheData := gcache.Get(constant.CacheKey + projectName[0])
	if cacheData == nil {
		return ConfigListBean{}
	}

	return cacheData.(ConfigListBean)
}

// 设置缓存列表 缓存：projectName，ConfigListBean
func SetCache(listBean ConfigListBean) {
	gcache.Set(constant.CacheKey+listBean.Name, listBean, 0)
}
