package battle

const (
	canReadPosition = 1 << iota
	canReadVelocity
	canWritePosition
)

type mockMovable struct { // implements movable for testing
	position, velocity *Vector
	canReadPosition    bool
	canReadVelocity    bool
	canWritePosition   bool
}

func NewMockMovable(x, y, dx, dy float64, grants int) *mockMovable {
	return &mockMovable{
		position:         NewVector(x, y),
		velocity:         NewVector(dx, dy),
		canReadPosition:  grants&canReadPosition > 0,
		canReadVelocity:  grants&canReadVelocity > 0,
		canWritePosition: grants&canWritePosition > 0,
	}
}

func (m *mockMovable) GetPosition() (*Vector, error) {
	if m.position.correct() && m.canReadPosition {
		return m.position, nil
	}
	return nil, ErrInvalidPosition
}

func (m *mockMovable) GetVelocity() (*Vector, error) {
	if m.velocity.correct() && m.canReadVelocity {
		return m.velocity, nil
	}
	return nil, ErrInvalidVelocity
}

func (m *mockMovable) SetPosition(position *Vector) error {
	if position.correct() && m.canWritePosition {
		m.position = position
		return nil
	}
	return ErrInvalidWritePosition
}
