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



# 二、 部署相关
`go build -ldflags "-H=windowsgui" -o goexcel.exe`

1、完全杀死 nginx 命令 
`taskkill /f /t /im nginx.exe`

2、完全杀死 goexcel.exe 命令
`taskkill /f /t /im goexcel.exe`



`CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o goexcel.exe`

重新启动

1、双击 nginx.exe 看到一个小黑窗一闪而过，启动成功

2、窗口中 输入cmd 回车 
再弹出的命令行窗口中输入 : 
`goexcel.exe -d true`
 后回车


2、安装win- 安装supervisor

参考网站
https://blog.csdn.net/dszgf5717/article/details/132195625


`pip3 install supervisor-win`

导出 修改参数
`echo_supervisord_conf.exe > D:\temp\supervisord.conf`


修改文件为
已复制到本目录下了
`supervisorctl.exe status/restart`


4.1 使用CMD启动并在后台运行
start /B supervisord -c D:\temp\supervisord.conf

开机启动 创建 VBS 脚本
`
Set ws = CreateObject("Wscript.Shell")
ws.run "cmd /c supervisord -c D:\Python39\etc\supervisord.conf",vbhide
`

vbs脚本双击即可自动运行，设置开机自启动：打开目录 C:\ProgramData\Microsoft\Windows\Start Menu\Programs\StartUp，拷贝 start.vbs 到文件夹里面即可。


4.3 将supervisor安装成服务的形式（必须使用绝对路径）
python -m supervisor.services install -c C:\python\etc\supervisord.conf

`
[program:test-example]
command=D:\\Python39\\python task
directory=H:\\Project\\test
autostart=true
autorestart=true
startsecs=3
startretries=10
stopasgroup=true
redirect_stderr=true
environment = PYTHONUNBUFFERED=1,PYTHONIOENCODING="UTF-8"
stdout_logfile=H:\\Project\\test\\log\\%(program_name)s.log
stdout_logfile_maxbytes=10MB
stdout_logfile_backups=5
`

