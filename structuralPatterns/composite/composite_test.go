package composite

import (
	"fmt"
	"testing"
)

/* COMPOSITE */
func TestAthleteComposite(t *testing.T) {
	localSwim := Swim
	swimmer := CompositeSwimmerA{
		MySwim: &localSwim,
	}
	(*swimmer.MySwim)()
	swimmer.MyAthlete.Train()
}

/* EMBED */
func TestAnimalEmbed(t *testing.T) {
	shark := Shark{
		Swim: Swim,
	}

	shark.Eat()
	shark.Swim()

}

/* INTERFACE */
func TestAthleteInterface(t *testing.T) {
	swimmer := SwimmerTrainer{
		&SwimmerImpl{},
		&Athlete{},
	}
	swimmer.Swim()
	swimmer.Train()
}

func TestTreeRecursiveComposite(t *testing.T) {
	root := Tree{
		LeafValue: 0,
		Right:     &Tree{5, &Tree{6, nil, nil}, nil},
		Left:      &Tree{4, nil, nil},
	}
	fmt.Println(root.Right.Right.LeafValue)
}
func TestTreeCompositeVsInheritance(t *testing.T) {
	son := Son{}
	GetParentField(&son.P)
}
