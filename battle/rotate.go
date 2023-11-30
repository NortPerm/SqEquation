package battle

type Rotable interface {
	GetDirection() (int, error)
	GetAngularVelocity() (int, error)
	GetDirectionsNumber() (int, error)
	SetDirection(direction int) error
}

type RotateCommand struct {
	rotable Rotable
}

func NewRotateCommand(r Rotable) *RotateCommand {
	return &RotateCommand{rotable: r}
}

func (rc *RotateCommand) Rotate(r Rotable) error {
	rc.rotable = r
	return nil
}

func (rc *RotateCommand) Execute() error {
	direction, err := rc.rotable.GetDirection()
	if err != nil {
		return err
	}
	directionsNumber, err := rc.rotable.GetDirectionsNumber()
	if err != nil {
		return err
	}
	velocity, err := rc.rotable.GetAngularVelocity()
	if err != nil {
		return err
	}

	return rc.rotable.SetDirection((direction + velocity) % directionsNumber)
}
