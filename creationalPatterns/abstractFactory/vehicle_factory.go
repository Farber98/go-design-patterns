package abstractFactory

import (
	"errors"
	"fmt"
)

type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

const (
	LUXURY_CAR       = 1
	FAMILY_CAR       = 2
	LUXURY_MOTORBIKE = 1
	FAMILY_MOTORBIKE = 2
)

type CarFactory struct {
}

func (*CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LUXURY_CAR:
		return new(LuxuryCar), nil
	case FAMILY_CAR:
		return new(FamilyCar), nil
	default:
		return nil, errors.New(fmt.Sprintf("vehicle of type %d not recognized\n", v))
	}
}

type MotorbikeFactory struct {
}

func (*MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LUXURY_MOTORBIKE:
		return new(LuxuryMotorbike), nil
	case FAMILY_MOTORBIKE:
		return new(FamilyMotorbike), nil
	default:
		return nil, errors.New(fmt.Sprintf("vehicle of type %d not recognized\n", v))
	}
}
