package builder

import "testing"

func TestBuilderPattern(t *testing.T) {
	manufacturingComplex := ManufacturingDirector{}

	carBuilder := &CarBuilder{}
	manufacturingComplex.SetBuilder(carBuilder)
	manufacturingComplex.Construct()

	car := carBuilder.GetVehicle()

	if car.Wheels != 4 {
		t.Errorf("Wheels on a car must be 4")
	}

	if car.Seats != 5 {
		t.Errorf("Seats on a car must be 5")
	}

	if car.Structure != "Car" {
		t.Errorf("Structure on a car must be 'Car'")
	}

	bikeBuilder := &BikeBuilder{}
	manufacturingComplex.SetBuilder(bikeBuilder)
	manufacturingComplex.Construct()

	bike := bikeBuilder.GetVehicle()

	if bike.Wheels != 2 {
		t.Errorf("Wheels on a bike must be 2")
	}

	if bike.Seats != 1 {
		t.Errorf("Seats on a bike must be 1")
	}

	if bike.Structure != "Bike" {
		t.Errorf("Structure on a bike must be 'Bike'")
	}
}
