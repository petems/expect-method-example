package main

import (
	"golang.org/x/crypto/ssh/terminal"
)

func getPassword(fd int) string {
	bytePassword, _ := terminal.ReadPassword(fd)
	return string(bytePassword)
}
