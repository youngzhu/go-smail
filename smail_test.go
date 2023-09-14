package smail

import "testing"

func TestSendMail(t *testing.T) {
	SendMail("Test", "Say hi")
}
