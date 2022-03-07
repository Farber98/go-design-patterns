package memento

import (
	"errors"
)

/*
MEMENTO:
- We'll have a type with some state and we want to be able to save milestons of its state.
- Having a finite amount of states saved, we can recoverr them if necesary for a variety of tasks-undo operations, historics, and so on.
- Usually has 3 actors:
	* Memento: A type that stores the type we want to save. Usually, we wont store the business type directly and we provide an extra layer of abstraction through this type.
	* Originator: A type that is in charge of creating mementos and storing the current active state. We said that the Memento type wraps states of the business type and we use originator as the creator of mementos.
	* Care taker: A type that stores the list of mementos that can have the logic to store them in a database or to not store more than specified number of them

OBJECTIVE:
- Memento is all about a sequence of actions over time, used to:
	* Capture an object state without modifying the object itself
	* Save limited amount of states so we can retrieve them later.

EXAMPLE: strings.
- Use a string as the state we want to save.
- The string stored in a field of a State instance, will be modified and we will be able to undo the operations done in this state.

ACCEPTANCE CRITERIA:
- We need to store a finite amount of states of type string
- We need a way to restore the current stored state to one of the state list.
*/

type State struct {
	Description string
}

type Memento struct {
	State State
}

type Originator struct {
	State State
}

func (o *Originator) NewMemento() Memento {
	return Memento{State: o.State}
}

func (o *Originator) ExtractAndStoreState(m Memento) {
	o.State = m.State
}

type CareTaker struct {
	MementoList []Memento
}

func (c *CareTaker) Add(m Memento) {
	c.MementoList = append(c.MementoList, m)
}

func (c *CareTaker) Memento(i int) (Memento, error) {
	if len(c.MementoList) < i || i < 0 {
		return Memento{}, errors.New("index not found")
	}
	return c.MementoList[i], nil
}
