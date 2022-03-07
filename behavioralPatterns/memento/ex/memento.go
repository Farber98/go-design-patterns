package main

import "fmt"

/*
The idea is to use a Command pattern to encapsulate a set of
different types of states (those that implement a Command interface)
and provide a small facade to automate the insertion in the CareTaker object.

We are going to develop a small ex. of an hypothetical audio mixer,
using the same Memento Pattern to save two types of states: Volume and Mute.
The volume state is oging to be a byte type and the Mute state a Boolean type.

We will use two completely different types to show the flex of this approach (and its drawbacks)

*/

type Command interface {
	GetValue() interface{}
}

type Volume byte

func (v Volume) GetValue() interface{} {
	return v
}

type Mute bool

func (m Mute) GetValue() interface{} {
	return m
}

type Memento struct {
	Memento Command
}

type Originator struct {
	Command Command
}

func (o *Originator) NewMemento() Memento {
	return Memento{o.Command}
}

func (o *Originator) ExtractAndStoreCommand(m Memento) {
	o.Command = m.Memento
}

type CareTaker struct {
	MementoStack []Memento
}

func (c *CareTaker) Add(m Memento) {
	c.MementoStack = append(c.MementoStack, m)
}

func (c *CareTaker) Pop() Memento {
	if len(c.MementoStack) > 0 {
		tempMemento := c.MementoStack[len(c.MementoStack)-1]
		c.MementoStack = c.MementoStack[0 : len(c.MementoStack)-1]
		return tempMemento
	}
	return Memento{}
}

type MementoFacade struct {
	Originator Originator
	CareTaker  CareTaker
}

func (m *MementoFacade) SaveSettings(s Command) {
	m.Originator.Command = s
	m.CareTaker.Add(m.Originator.NewMemento())
}

func (m *MementoFacade) RestoreSettings(i int) Command {
	m.Originator.ExtractAndStoreCommand(m.CareTaker.MementoStack[i])
	return m.Originator.Command
}

func main() {
	m := MementoFacade{}
	m.SaveSettings(Volume(4))
	m.SaveSettings(Mute(false))
	AssertAndPrint(m.RestoreSettings(0))
	AssertAndPrint(m.RestoreSettings(1))
}

func AssertAndPrint(c Command) {
	switch cast := c.(type) {
	case Volume:
		fmt.Printf("Volume:\t%d\n", cast)
	case Mute:
		fmt.Printf("Mute:\t%t\n", cast)
	}
}
