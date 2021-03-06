package task

import (
	"encoding/json"
	"github.com/goflyfox/gcsc/client"
	"github.com/goflyfox/gcsc/constant"
	"github.com/goflyfox/gcsc/utils/resp"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gstr"
)

func CheckVersion() {
	// form := base.BaseForm{}
	// config.TbConfigPublic{}.GetCacheModel(&form)
	projectNameStr := g.Config().GetString(constant.ParamProjectName)
	projectSecretStr := g.Config().GetString(constant.ParamProjectSecret)
	if projectNameStr == "" {
		glog.Warning("[ConfigClient]updateData projects is null")
		return
	}

	projectNames := gstr.Split(projectNameStr, ",")
	projectSecrets := gstr.Split(projectSecretStr, ",")
	if len(projectNames) != len(projectSecrets) {
		glog.Error("[ConfigClient]projects config length err")
		return
	}

	serverUrl := g.Config().GetString(constant.ParamServerUrl, "http://127.0.0.1")
	serverVersionPath := g.Config().GetString(constant.ParamServerVersionPath, "/config/api/version")
	serverDataPath := g.Config().GetString(constant.ParamServerDataPath, "/config/api/data")
	for index, projectName := range projectNames {
		projectSecret := projectSecrets[index]
		nowTime := gtime.Now().Format("YmdHis")
		mac, err := gmd5.Encrypt(projectName + nowTime + projectSecret)
		if err != nil {
			glog.Error("[ConfigClient]updateData md5 err", err)
			continue
		}

		// 查询版本
		var serverVersion string
		reqUrl := serverUrl + serverVersionPath + "?name=" + projectName + "&no=" + nowTime + "&mac=" + mac
		if r, e := ghttp.Get(reqUrl); e != nil {
			glog.Error("[ConfigClient]reqUrl error:"+reqUrl, err)
		} else {
			defer r.Close()

			content := string(r.ReadAll())
			var respData resp.Resp
			err := json.Unmarshal([]byte(content), &respData)
			if err != nil {
				glog.Error("[ConfigClient]reqUrl resp to object error:"+reqUrl, err)
				continue
			}

			if !respData.Success() {
				glog.Error("[ConfigClient]projects config length error:" + reqUrl)
				continue
			}

			serverVersion = respData.GetString("version")
		}

		cacheVersion := client.GetCache(projectName).Version
		// 版本相同不需要处理
		if cacheVersion == serverVersion {
			continue
		}

		// 更新数据
		listBean := new(client.ConfigListBean)
		nowTime = gtime.Now().Format("YmdHis")
		mac, err = gmd5.Encrypt(projectName + nowTime + projectSecret)
		if err != nil {
			glog.Error("[ConfigClient]updateData md5 err", err)
			continue
		}
		reqUrl = serverUrl + serverDataPath + "?name=" + projectName + "&no=" + nowTime + "&mac=" + mac
		if r, e := ghttp.Get(reqUrl); e != nil {
			glog.Error("[ConfigClient]reqUrl error:"+reqUrl, err)
		} else {
			defer r.Close()

			content := string(r.ReadAll())
			glog.Info("[ConfigClient]reqUrl success", reqUrl, content)
			var respData resp.Resp
			err := json.Unmarshal([]byte(content), &respData)
			if err != nil {
				glog.Error("[ConfigClient]reqUrl resp to object error:"+reqUrl, err)
				continue
			}

			if !respData.Success() {
				glog.Error("[ConfigClient]projects config length error:" + reqUrl)
				continue
			}

			dataContent := respData.GetString("content")
			dataVersion := respData.GetString("version")

			var dataList []client.ConfigBean
			err = gjson.DecodeTo(dataContent, &dataList)
			if err != nil {
				glog.Error("[ConfigClient]reqUrl resp to object error:"+reqUrl, err)
				continue
			}

			listBean.Version = dataVersion
			listBean.Name = projectName
			listBean.ListData = dataList

			// 存储文件
			data := projectName + "\r\n" + dataVersion + "\r\n" + dataContent
			dataFilePath := GetDataPath() + gfile.Separator + projectName + ".txt"
			gfile.PutBytes(dataFilePath, []byte(data))

			// 设置缓存
			client.SetCache(*listBean)

			glog.Info("[ConfigClient]load cache and file:" + dataFilePath)
		}

	}

}

// 缓存数据初始化
func InitConfigData() {
	glog.Info("[ConfigClient]initConfigData start...")
	dataPath := GetDataPath()

	glog.Info("[ConfigClient]dataPath:" + dataPath)

	// 获取文件列表
	files, err := gfile.DirNames(dataPath)
	if err != nil {
		glog.Error("[ConfigClient]dataPath list file error", err)
		return
	}

	// 如果第一次不存在，进行一次获取
	if len(files) <= 0 {
		CheckVersion()
	}

	// 再次获取文件列表，不存在打印异常
	files, err = gfile.DirNames(dataPath)
	if err != nil {
		glog.Error("[ConfigClient]dataPath list file error", err)
		return
	}

	if len(files) <= 0 {
		glog.Error("[ConfigClient]dataPath not exist file")
		return
	}

	for _, filename := range files {
		filePath := dataPath + gfile.Separator + filename
		if !gfile.IsFile(filePath) {
			continue
		}

		content := gfile.GetContents(filePath)
		if content == "" {
			glog.Error("[ConfigClient]file content empty :" + filePath)
			continue
		}
		glog.Info(content)
		lines := gstr.Split(content, "\r\n")
		if len(lines) != 3 {
			glog.Error("[ConfigClient]file content error :" + filePath)
			continue
		}

		var dataList []client.ConfigBean
		err = gjson.DecodeTo(lines[2], &dataList)
		if err != nil {
			glog.Error("[ConfigClient]reqUrl resp to object error", err)
			continue
		}

		listBean := new(client.ConfigListBean)
		listBean.Name = lines[0]
		listBean.Version = lines[1]
		listBean.ListData = dataList
		// 设置缓存
		client.SetCache(*listBean)
	}

}

// 获取数据目录
func GetDataPath() string {
	dataPath := g.Config().GetString(constant.ParamDataPath)
	if dataPath == "" {
		dataPath = gfile.TempDir() + gfile.Separator + "configClient"
	}

	if !gfile.Exists(dataPath) {
		err := gfile.Mkdir(dataPath)
		if err != nil {
			glog.Error("[ConfigClient]create config path error")
		}
	}

	return dataPath
}
