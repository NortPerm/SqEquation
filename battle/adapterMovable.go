// этот файл задел на будущее, в данной работе адаптер не нужен и работа с реальными объектами тоже

package battle

import "math"

type MovableAdapter struct { // implements movable for real object
	obj GameObject
}

func NewMovableAdapter(obj GameObject) *MovableAdapter {
	return &MovableAdapter{obj: obj}
}

func (ma *MovableAdapter) GetPosition() (*Vector, error) {
	uptypedPosition, err := ma.obj.GetProperty(objPosition)
	if err != nil {
		return nil, ErrInvalidPosition
	}
	position, ok := uptypedPosition.(*Vector)
	if !ok {
		return nil, ErrInvalidTypecast
	}
	return position, nil
}

func (ma *MovableAdapter) GetVelocity() (*Vector, error) {
	// да из-за наличия трай-кетчей приходится проверять каждую операцию, что делает код не столь компактным как на шарпе
	untypedVelocity, err := ma.obj.GetProperty(objVelocity)
	if err != nil {
		return nil, ErrInvalidVelocity
	}
	velocity, ok := untypedVelocity.(int)
	if !ok {
		return nil, ErrInvalidTypecast
	}

	untypedDir, err := ma.obj.GetProperty(objDirection)
	if err != nil {
		return nil, ErrInvalidVelocity
	}
	direction, ok := untypedDir.(int)
	if !ok {
		return nil, ErrInvalidTypecast
	}

	untypedDirNums, err := ma.obj.GetProperty(objDirectionsNumber)
	if err != nil {
		return nil, ErrInvalidVelocity
	}
	dirNums, ok := untypedDirNums.(int)
	if !ok {
		return nil, ErrInvalidTypecast
	}

	angle := 2 * math.Pi * float64(direction) / float64(dirNums)
	return NewVector(
		float64(velocity)*math.Cos(angle),
		float64(velocity)*math.Sin(angle),
	), nil
}

func (ma *MovableAdapter) SetPosition(position *Vector) error {
	if err := ma.obj.SetProperty(objPosition, position); err != nil {
		return ErrInvalidWritePosition
	}
	return nil
}
