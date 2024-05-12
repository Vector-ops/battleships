package main

import (
	"math/rand"

	"github.com/Vector-ops/battleships/enums"
	"github.com/Vector-ops/battleships/types"
)

func getValidCoordinates(s types.ShipSize, gm types.GameMap) []types.ValidCoordinates {
	size := 0
	ship := []string{"s", "m", "l"}
	var valid []types.ValidCoordinates
	for i := range 5 {
		var coord types.Coordinates
		for j := range 5 {
			if gm[types.Coordinates{X: j, Y: i}] == " " {
				if size == 0 {
					coord = types.Coordinates{X: j, Y: i}
				}
				size++
				if size == int(s) {
					valid = append(valid, types.ValidCoordinates{Direction: enums.Horizontal, StCd: coord, EdCd: types.Coordinates{X: j, Y: i}, Ship: ship[int(s)-1]})
					size = 0
				}
			} else {
				size = 0
			}
		}
		size = 0
	}
	for i := range 5 {
		var coord types.Coordinates
		for j := range 5 {
			if gm[types.Coordinates{X: i, Y: j}] == " " {
				if size == 0 {
					coord = types.Coordinates{X: i, Y: j}
				}
				size++
				if size == int(s) {
					valid = append(valid, types.ValidCoordinates{Direction: enums.Vertical, StCd: coord, EdCd: types.Coordinates{X: i, Y: j}, Ship: ship[int(s)-1]})
					size = 0
				}
			} else {
				size = 0
			}
		}
		size = 0
	}
	return valid
}

func PlaceShips(shipSize types.ShipSize, count int, gm *types.GameMap) {
	validCoords := getValidCoordinates(shipSize, *gm)
	for range count {
		// pick a random coordinate
		ind := rand.Intn(len(validCoords))
		choice := validCoords[ind]

		// remove chosen coordinate from the list of available coordinates
		validCoords = append(validCoords, validCoords[:ind]...)
		validCoords = append(validCoords, validCoords[:ind+1]...)

		if choice.Direction == enums.Horizontal {
			for i := choice.StCd.X; i <= choice.EdCd.X; i++ {
				(*gm)[types.Coordinates{X: i, Y: choice.StCd.Y}] = choice.Ship
			}
		} else {
			for i := choice.StCd.Y; i <= choice.EdCd.Y; i++ {
				(*gm)[types.Coordinates{X: choice.StCd.X, Y: i}] = choice.Ship
			}
		}
	}
}
