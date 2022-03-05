package responsability

import (
	"fmt"
	"io"
	"strings"
)

/*
CHAIN OF RESPONSiBILITY:
- It consists of a chain where each link follows the single responsibility principle
- Implies that a type, function, method or any similar abstraction must have one single responsibility onli and it must do it quite well.
- This way, we can apply mani functions that achieve one specific thieng each to some struct, slice, map and so on.
- We can chain them to execute in order.

OBJECTIVE:
- Provide a way to chain actions at runtime, based on some input.
- Actions are chained to each other and each link will execute some action and pass the request to the next link (or not).
- Pass a request through a chain of processors untile one of them can proces it, in which case the chain could be stopped.

EXAMPLE: Multi-logger chain.
- Use two different loggers. Console loggers and one general purpose logger.

ACCEPTANCE CRITERIA:
- We need a simple logger that logs the text of a request with a prefix FirstLogger and passes it to the next link in the chain
- A second logger will write on the console if the incoming text has the word hello and pass the request to a third logger. If not, the chain will be broken and return.
- A third logger type is a general purpose logger called WriterLogger that useas an io.Writer interface to log
- A concrete implementation of the WriterLogger writes to a file and represents the third link in the chain
*/

type ChainLogger interface {
	Next(string)
}

type FirstLogger struct {
	NextChain ChainLogger
}

func (f *FirstLogger) Next(s string) {
	fmt.Printf("First logger: %s\n", s)
	if f.NextChain != nil {
		f.NextChain.Next(s)
	}
}

type SecondLogger struct {
	NextChain ChainLogger
}

func (f *SecondLogger) Next(s string) {
	if strings.Contains(strings.ToLower(s), "hello") {
		fmt.Printf("Second logger: %s\n", s)
		if f.NextChain != nil {
			f.NextChain.Next(s)
		}
		return
	}
	fmt.Printf("Finishing in second logging\n\n")
}

type WriterLogger struct {
	NextChain ChainLogger
	Writer    io.Writer
}

func (f *WriterLogger) Next(s string) {
	if f.Writer != nil {
		f.Writer.Write([]byte("WriterLogger: " + s))
	}
	if f.NextChain != nil {
		f.NextChain.Next(s)
	}
}

type ClosureChain struct {
	NextChain ChainLogger
	Closure   func(string)
}

func (c *ClosureChain) Next(s string) {
	if c.Closure != nil {
		c.Closure(s)
	}
	if c.NextChain != nil {
		c.Next(s)
	}
}
