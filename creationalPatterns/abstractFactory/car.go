package abstractFactory

type Car interface {
	NumDoors() int
}

type LuxuryCar struct {
}

func (*LuxuryCar) NumWheels() int {
	return 4
}
func (*LuxuryCar) NumSeats() int {
	return 4
}
func (*LuxuryCar) NumDoors() int {
	return 5
}

type FamilyCar struct {
}

func (*FamilyCar) NumWheels() int {
	return 4
}
func (*FamilyCar) NumSeats() int {
	return 5
}
func (*FamilyCar) NumDoors() int {
	return 5
}
