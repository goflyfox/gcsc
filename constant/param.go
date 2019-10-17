package constant

const (
	ParamProjectName      = "config.project-name"       // 项目名称
	ParamProjectSecret    = "config.project-secret"     // 项目秘钥
	ParamDataPath         = "config.data-path"          // 数据存储路径
	ParamCronCheckVersion = "config.cron-check-version" // 定时任务配置

	// 服务端配置
	ParamServerUrl         = "config.server-url"       // 服务端URL
	ParamServerVersionPath = "config.url-version-path" // 版本获取路径
	ParamServerDataPath    = "config.url-data-path"    // 数据获取路径

	CacheKey = "GCSC:CONFIG" // 缓存Key
)
