# QAsystem

仿知乎问答系统服务端

目前正在完善中。。。

## Docker 快速部署

```shell
$ docker-compose up
```

## 常规运行

### 1.获取代码

```shell
$ git clone https://github.com/SSunSShine/QAsystem

$ cd QAsystem
```

### 2.下载依赖

```shell
$ go mod tidy
```

### 3.修改配置信息
```shell
$ touch ./conf/configuration.yaml
```

```yaml
# 数据库
db:
  driver: mysql
  addr: root:123456@/qasystem?charset=utf8&parseTime=True&loc=Local

# Redis
redis:
  addr: 127.0.0.1:6379
  password:
  db: 0

# jwt认证密钥
jwtKey: 自己设置密钥

# 端口
address: :8080
```

### 4.初始化并运行

```shell
$ (sudo) go run ./ -init
```

## 技术栈

| 技术              | 链接                                                         |
| ----------------- | ------------------------------------------------------------ |
| gin               | https://github.com/gin-gonic/gin                             |
| gorm              | https://github.com/jinzhu/gorm                               |
| viper             | https://github.com/spf13/viper                               |
| go-redis          | https://github.com/go-redis/redis/v8                         |
| validator         | https://github.com/go-playground/validator/v10               |
| jwt-go            | https://github.com/dgrijalva/jwt-go                          |
| bcrypt            | https://golang.org/x/crypto/bcrypt                           |
| httpexpect        | https://github.com/gavv/httpexpect/v2                        |
| docker            | https://docs.docker.com/                                     |
