package main

import (
	"fmt"
	"testing"
)

func TestSendMail(t *testing.T) {

	mailConf := MailConf{
		MailConn: map[string]string{
			"user": "wanglulu@**.com",
			"pass": "Wang**.",
			"host": "smtp.exmail.qq.com",
			"port": "465",
		},
		MailTo: map[string]string{
			"wanglulu@***.com": "露露",
		},
		MailCc: map[string]string{
			"wanglulu@***.com": "露露",
		},
		MailBcc: map[string]string{
			"wanglulu@***.com": "露露",
		},
		Alias:   "from",
		Body:    "<h1>正文</h1>",
		Subject: "邮件主题",
		Attachs: []string{"/Users/lulu/workspace/dev/go/src/work/lumail/main.go"},
		Html:    true,
	}

	if err := SendMail(&mailConf); err != nil {
		fmt.Println("Error ", err)
	}

}