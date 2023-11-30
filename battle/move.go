package battle

type Movable interface {
	GetPosition() (*Vector, error)
	GetVelocity() (*Vector, error)
	SetPosition(position *Vector) error
}

type MoveCommand struct {
	movable Movable
}

func NewMoveCommand(m Movable) *MoveCommand {
	return &MoveCommand{movable: m}
}

func (mc *MoveCommand) Move(m Movable) error {
	mc.movable = m
	return nil
}

func (mc *MoveCommand) Execute() error {
	position, err := mc.movable.GetPosition()
	if err != nil {
		return err
	}
	velocity, err := mc.movable.GetVelocity()
	if err != nil {
		return err
	}
	newPosition, err := position.Add(velocity)
	if err != nil {
		return err
	}
	return mc.movable.SetPosition(newPosition)
}
