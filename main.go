package main

import (
  "bufio"
  "fmt"
  "io"
  "bytes"

  "time"
  expect "github.com/Netflix/go-expect"
)

func feeder(stdin io.Reader) (string) {

  scanner := bufio.NewScanner(stdin)
  var text string

  fmt.Print("Enter some text: ")
  scanner.Scan()
  text = scanner.Text()

  return text
}

func main() {
  buf := new(bytes.Buffer)
  c, _ := expect.NewConsole(expect.WithStdout(buf))

  defer c.Close()

  donec := make(chan struct{})
  go func() {
    defer close(donec)
    time.Sleep(time.Second)
    c.Send("iHello world")
    time.Sleep(time.Second)
    c.ExpectEOF()
  }()

  echoText := feeder(c.Tty())

  <-donec

  fmt.Print("Reponse", echoText)
}
