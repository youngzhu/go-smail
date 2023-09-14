package smail

import (
	"crypto/tls"
	"fmt"
	gomail "gopkg.in/gomail.v2"
	"log"
	"os"
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

	m.SetHeader("Subject", subject)

	m.SetBody("text/plain", body)

	d := gomail.NewDialer("smtp.163.com", 465, from, fromPwd)
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
