package memento

import "testing"

func TestCareTaker_Add(t *testing.T) {
	originator := Originator{State{Description: "Idle"}}
	CareTaker := CareTaker{}
	mem := originator.NewMemento()
	if mem.State.Description != "Idle" {
		t.Error("Expected state was not found")
	}
	currentLen := len(CareTaker.MementoList)
	CareTaker.Add(mem)
	if len(CareTaker.MementoList) != currentLen+1 {
		t.Error("no new elements were added on the list")
	}
}

func TestCareTaker_Memento(t *testing.T) {
	originator := Originator{State{Description: "Idle"}}
	CareTaker := CareTaker{}
	CareTaker.Add(originator.NewMemento())
	mem, err := CareTaker.Memento(0)
	if err != nil {
		t.Fatal(err)
	}
	if mem.State.Description != "Idle" {
		t.Error("Unexpected state")
	}
	mem, err = CareTaker.Memento(-1)
	if err == nil {
		t.Fatal("An error is expected when asking for a negative number but no error was found")
	}
}

func TestOriginator_ExtractAndStoreState(t *testing.T) {
	originator := Originator{State: State{"Idle"}}
	idleMemento := originator.NewMemento()

	originator.ExtractAndStoreState(idleMemento)
	if originator.State.Description != "Idle" {
		t.Error("Unexpected state found")
	}
}
