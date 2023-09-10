# Virtual-Admin

简体中文

**暂未完成，无法使用**

基于Gin + Vue + Element UI 的前后端分离录播权限管理系统

本项目基于[Go-Admin](https://github.com/go-admin-team/go-admin)作为权限管理框架进行开发，感谢项目开发者的二次开发授权

[在线文档](https://vup.monster)

[前端项目](https://github.com/KiClover/virtual-frontend)


## 🎬 功能

- 以用户为单位粒度控制录播权限
- 多录播节点部署，管理便携
- 对接S3 Object Storage,BiliBili Upload API,阿里云盘API，支持边缘节点自动上传
- 对接腾讯云数据万象，支持COS存储策略进行云端基础切片
- 支持中心化存储，无需使用用户自行配置存储策略信息


## ✨ 特性

- 遵循 RESTful API 设计规范

- 基于 GIN WEB API 框架，提供了丰富的中间件支持（用户认证、跨域、访问日志、追踪ID等）

- 基于Casbin的 RBAC 访问控制模型

- JWT 认证

- 支持 Swagger 文档(基于swaggo)

- 基于 GORM 的数据库存储，可扩展多种类型数据库

- 多指令模式

- 多租户的支持

- TODO: 单元测试


## 准备工作

你需要在本地安装 go gin [node](http://nodejs.org/) 和 [git](https://git-scm.com/) 

同时配套了系列教程包含视频和文档，如何从下载完成到熟练使用，强烈建议大家先看完这些教程再来实践本项目！！！


## 📦 本地开发

### 环境要求

go 1.18

node版本: v14.16.0

npm版本: 6.14.11

### 开发目录创建

```bash

# 创建开发目录
mkdir virtualadmin
cd virtualadmin
```

### 获取代码

> 重点注意：两个项目必须放在同一文件夹下；

```bash
# 获取后端代码
git clone https://github.com/KiClover/virtual-backend.git

# 获取前端代码
git clone https://github.com/KiClover/virtual-frontend.git

```

### 启动说明

#### 服务端启动说明

```bash
# 进入 go-admin 后端项目
cd ./virtual-backend

# 更新整理依赖
go mod tidy

# 编译项目
go build

# 修改配置 
# 文件路径  go-admin/config/settings.yml
vi ./config/settings.yml

# 1. 配置文件中修改数据库信息 
# 注意: settings.database 下对应的配置数据
# 2. 确认log路径
```

:::tip ⚠️注意 在windows环境如果没有安装中CGO，会出现这个问题；

```bash
E:\virtual-admin>go build
# github.com/mattn/go-sqlite3
cgo: exec /missing-cc: exec: "/missing-cc": file does not exist
```

or

```bash
D:\Code\virtual-admin>go build
# github.com/mattn/go-sqlite3
cgo: exec gcc: exec: "gcc": executable file not found in %PATH%
```


:::

#### 初始化数据库，以及服务启动

``` bash
# 首次配置需要初始化数据库资源信息
# macOS or linux 下使用
$ ./virtual-admin migrate -c config/settings.dev.yml

# ⚠️注意:windows 下使用
$ virtual-admin.exe migrate -c config/settings.dev.yml


# 启动项目，也可以用IDE进行调试
# macOS or linux 下使用
$ ./virtual-admin server -c config/settings.yml


# ⚠️注意:windows 下使用
$ virtual-admin.exe server -c config/settings.yml
```

#### sys_api 表的数据如何添加

在项目启动时，使用`-a true` 系统会自动添加缺少的接口数据
```bash
./virtual-admin server -c config/settings.yml -a true
```

#### 使用docker 编译启动

```shell
# 编译镜像
docker build -t virtual-admin .

# 启动容器，第一个virtual-admin是容器名字，第二个go-admin是镜像名称
# -v 映射配置文件 本地路径：容器路径
docker run --name virtual-admin -p 8000:8000 -v /config/settings.yml:/config/settings.yml -d go-admin-server
```

#### 文档生成

```bash
go generate
```

#### 交叉编译

```bash
# windows
env GOOS=windows GOARCH=amd64 go build main.go

# or
# linux
env GOOS=linux GOARCH=amd64 go build main.go
```

### UI交互端启动说明

```bash
# 安装依赖
npm install

# 建议不要直接使用 cnpm 安装依赖，会有各种诡异的 bug。可以通过如下操作解决 npm 下载速度慢的问题
npm install --registry=https://registry.npmmirror.com

# 启动服务
npm run dev
```

