package client

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gcache"
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

const CacheKey = "GCSC:CONFIG"

func Value(key string) string {
	return ""
}

func Code(key string) string {
	return ""
}

func ValueByCode(code string, configKey string) string {
	return ""
}

func List(configKey string) []ConfigBean {
	return nil
}

func ValueByProject(projectName string, key string) string {
	return ""
}

func CodeByProject(projectName string, key string) string {
	return ""
}

func ValueByCodeProject(projectName string, code string, configKey string) string {
	return ""
}

func ListByProject(projectName string, configKey string) []ConfigBean {
	return nil
}

// 获取缓存列表 缓存：projectName，ConfigListBean
func GetCache(projectName string) ConfigListBean {
	if projectName == "" {
		projectName = g.Config().GetString("projectName")
	}

	cacheData := gcache.Get(CacheKey + projectName)
	if cacheData == nil {
		return ConfigListBean{}
	}

	return cacheData.(ConfigListBean)
}

// 设置缓存列表 缓存：projectName，ConfigListBean
func SetCache(projectName string, listBean ConfigListBean) {
	gcache.Set(CacheKey+projectName, listBean, 0)
}
