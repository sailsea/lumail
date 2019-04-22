package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"gopkg.in/gomail.v2"
	"regexp"
	"strconv"
	"strings"
)

var conn = flag.String("conn", "", "必填 连接信息: -conn=用户名:密码[服务器:端口]")
var to = flag.String("to", "", `必填 用户列表,格式 -to=邮箱地址:别名,邮箱地址:别名`)
var cc = flag.String("cc", "", `用户列表,格式 -cc=邮箱地址:别名,邮箱地址:别名`)
var bcc = flag.String("bcc", "", `用户列表,格式 -bcc=邮箱地址:别名,邮箱地址:别名`)
var att = flag.String("att", "", "附件列表,格式 -att=/123.png,/tmp/456.pdf")
var sub = flag.String("sub", "Test标题", "邮件标题,格式 -sub=邮件标题")
var body = flag.String("body", "Test正文", "邮件正文,格式 -body=邮件内容")
var ali = flag.String("ali", "TestAlias", "From 别名，只支持英文,格式 -ali=lulu")
var html = flag.Bool("html", false, "true启用html格式，默认false,格式 -html")


func main() {
	flag.Parse()
	fmt.Println("-conn:", *conn) //-conn='用户名:密码[服务器:端口]'
	fmt.Println("-to:", *to)     //-to='邮箱地址:别名,邮箱地址:别名'
	fmt.Println("-cc:", *cc)     //-cc='邮箱地址:别名,邮箱地址:别名'
	fmt.Println("-bcc:", *bcc)   //-bcc='邮箱地址:别名,邮箱地址:别名'
	fmt.Println("-att:", *att)   //-att='/123.png,/tmp/456.pdf'
	fmt.Println("-sub:", *sub)   //-sub='邮件标题'
	fmt.Println("-body:", *body) //-body='邮件内容'
	fmt.Println("-ali:", *ali)   //-ali='lulu'
	fmt.Println("-html:", *html) //-html
	fmt.Println("其他参数：", flag.Args())
	var mailConf MailConf
	mc, err := MailConfHandle(mailConf)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	mailJson, _ := json.Marshal(mc)
	fmt.Println(string(mailJson))

	if err := SendMail(mc); err != nil {
		fmt.Println("发送失败,Error", err)
		return
	}
	fmt.Println("发送成功")
}

func MailConfHandle(mailConf MailConf) (*MailConf, error) {
	//处理conn
	conn := *conn
	if conn == "" {
		return &mailConf, errors.New("conn 不能为空.")
	}
	var connMap = make(map[string]string, 0)
	connRegexp := regexp.MustCompile(`\[(.*?)\]`)
	serverInfoList := connRegexp.FindStringSubmatch(conn)
	userInfo := strings.Replace(conn, serverInfoList[0], "", -1)
	serverList := strings.Split(serverInfoList[1], ":")
	userList := strings.Split(userInfo, ":")
	if len(serverList) != 2 || len(userList) != 2 {
		return &mailConf, errors.New("conn 格式错误,conn:" + conn)
	}
	connMap["user"] = userList[0]
	connMap["pass"] = userList[1]
	connMap["host"] = serverList[0]
	connMap["port"] = serverList[1]
	mailConf.MailConn = connMap

	//处理to -to=邮箱地址:别名;邮箱地址,别名
	to := *to
	if to != "" {
		toConf, err := StrToMap(to)
		if err != nil {
			return &mailConf, errors.New(fmt.Sprintf("to %v,to: %s", err, to))
		}
		mailConf.MailTo = toConf
	}

	//处理cc -cc=邮箱地址:别名;邮箱地址,别名
	cc := *cc
	if cc != "" {
		ccConf, err := StrToMap(cc)
		if err != nil {
			return &mailConf, errors.New(fmt.Sprintf("cc %v,cc: %s", err, cc))
		}
		mailConf.MailCc = ccConf
	}

	//处理to -to=邮箱地址:别名;邮箱地址,别名
	bcc := *bcc
	if bcc != "" {
		bccConf, err := StrToMap(bcc)
		if err != nil {
			return &mailConf, errors.New(fmt.Sprintf("bcc %v,bcc: %s", err, bcc))
		}
		mailConf.MailBcc = bccConf

	}

	//处理att -att=/123.png,/tmp/456.pdf
	att := *att
	if att != "" {
		var attList []string
		attList = strings.Split(att, ",")
		mailConf.Attachs = attList
	}

	//处理sub
	sub := *sub
	if sub != "" {
		mailConf.Subject = sub
	}

	//处理body
	body := *body
	if body != "" {
		mailConf.Body = body
	}

	//处理ali
	ali := *ali
	if ali != "" {
		mailConf.Alias = ali
	}

	//处理html
	html := *html
	mailConf.Html = html

	return &mailConf, nil
}

func StrToMap(string2 string) (map[string]string, error) {
	var m = make(map[string]string, 0)

	strList := strings.Split(string2, ",")
	for _, v := range strList {
		strList := strings.Split(v, ":")
		if len(strList) != 2 {
			return m, errors.New("格式错误")
		}
		m[strList[0]] = strList[1]
	}
	return m, nil
}

type MailConf struct {
	MailConn map[string]string //服务器连接信息和用户名密码
	MailTo   map[string]string //用户发送列表
	MailCc   map[string]string //用户抄送列表
	MailBcc  map[string]string //用户暗送列表
	Subject  string            //标题
	Body     string            //正文
	Alias    string            //发送者别名 英文
	Attachs  []string          //附件列表 绝对路径
	Html     bool              //是否发送html
}

func SendMail(mc *MailConf) error {
	//定义邮箱服务器连接信息，如果是阿里邮箱 pass填密码，qq邮箱填授权码
	mailConn := mc.MailConn
	//转换端口类型为int
	port, err := strconv.Atoi(mailConn["port"])
	if err != nil {
		return errors.New("端口格式错误.")
	}
	m := gomail.NewMessage()
	//如alias不为空设置发送者别名
	if mc.Alias != "" {
		m.SetHeader("From", mc.Alias+"<"+mailConn["user"]+">") //设置别名
	} else {
		m.SetHeader("From", mailConn["user"])
	}

	//设置发送给多个用户
	if len(mc.MailTo) != 0 {
		toList := []string{}
		for k, v := range mc.MailTo {
			toList = append(toList, m.FormatAddress(k, v))
		}
		m.SetHeader("To", toList...)
	}

	//设置抄送给多个用户
	if len(mc.MailCc) != 0 {
		ccList := []string{}
		for k, v := range mc.MailCc {
			ccList = append(ccList, m.FormatAddress(k, v))
		}
		m.SetHeader("Cc", ccList...)
	}

	//设置暗送给多个用户
	if len(mc.MailBcc) != 0 {
		bccList := []string{}
		for k, v := range mc.MailBcc {
			bccList = append(bccList, m.FormatAddress(k, v))
		}
		m.SetHeader("Bcc", bccList...)
	}

	//设置邮件主题
	if mc.Subject != ""{
		m.SetHeader("Subject", mc.Subject)
	}

	//设置邮件正文
	if mc.Html {
		m.SetBody("text/html", mc.Body)
	} else {
		m.SetBody("text/plain", mc.Body)
	}

	//设置附件
	if len(mc.Attachs) != 0 {
		for _, v := range mc.Attachs {
			m.Attach(v)
		}
	}

	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])

	return d.DialAndSend(m)
}
