// этот файл задел на будущее - абстрактный объект, который может быть передан в нужные адаптeры

package battle

import (
	"errors"
	"sync"
)

type GameObject interface {
	GetProperty(key objectState) (any, error) // ключи у нас enum, определенные ниже через йоту
	SetProperty(key objectState, value any) error
}

type objectState int

const (
	objUnknown objectState = iota
	objPosition
	objVelocity
	objAngularVelocity
	objDirection
	objDirectionsNumber
)

type gameObject struct { // реализует GameObject через мап-сторадж
	mu    sync.RWMutex
	state map[objectState]any
}

func NewGameobject() *gameObject {
	gameObject := &gameObject{
		state: make(map[objectState]any),
	}
	return gameObject
}

func (g *gameObject) GetProperty(key objectState) (any, error) {
	g.mu.RLock()
	defer g.mu.RUnlock()
	property, ok := g.state[key]
	if !ok {
		return nil, errors.New("property not defined")
	}
	return property, nil
}

func (g *gameObject) SetProperty(key objectState, value any) error {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.state[key] = value
	return nil
}
