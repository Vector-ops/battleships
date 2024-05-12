package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"

	"github.com/Vector-ops/battleships/enums"
	"github.com/Vector-ops/battleships/types"
)

var clear map[string]func()

const (
	// to render ships on map during gameplay
	debugDraw bool = false

	// to stop gameplay and check the ship placements on map
	debugMap bool = false

	targetHit  string = "*"
	targetMiss string = "o"
)

var gameMap types.GameMap

func init() {
	log.SetFlags(0)

	clear = make(map[string]func())

	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func main() {
	var diff int
	flag.IntVar(&diff, "d", 0, "Set the difficulty of the game")
	flag.Parse()
	clearConsole()

	gameMap = make(types.GameMap, 25)
	var saveFile types.SaveFile
	saveFile.StartMap = make(types.GameMap, 25)
	saveFile.EndMap = make(types.GameMap, 25)

	fillMap(&gameMap)
	saveFile.Time = time.Now().Format(time.RFC3339)
	CopyMap(gameMap, &saveFile.StartMap)
	var tries int = 5
	var hit bool

	if debugDraw {
		drawMap(hit, tries, gameMap)
	} else {
		drawEmptyMap(hit, tries, nil)
	}

	if !debugMap {
		for tries > 0 && !checkWin(gameMap) {
			hit, err := userInput(&gameMap)
			if !hit && err == nil {
				tries--
			}
			if debugDraw {
				drawMap(hit, tries, gameMap)
			} else {
				drawEmptyMap(hit, tries, err)
			}
		}
		drawMap(hit, tries, gameMap)
		if tries <= 0 {
			fmt.Printf("You lost!\n")
		} else if checkWin(gameMap) {
			fmt.Printf("You Won!\n")
		}
		CopyMap(gameMap, &saveFile.EndMap)
		saveFile.TriesLeft = tries
		saveFile.Win = checkWin(gameMap)
	}
	SaveMap(saveFile)
}

func clearConsole() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	}
}

// draw empty map
func drawEmptyMap(hit bool, tries int, err error) {
	clearConsole()
	if err != nil {
		fmt.Printf("tries left: %d   hit: %s\n", tries, err)
	} else {
		fmt.Printf("tries left: %d   hit: %t\n", tries, hit)
	}
	fmt.Println("  A   B   C   D   E")
	for i := range 5 {
		fmt.Println("---------------------")
		for j := range 5 {
			if gameMap[types.Coordinates{X: j, Y: i}] == targetHit || gameMap[types.Coordinates{X: j, Y: i}] == targetMiss {
				fmt.Printf("| %s ", gameMap[types.Coordinates{X: j, Y: i}])
			} else {
				fmt.Printf("|   ")
			}
		}
		fmt.Printf("| %d\n", i+1)
	}
	fmt.Println("---------------------")
}

// draw map after game ends
func drawMap(hit bool, tries int, gameMap types.GameMap) {
	clearConsole()
	fmt.Printf("tries left: %d   hit: %t\n", tries, hit)
	fmt.Println("  A   B   C   D   E")

	for i := range 5 {
		fmt.Println("---------------------")
		for j := range 5 {
			fmt.Printf("| %s ", gameMap[types.Coordinates{X: j, Y: i}])
		}
		fmt.Printf("| %d\n", i+1)
	}
	fmt.Println("---------------------")
}

// fill map randomly
func fillMap(gameMap *types.GameMap) {

	// initialize map with " "
	for i := range 5 {
		for j := range 5 {
			(*gameMap)[types.Coordinates{X: i, Y: j}] = " "
		}
	}

	PlaceShips(enums.Large, 2, gameMap)
	PlaceShips(enums.Medium, 3, gameMap)
	PlaceShips(enums.Small, 4, gameMap)
}

// accept user input and validate it
func userInput(gameMap *types.GameMap) (bool, error) {

	var hit bool = false
	coordinates := types.Coordinates{}

	r := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the coordinates in this format (A 1): ")
	inp, _, err := r.ReadLine()

	if err != nil {
		return false, errors.New("invalid coordinates")
	}
	coordinates.X = int(rune(inp[0]) - 'A')
	coordinates.Y, err = strconv.Atoi(string(inp[2:]))
	if err != nil {
		return false, errors.New("invalid y coordinate")
	}

	if coordinates.X > 4 || coordinates.X < 0 {
		return hit, errors.New("invalid x coordinate")
	}

	if coordinates.Y > 5 {
		return hit, errors.New("invalid y coordinate")
	}

	coordinates.Y = coordinates.Y - 1

	if (*gameMap)[coordinates] != " " && (*gameMap)[coordinates] != targetHit && (*gameMap)[coordinates] != targetMiss {
		(*gameMap)[coordinates] = targetHit
		hit = true
	} else if (*gameMap)[coordinates] == " " {
		(*gameMap)[coordinates] = targetMiss
		hit = false
	} else if (*gameMap)[coordinates] == targetHit || (*gameMap)[coordinates] == targetMiss {
		return hit, errors.New("coordinates already hit")
	}
	return hit, nil
}

func checkWin(gameMap types.GameMap) bool {
	for i := range 5 {
		for j := range 5 {
			if gameMap[types.Coordinates{X: i, Y: j}] != " " && gameMap[types.Coordinates{X: i, Y: j}] != targetHit {
				return false
			}
		}
	}
	return true
}
