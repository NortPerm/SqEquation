package battle

import (
	"errors"

	"github.com/NortPerm/SqEquation/solver"
)

var ErrInvalidVector = errors.New("invalid vector")

type Vector struct {
	x, y float64
}

func NewVector(x, y float64) *Vector {
	return &Vector{x: x, y: y}
}

func (v *Vector) Add(v1 *Vector) (*Vector, error) {
	if !v.correct() || !v1.correct() {
		return nil, ErrInvalidVector
	}
	return NewVector(v.x+v1.x, v.y+v1.y), nil
}

func (v *Vector) correct() bool {
	return solver.IsNumberList(v.x, v.y)
}
