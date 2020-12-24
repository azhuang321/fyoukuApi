# 微服务 服务发现(go version <= 1.14)
## 1.安装etcd(3.4.7)
   地址:https://github.com/etcd-io/etcd/releases  
   下载后直接运行 目录下的etcd

## 安装go-micro
    1. go get github.com/micro/go-micro
    2. go get github.com/micro/protobuf/{proto,proto-gin-go}
    2.go get github.com/micro/protoc-gen-micro/v2
    3. go get go get github.com/micro/protoc-gin-micro
    4.生成相关protobuf protoc -proto_path=xxx -go_out=xxx -micro_out=xxx xxx.proto

## 安装micro
    1. go get 安装或二进制包安装(2.x)
    2. go get github.com/micro/micro
    3.export MICRO_REGISTRY=etcd
    
 ## 启动micro api
    `micro api --address=0.0.0.0:8085 --handler=api`

## 启动micro web
    `micro web`
    
## 启动所有服务

```shell script
    micro api --address=0.0.0.0:8085 --handler=api
    micro web
    etcd
    micro main
    api main

```
