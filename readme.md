# 一、 部署步骤

1、 安装环境
    包括数据库 mysql redis 等。

    drop database haibingoffice;
    create database haibingoffice;

2、 修改 env 文件
    主要是数据库链接 地址 端口 密码等。

3、 执行命令导入表结构
    `go run main.go migrate up`

4、 执行命令 seed 预定义数据
    `go run main.go seed`

5、 导入商品
    ``



