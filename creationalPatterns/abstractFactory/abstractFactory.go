package abstractFactory

import (
	"errors"
	"fmt"
)

/*
ABSTRACT FACTORY:
- Grouping related families of objects is very convenient when your object number is growing too much.
- Provides a new layer of encapsulation for Factory methods that return a common interface for all factories.
- Group common factories into a super Factory (factory of factories)

EXAMPLE: Vehicle factory example
- We must retrieve a Vehicle object using a factory returned by the abstract factory.
- The vehicle must be a concrete implementation of a Motorbike or a Car that implements both interfaces (Vehicle and Car or Motorbike)
*/
const (
	CAR_FACTORY_TYPE       = 1
	MOTORBIKE_FACTORY_TYPE = 2
)

func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CAR_FACTORY_TYPE:
		return new(CarFactory), nil
	case MOTORBIKE_FACTORY_TYPE:
		return new(MotorbikeFactory), nil
	default:
		return nil, errors.New(fmt.Sprintf("Factory with id %d not recognized \n", f))
	}
}
