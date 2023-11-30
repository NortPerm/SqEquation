package battle

type Spaceship struct {
	gameObject // композиция над игровым объектом
}

func NewSpaceship(x, y, v float64, d, dcount int) *Spaceship {
	spaceShip := &Spaceship{}
	spaceShip.gameObject = *NewGameobject()
	spaceShip.SetProperty(objDirection, d)
	spaceShip.SetProperty(objDirectionsNumber, dcount)
	spaceShip.SetProperty(objPosition, NewVector(x, y))
	spaceShip.SetProperty(objVelocity, v)
	return spaceShip
}
