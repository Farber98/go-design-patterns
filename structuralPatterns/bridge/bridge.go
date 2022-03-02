package bridge

import (
	"errors"
	"io"
)

/*
BRIDGE:
- Decouples an abstraction from its implementation so that the two can vary indenepdently.
- Means you could even decouple an abstraction (eg. object) from what it does (eg. its implementation).
- Allows us to change the abstracted object while reusing the same implementation.

OBJECTIVE:
- Bring flexibiliy to a struct that changes often.
- Knowing the I/O of a method, it allows us to change code without knowing too much about it and leaving the freedom for both sides to be modified more easily

EXAMPLE: Two printers and two ways of printing for each

ACCEPTANCE CRITERIA:
- We will have two objects (Packt and Normal printer) and two implementations(PrinterImpl1 and PrinterImpl2) that we will join by using the Bridge design pattern.
- A PrinterAPI that accepts a message to print
- An implementation of the API that simply prints the message to the console
- An implementation of the API that prints to an io.Writer interface
- A Printer abstraction with a Print method to implement printing types.
- A Normal printer object, which will implement the Printer and the PrinterAPI interface.
- The normal printer will forward the message directly to the implementation.
- A Packt printer, which will implement the Printer abstraction and the PrinterAPI interface.
- The Packt printer will append the message Message from Packt: to all prints.

*/

type PrinterAPI interface {
	PrintMessage(string) error
}

type PrinterImpl1 struct {
}

func (impl1 *PrinterImpl1) PrintMessage(msg string) error {
	return errors.New("not implemented yet")
}

type PrinterImpl2 struct {
	Writer io.Writer
}

func (impl2 *PrinterImpl2) PrintMessage(msg string) error {
	return errors.New("not implemented yet")

}

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = errors.New("content received on Writer was empty")
	return
}

type PrinterAbstraction interface {
	Print() error
}

type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *NormalPrinter) Print() error {
	return errors.New("NotImplementedYet")
}

type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *PacktPrinter) Print() error {
	return errors.New("NotImplenetedYet")
}
