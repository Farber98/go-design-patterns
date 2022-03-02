package adapter

import "fmt"

/*
ADAPTER:
- Allows us to use something that wasn't built for a specific task at the beggining.
- Useful when an interface gets outdated and it's not possible to replace it easily or fast. Instead, you create a new interface to deal with the current needs of your app, which, under the hood, uses implementations of the old interface.

OBJECTIVE:
- Helps to fit the needs of two parts of the code that are incompatible at first and must work together.

EXAMPLE: Using an incompatible interface wih an adapter object.
- We will have an old Printer interface and a new one.
- Users of the new interface don't expect the signature that the old one has.
- We need an Adapter so that users can still se old implementatios if necessary.

ACCEPTANEC CRITERIA:
- Create an Adapter object that implements the ModernPrinter interface
- The new Adapter object must contain an instance of the LegacyPrinter interface
- When using ModernPrinter, it must call the LegacyPrinter interface under the hood, prefixing it with the text Adapter.

*/

type LegacyPrinter interface {
	Print(s string) string
}

type MyLegacyPrinter struct {
}

func (l *MyLegacyPrinter) Print(s string) (newMsg string) {
	newMsg = fmt.Sprintf("Legacy Printer: %s\n", s)
	fmt.Println(newMsg)
	return
}

type ModernPrinter interface {
	PrintStored() string
}

type MyModernPrinter struct {
}

type PrinterAdapter struct {
	OldPrinter LegacyPrinter
	Msg        string
}

func (p *PrinterAdapter) PrintStored() (newMsg string) {
	if p.OldPrinter != nil {
		newMsg = fmt.Sprintf("Adapter: %s", p.Msg)
		newMsg = p.OldPrinter.Print(newMsg)
	} else {
		newMsg = p.Msg
	}
	return
}
