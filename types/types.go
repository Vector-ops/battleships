package types

type Coordinates struct {
	X int
	Y int
}

type GameMap map[Coordinates]string

type ShipSize int
type Direction string

type ValidCoordinates struct {
	Direction Direction
	StCd      Coordinates
	EdCd      Coordinates
	Ship      string
}
