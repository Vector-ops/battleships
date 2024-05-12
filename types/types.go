package types

type Coordinates struct {
	X int `json:"x"`
	Y int `json:"y"`
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

type SaveFile struct {
	Time      string  `json:"time"`
	StartMap  GameMap `json:"start_map"`
	EndMap    GameMap `json:"end_map"`
	Win       bool    `json:"win"`
	TriesLeft int     `json:"tries_left"`
}
