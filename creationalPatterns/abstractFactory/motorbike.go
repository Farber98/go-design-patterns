package abstractFactory

type Motorbike interface {
	GetMotorbikeType() int
}

type LuxuryMotorbike struct {
}

func (*LuxuryMotorbike) NumWheels() int {
	return 2
}
func (*LuxuryMotorbike) NumSeats() int {
	return 1
}
func (*LuxuryMotorbike) GetMotorbikeType() int {
	return LUXURY_MOTORBIKE
}

type FamilyMotorbike struct {
}

func (*FamilyMotorbike) NumWheels() int {
	return 2
}
func (*FamilyMotorbike) NumSeats() int {
	return 2
}
func (*FamilyMotorbike) GetMotorbikeType() int {
	return FAMILY_MOTORBIKE
}
