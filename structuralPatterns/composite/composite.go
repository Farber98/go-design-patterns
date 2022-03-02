package composite

import "fmt"

/*
COMPOSITE:
- Favors composition ('has a' relationship) over inheritance ('is a' relationship).
- Create hierarchies and trees of objects.
- Objects have different objects with their own fields and methods inside them.

OBJECTIVE:
- Avoid hierarchy hell where the complexity of an app could grow too much and the clarity of code is affected.

GOLANG:
- Direct composition.
- Embedding composition.

EXAMPLE: The swimmer and the fish.
- We'll have and athlete and a swimmer.
- We'll have an animal and a shark.
- The Swimmer and the Fish methods must share the code
- The athlete mus train and the animal must eat.

ACCEPTANEC CRITERIA:
- We must have an Athlete struct with Train method
- We must have a Swimmer with Swim method
- We must have an Animal struct with an Eat method
- We must have a shark struct with a swim method that is shared with the Swimmer, and not have inheritance or hierarchy issues.

*/

/* COMPOSITE */
type Athlete struct {
}

func (a *Athlete) Train() {
	fmt.Println("Training")
}

func Swim() {
	fmt.Println("Swimming")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    *func()
}

/* ---------------------------------------------------- */

/* EMBED */
type Animal struct{}

func (a *Animal) Eat() {
	fmt.Println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

/* ---------------------------------------------------- */
/* BEST APPROACH WITH INTERFACES */
type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}

type SwimmerImpl struct {
}

func (s *SwimmerImpl) Swim() {
	fmt.Println("Swimming")
}

/* type Athlete struct {
}

func (a *Athlete) Train() {
	fmt.Println("Training")
} */

type SwimmerTrainer struct {
	Swimmer
	Trainer
}

/* ---------------------------------------------------- */
/* RECURSIVE COMPOSITING */
type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}

/* ---------------------------------------------------- */
/* COMPOSITE PATTERN VERSUS INHERITANCE */
type Parent struct {
	SomeField int
}

type Son struct {
	P Parent // Composite son struct with the parent to access parent instance
}

func GetParentField(p *Parent) int {
	return p.SomeField
}
