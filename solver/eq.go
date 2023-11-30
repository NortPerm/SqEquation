package solver

import (
	"errors"
	"math"
)

var (
	errInvalidFirstCoeff = errors.New("a must not be zero")
	errNotNumberCoeff    = errors.New("at least one coefficient is not a number")
)

const eps float64 = 1e-10

type squareEquation struct {
	a, b, c float64
}

func NewSquareEquation(a, b, c float64) *squareEquation {
	// строго говоря конструктор должен валидировать коэффициенты и проверка А происходит тут, а не в момент рещения
	// исходя из требований ДЗ проверка перенесена в solve
	// по этой причине подразумеваем, что конструктор не может выбросить ошибку
	return &squareEquation{a: a, b: b, c: c}
}

func (se *squareEquation) solve() ([]float64, error) {
	if !isNumberList(se.a, se.b, se.c) {
		return nil, errNotNumberCoeff
	}
	if isFloatEqual(se.a, 0) {
		return nil, errInvalidFirstCoeff
	}
	d := se.b*se.b - 4*se.a*se.c
	if d < 0 {
		return []float64{}, nil
	}
	if isFloatEqual(d, 0) {
		return []float64{-se.b / (2 * se.a)}, nil
	}
	return []float64{(-se.b + math.Sqrt(d)) / (2 * se.a), (-se.b - math.Sqrt(d)) / (2 * se.a)}, nil
}

func isFloatEqual(a, b float64) bool {
	return math.Abs(a-b) < eps
}

func isNumber(a float64) bool {
	return !math.IsInf(a, 1) && !math.IsInf(a, -1) && !math.IsNaN(a)
}

func isNumberList(list ...float64) bool {
	for _, v := range list {
		if !isNumber(v) {
			return false
		}
	}
	return true
}

func IsNumberList(list ...float64) bool {
	return isNumberList(list...)
}
