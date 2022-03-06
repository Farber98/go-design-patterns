package main

import (
	"fmt"
	"time"
)

type Command interface {
	Info() string
}

type TimePassed struct {
	Start time.Time
}

func (t *TimePassed) Info() string {
	return time.Since(t.Start).String()
}

type HelloMessage struct {
}

func (h HelloMessage) Info() string {
	return "Hello World!"
}

func main() {
	timeCommand := &TimePassed{time.Now()}
	HelloCommand := &HelloMessage{}
	time.Sleep(time.Second)
	fmt.Println(timeCommand.Info())
	fmt.Println(HelloCommand.Info())
}
