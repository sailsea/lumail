# lumail
一个命令行邮件发送工具

#### 参数说明

|参数|必选|默认|注释|
|:---|:---|:---|:---|
|conn|是|""|必填 连接信息: -conn=用户名:密码[服务器:端口]|
|to|是|""|必填 用户列表,格式 -to=邮箱地址:别名,邮箱地址:别名|
|cc|否|""|用户列表,格式 -cc=邮箱地址:别名,邮箱地址:别名|
|bcc|否|""|用户列表,格式 -bcc=邮箱地址:别名,邮箱地址:别名|
|att|否|""|附件列表,格式 -att=/123.png,/tmp/456.pdf|
|sub|否|""|邮件标题,格式 -sub=邮件标题|
|body|否|""|邮件正文,格式 -body=邮件内容|
|ali|否|""|From 别名，只支持英文,格式 -ali=lulu|
|html|否|false|true启用html格式，默认false,格式 -html|

#### 使用方法

打开链接https://github.com/sailsea/datadict/releases/tag/v1.0.0
下载自己系统对应的版本
```
lumail-darwin-386.zip
lumail-darwin-amd64.zip
lumail-linux-386.zip
lumail-linux-amd64.zip
lumail-windows-386.zip
lumail-windows-amd64.zip
```

解压后放在环境变量的目录中

使用例子

配置必须项[-conn,-to]发送测试邮件(conn要加单引号)
```bash
lumail \
    -conn='wanglulu@**.com:ang1234.[smtp.exmail.qq.com:465]' \
    -to=wanglulu@**.com:露露,wanglulu@**.com:露露 
```

全配置发送邮件

```bash
lumail \
    -conn='wanglulu@***.com:Wang56688.[smtp.exmail.qq.com:465]' \
	-to=wanglulu@***.com:露露,wanglulu@***.com:露露 \
	-cc=wanglulu@***.com:露露,wanglulu@***.com:露露 \
	-bcc=wanglulu@***.com:露露,wanglulu@***.com:露露 \
	-att=/Users/lulu/workspace/dev/go/src/work/mail/main.go,/Users/lulu/workspace/dev/go/src/work/mail/gomail.go \
	-sub=subTest \
	-body='text<h1>h1</h1>  露露' \
	-ali=lulu \
	-html
```


执行效果图
![执行效果图](https://github.com/sailsea/lumail/master/image/3.png)
生成效果图
#### 必选项发送
![必选项发送](https://github.com/sailsea/lumail/master/image/1.png)
#### 全选发送
![全选发送](https://github.com/sailsea/lumail/master/image/2.png)