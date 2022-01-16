package emailutil

import (
	"net/smtp"
	"strings"

	"github.com/dajinkuang/errors"
)

// SendToMail 发送邮件 smtp
//  user       发件人邮箱账号
//  password   发件人邮箱密码
//  host       发件人服务器
//  to         收件人
//  subject    邮件主题
//  body       邮件内容
//  emailType  邮件类型 传html
func SendToMail(user, password, host, to, subject, body, emailType string) (err error) {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var contentType string
	switch strings.ToLower(emailType) {
	case "html":
		contentType = "Content-Type: text/" + emailType + "; charset=UTF-8"
	default:
		contentType = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	sendTo := strings.Split(to, ";")
	if err = smtp.SendMail(host, auth, user, sendTo, msg); err != nil {
		err = errors.Wrap(err)
		return
	}
	return
}
