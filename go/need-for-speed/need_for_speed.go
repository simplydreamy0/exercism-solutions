package speed

// TODO: define the 'Car' type struct
type Car struct {
	battery, batteryDrain, speed, distance int
}

// NewCar creates a new remote controlled car with full battery and given specifications.
func NewCar(speed, batteryDrain int) Car {
	return Car{
		battery: 100,
		speed: speed,
		batteryDrain: batteryDrain,
		distance: 0,
	};
}

type Track struct {
	distance int
}
// NewTrack creates a new track
func NewTrack(distance int) Track {
	return Track{ distance: distance };
}

// Drive drives the car one time. If there is not enough battery to drive one more time,
// the car will not move.
func Drive(car Car) Car {
	if car.battery - car.batteryDrain < 0 {
		return car
	}
	return Car{
		battery: car.battery - car.batteryDrain,
		speed: car.speed,
		batteryDrain: car.batteryDrain,
		distance: car.distance + car.speed,
	}
}

// CanFinish checks if a car is able to finish a certain track.
func CanFinish(car Car, track Track) bool {
	numberOfDrives := car.battery / car.batteryDrain;
	maxDistance := numberOfDrives * car.speed;
	return maxDistance >= track.distance;
}
