package emailutil

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/dajinkuang/errors"
)

// SendTlsEmail tls 发送邮件
func SendTlsEmail(host string, port int, fromEmailAccount, fromEmailAccountName, fromEmailPwd, toEmailAccount, content, subject, contentType string) (err error) {
	header := make(map[string]string)

	header["From"] = fromEmailAccountName + "<" + fromEmailAccount + ">"
	header["To"] = toEmailAccount + "<" + toEmailAccount + ">"
	header["Subject"] = subject
	header["Content-Type"] = contentType
	message := ""

	for k, v := range header {
		message += fmt.Sprintf("%s:%s\r\n", k, v)
	}

	message += "\r\n" + content

	auth := smtp.PlainAuth(
		"",
		fromEmailAccount,
		fromEmailPwd,
		host,
	)

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		err = errors.Wrap(err)
		return
	}
	// create smtp client
	c, err := smtp.NewClient(conn, host)
	if err != nil {
		err = errors.Wrap(err)
		return
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				err = errors.Wrap(err)
				return
			}
		}
	}

	if err = c.Mail(fromEmailAccount); err != nil {
		err = errors.Wrap(err)
		return
	}

	if err = c.Rcpt(toEmailAccount); err != nil {
		err = errors.Wrap(err)
		return
	}

	w, err := c.Data()
	if err != nil {
		err = errors.Wrap(err)
		return
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		err = errors.Wrap(err)
		return
	}

	err = w.Close()
	if err != nil {
		err = errors.Wrap(err)
		return
	}

	err = c.Quit()
	if err != nil {
		err = errors.Wrap(err)
		return
	}
	return
}
