package abstractFactory

import "testing"

func TestMotorbikeFactory(t *testing.T) {
	motorbikeFactory, err := BuildFactory(MOTORBIKE_FACTORY_TYPE)
	if err != nil {
		t.Fatal(err)
	}

	motorbikeVehicle, err := motorbikeFactory.NewVehicle(LUXURY_MOTORBIKE)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motorbike vehicle has %d wheels\n", motorbikeVehicle.NumWheels())

	luxuryBike, ok := motorbikeVehicle.(Motorbike)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Luxury motorbike has type %d\n", luxuryBike.GetMotorbikeType())
}

func TestCarFactory(t *testing.T) {
	carFactory, err := BuildFactory(CAR_FACTORY_TYPE)
	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := carFactory.NewVehicle(LUXURY_CAR)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car vehicle has %d wheels\n", carVehicle.NumWheels())

	luxuryCar, ok := carVehicle.(Car)
	if !ok {
		t.Fatal("Struct assertion has failed")
	}
	t.Logf("Luxury motorbike has type %d\n", luxuryCar.NumDoors())
}
