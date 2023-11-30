package battle

import "errors"

// interface
type (
	Command interface {
		Execute() error // го не умеет в исключения, потому любой вызов команды должен вернуть наружу ошибку выполнения
	}
)

var (
	ErrInvalidPosition      = errors.New("can not get position of object")
	ErrInvalidVelocity      = errors.New("can not get velocity of object")
	ErrInvalidWritePosition = errors.New("can not set position of object")
	ErrInvalidTypecast      = errors.New("can not typecast interface to object")

	ErrInvalidDirection       = errors.New("can not get position of object")
	ErrInvalidAngularVelocity = errors.New("can not get angular velocity of object")
	ErrInvalidWriteDirection  = errors.New("can not set position of object")
	ErrInvalidDirectionsCount = errors.New("directions count must be above the zero")
)
