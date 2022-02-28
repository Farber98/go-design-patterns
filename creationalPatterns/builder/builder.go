package builder

/*
BUILDER:
- Abstract complex creations so that the object creation is separated from the object user.
- Create an object step by step by filling its fields and creating embedded objects.
- Reuse the object creation algorithm between many objects that share the same characteristic.

WHEN TO USE IT:
- Need to create similar objects with small diferences.

EXAMPLE: Vehicle manufacturing.
- The process of creating a vehicle is more or less the same for every kind of vehicle.
- We must have a manufacturing type that constructs everything that a vehicle needs.
- When using a car builder, the VehicleProduct with four wheels, five seats and a structure definde as Car must be returned
- When using a motorbike builder, the VehicleProduct with two wheels, two seats and a structure defined as Motorbike must be returned.
- A VehicleProduct built by any BuildProcess bulder must be open to modifications.

*/

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

/* In charge of accepting builders.  */
type ManufacturingDirector struct {
	builder BuildProcess
}

/* Uses the builder stored in Manufacturing, and reproduce required stepts.  */
func (f *ManufacturingDirector) Construct() {
	f.builder.SetWheels().SetSeats().SetStructure()
}

/* Allows us to change the builder that is b eing used in Manufacturing Director. */
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

/* CAR BUILDER */
type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

/* MOTORBIKE BUILDER */
type MotorbikeBuilder struct {
	v VehicleProduct
}

func (b *MotorbikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}
func (b *MotorbikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 1
	return b
}
func (b *MotorbikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}
func (b *MotorbikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

/* MOTORBIKE BUILDER */
type BusBuilder struct {
	v VehicleProduct
}

func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 8
	return b
}
func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 24
	return b
}
func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}
func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}
