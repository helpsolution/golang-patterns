package builder

type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehickleProduct
}

//Director
type ManufacturingDirector struct {
	builder BuildProcess
}

func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

//Product
type VehickleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

//Concrete builder Car

type CarBuilder struct {
	v VehickleProduct
}

func (b *CarBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 4
	return b
}

func (b *CarBuilder) SetSeats() BuildProcess {
	b.v.Seats = 5
	return b
}

func (b *CarBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Car"
	return b
}

func (b *CarBuilder) GetVehicle() VehickleProduct {
	return b.v
}

//Concrete builder Bike

type BikeBuilder struct {
	v VehickleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 1
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehickleProduct {
	return b.v
}
