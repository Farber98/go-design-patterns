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

type ChainLogger interface {
	Next(Command)
}

type Logger struct {
	NextChain ChainLogger
}

func (f *Logger) Next(c Command) {
	time.Sleep(time.Second)
	fmt.Printf("Elapsed time from creation: %s\n", c.Info())

	if f.NextChain != nil {
		f.NextChain.Next(c)
	}
}

func main() {
	second := new(Logger)
	first := Logger{second}
	command := &TimePassed{time.Now()}
	first.Next(command)
}
