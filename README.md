# gcsc

#### 介绍
gcsc(go config server client) 配置管理客户端,此项目基于gf框架开发，主要针对配置管理平台（https://github.com/goflyfox/gcs）项目的客户端实现；

* github地址：https://github.com/goflyfox/gcsc
* gitee地址：https://gitee.com/goflyfox/gcsc

#### 安装教程

* gopath模式: `go get github.com/goflyfox/gcsc`
* 或者 使用go.mod添加 :`require github.com/goflyfox/gcsc latest`

#### 功能模块

1. 配置管理平台客户端实现，仅需使用go mod引用加载初始化即可。
2. 实现客户端动态加载服务端发布数据，客户端无需感知自动更新。
3. 调用方法简单易用，等于与本地调用。
4. 支持本地缓存，参数频繁调用无性能瓶颈。
5. 支持重启本地加载，再服务端异常情况下不影响客户端可用性。
6. 支持多项目配置加载及按项目配置获取，便于配置维护。

#### 使用说明

1. 此项目为配置管理平台（https://github.com/goflyfox/gcs）项目的客户端实现
2. 需要配置配置管理平台服务端地址，项目名称，秘钥以及存储地址；
```toml
[config]
    # 服务端地址
    server-url = "http://127.0.0.1"
    # 项目名称，支持多个项目，用逗号分割
    project-name = "test"
    # 项目秘钥，与项目名称对应
    project-secret = "12345678"
    # 数据存储位置，用于启动加载，避免数据丢失
    data-path = "data"
```
3. 启动需要进行初始化;
`_ "github.com/goflyfox/gcsc/task"`
4. 直接调用即可获取到配置管理平台数据，使用示例如下：
```go
// 可获取到配置管理平台test对应的值
client.Value("test")
```

#### 感谢

1. gf框架 [https://github.com/gogf/gf](https://github.com/gogf/gf) 

#### 项目支持

- 项目的发展，离不开大家得支持~！~

- [阿里云最新活动：双11最新活动，低至1折；还有新人礼包；请点击这里](https://www.aliyun.com/1111/2019/home?spm=5176.11533457.1089570.70.4fe277e3TKVLoB&userCode=c4hsn0gc)
- 1核2G1M40G盘，86元/1年， 
- 2核4G3M40G盘，799元/3年，
- 2核8G5M40G盘，1399元/3年。

- [阿里云：ECS云服务器2折起；请点击这里](https://www.aliyun.com/acts/limit-buy?spm=5176.11544616.khv0c5cu5.1.1d8e23e8XHvEIq&userCode=c4hsn0gc)
- [阿里云：ECS云服务器新人优惠券；请点击这里](https://promotion.aliyun.com/ntms/yunparter/invite.html?userCode=c4hsn0gc)

- 也可以请作者喝一杯咖啡:)

![jflyfox](https://raw.githubusercontent.com/jflyfox/jfinal_cms/master/doc/pay01.jpg "Open source support")

