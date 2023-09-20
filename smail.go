package smail

import (
	"crypto/tls"
	"fmt"
	gomail "gopkg.in/gomail.v2"
	"log"
	"os"
	"strings"
)

const (
	envKeyFrom    = "GO_MAIL_FROM"
	envKeyFromPwd = "GO_MAIL_FROM_PWD"
	envKeyTo      = "GO_MAIL_TO"

	secretErr = "请先设置变量[%s]\n"
)

func SendMail(subject, body string) error {
	from, err := getSecret(envKeyFrom)
	if err != nil {
		return err
	}
	fromPwd, err := getSecret(envKeyFromPwd)
	if err != nil {
		return err
	}
	to, err := getSecret(envKeyTo)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to)

	if subject == "" {
		return fmt.Errorf("邮件主题不能为空\n")
	}
	m.SetHeader("Subject", subject)

	if body == "" {
		body = "RT"
	}
	m.SetBody("text/plain", body)

	smtp := getSmtpConfig(from)

	d := gomail.NewDialer(smtp.host, smtp.port, from, fromPwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	if err := d.DialAndSend(m); err != nil {
		log.Print("发送邮件失败：", err)
	}

	return nil
}

func getSecret(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		log.Printf(secretErr, key)
		return "", fmt.Errorf(secretErr, key)
	}
	return val, nil
}

type smtpConfig struct {
	host string
	port int
}

var smtpConfigs = map[string]smtpConfig{
	"163.com": {
		host: "smtp.163.com",
		port: 465,
	},
	"hotmail.com": {
		host: "smtp.office365.com",
		port: 587,
	},
}

func getSmtpConfig(mail string) smtpConfig {
	if !strings.Contains(mail, "@") {
		panic("无效的邮箱地址:" + mail)
	}

	suffix := strings.Split(mail, "@")[1]

	smtp, ok := smtpConfigs[suffix]
	if !ok {
		panic("尚不支持的邮箱类型:" + mail)
	}
	return smtp
}
