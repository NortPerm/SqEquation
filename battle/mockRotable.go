package battle

const (
	canReadDirection = 1 << iota
	canReadAngularVelocity
	canWriteDirection
)

type mockRotable struct { // implements movable for testing
	direction, directionsNumber, velocity int
	canReadDirection                      bool
	canReadAngularVelocity                bool
	canWriteDirection                     bool
}

func NewMockRotable(direction, directionsNumber, velocity, grants int) *mockRotable {
	return &mockRotable{
		direction:              direction,
		directionsNumber:       directionsNumber,
		velocity:               velocity,
		canReadDirection:       grants&canReadDirection > 0,
		canReadAngularVelocity: grants&canReadAngularVelocity > 0,
		canWriteDirection:      grants&canWriteDirection > 0,
	}
}

func (r *mockRotable) GetDirection() (int, error) {
	if r.canReadDirection {
		return r.direction, nil
	}
	return 0, ErrInvalidDirection
}

func (r *mockRotable) GetDirectionsNumber() (int, error) {
	if r.directionsNumber > 0 {
		return r.directionsNumber, nil
	}
	return 0, ErrInvalidDirectionsCount
}

func (r *mockRotable) GetAngularVelocity() (int, error) {
	if r.canReadAngularVelocity {
		return r.velocity, nil
	}
	return 0, ErrInvalidAngularVelocity
}

func (r *mockRotable) SetDirection(direction int) error {
	if r.canWriteDirection {
		r.direction = direction
		return nil
	}
	return ErrInvalidWriteDirection
}
