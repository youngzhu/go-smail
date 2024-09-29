package smail

import "testing"

func TestSendMail(t *testing.T) {
	err := SendMail("Test", "Say hi")
	if err != nil {
		t.Error(err)
	}
}

func TestGetSmtpConfig(t *testing.T) {
	want := "smtp.office365.com"

	got := getSmtpConfig("xx@hotmail.com").host

	if got != want {
		t.Errorf("want %s, but got %s", want, got)
	}
}
