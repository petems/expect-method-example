package main

import (
	"testing"

	"github.com/Netflix/go-expect"
)

func TestGetPassword(t *testing.T) {
	c, _ := expect.NewConsole()
	defer c.Close()
	donec := make(chan struct{})
	go func() {
		defer close(donec)
		c.SendLine("hunter2")
	}()
	echoText := getPassword(int(c.Tty().Fd()))
	<-donec

	expected := "hunter2"
	actual := echoText

	if actual != expected {
		t.Errorf("Expected getPassword function to be `%s` instead got `%s`!", expected, actual)
	}
}
