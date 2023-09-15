package smail

import "testing"

func TestSendMail(t *testing.T) {
	SendMail("Test", "Say hi")
}

func TestGetSmtpConfig(t *testing.T) {
	want := "smtp.office365.com"

	got := getSmtpConfig("xx@hotmail.com").host

	if got != want {
		t.Errorf("want %s, but got %s", want, got)
	}
}
