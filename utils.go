package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"path/filepath"

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

func CopyMap(src types.GameMap, dest *types.GameMap) {
	for key, value := range src {
		(*dest)[key] = value
	}
}

func SaveMap(saveFile types.SaveFile) {

	dirName := "saves"
	filePath := filepath.Join(dirName, "saves.txt")
	dirPath := filepath.Join(".", dirName)
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		log.Println(err)
	}

	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("Could not save file: %s\nerror: %T", filePath, err)
	}
	defer f.Close()

	var save string
	save += fmt.Sprintln("----------------------------------", saveFile.Time, "----------------------------------")
	save += "------------------------------Start Map-----------------------------------------\n"
	save += generateMapString(saveFile.StartMap)
	save += "------------------------------End Map-------------------------------------------\n"
	save += generateMapString(saveFile.EndMap)
	save += fmt.Sprintln("Tries left: ", saveFile.TriesLeft, "   ", "Win: ", saveFile.Win)

	_, err = f.WriteString(save)
	if err != nil {
		log.Println(err)
	}
}

func generateMapString(gameMap types.GameMap) string {
	var gmap string
	gmap += fmt.Sprintln("  A   B   C   D   E")

	for i := range 5 {
		gmap += fmt.Sprintln("---------------------")
		for j := range 5 {
			gmap += fmt.Sprintf("| %s ", gameMap[types.Coordinates{X: j, Y: i}])
		}
		gmap += fmt.Sprintf("| %d\n", i+1)
	}
	gmap += fmt.Sprintln("---------------------")

	return gmap
}
