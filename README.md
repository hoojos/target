# target简介

target是一个由Go开发高性能 HTTP API mock 服务器

## 主要功能：

- 支持返回头、返回体、返回耗时自定义

## 安装

```bash
git clone https://github.com/hoojos/target.git
cd target
go build -o target main.go
```

## 用法
1. 创建接口配置文件，文件格式参考`target.yaml`
2. 执行`./target -c ${filename}`