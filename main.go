package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

var clear map[string]func()

const (
	debugFill bool = false
	debugDraw bool = true
)

type Coordinates struct {
	x int
	y int
}

var gameMap map[Coordinates]string

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

	gameMap = make(map[Coordinates]string, 25)
	fillMap(&gameMap)
	var tries int = 5
	var hit bool

	if debugDraw {
		drawMap(hit, tries, gameMap)
	} else {
		drawEmptyMap(hit, tries, nil)
	}
	// fmt.Println(gameMap)
	// valid := GetValidCoordinates(Large, gameMap)
	// fmt.Println(valid)

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
	if tries <= 0 {
		if debugDraw {
			drawMap(hit, tries, gameMap)
			fmt.Printf("You lost!\n")
		} else {
			drawEmptyMap(hit, tries, nil)
			fmt.Printf("You lost!\n")
		}
	} else if checkWin(gameMap) {
		if debugDraw {
			drawMap(hit, tries, gameMap)
			fmt.Printf("You Won!\n")
		} else {
			drawEmptyMap(hit, tries, nil)
			fmt.Printf("You Won!\n")
		}

	}
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
			if gameMap[Coordinates{x: i, y: j}] != "b" {
				fmt.Printf("| %s ", gameMap[Coordinates{x: j, y: i}])
			} else {
				fmt.Printf("|   ")
			}
		}
		fmt.Printf("| %d\n", i+1)
	}
	fmt.Println("---------------------")
}

// draw map after game ends
func drawMap(hit bool, tries int, gameMap map[Coordinates]string) {
	clearConsole()
	fmt.Printf("tries left: %d   hit: %t\n", tries, hit)
	fmt.Println("  A   B   C   D   E")

	for i := range 5 {
		fmt.Println("---------------------")
		for j := range 5 {
			fmt.Printf("| %s ", gameMap[Coordinates{x: j, y: i}])
		}
		fmt.Printf("| %d\n", i+1)
	}
	fmt.Println("---------------------")
}

// fill map randomly
func fillMap(gameMap *map[Coordinates]string) {
	// TODO: increase or decrease the number of empty slots to increase or decrese difficulty
	l := 2
	ch := []string{" ", "b"}

	if debugFill {
		ch = []string{" ", "l", "s", "m"}
		l = 4
	}
	d := 0
	count := 0

	for i := range 5 {
		for j := range 5 {
			if count < (5 - d) {
				(*gameMap)[Coordinates{x: i, y: j}] = ch[rand.Intn(l)]
				count++
			}
		}
		count = 0
	}
}

// accept user input and validate it
func userInput(gameMap *map[Coordinates]string) (bool, error) {

	var hit bool = false
	coordinates := Coordinates{}

	r := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the coordinates in this format (A 1): ")
	inp, _, err := r.ReadLine()

	if err != nil {
		return false, errors.New("invalid coordinates")
	}
	coordinates.x = int(rune(inp[0]) - 'A')
	coordinates.y, err = strconv.Atoi(string(inp[2:]))
	if err != nil {
		return false, errors.New("invalid y coordinate")
	}

	if coordinates.x > 4 || coordinates.x < 0 {
		return hit, errors.New("invalid x coordinate")
	}

	if coordinates.y > 5 {
		return hit, errors.New("invalid y coordinate")
	}

	coordinates.y = coordinates.y - 1

	//TODO: check for when the coord contains "*"
	//TODO: check for win condition

	if (*gameMap)[coordinates] == "b" {
		(*gameMap)[coordinates] = "*"
		hit = true
	}
	return hit, nil
}

func checkWin(gameMap map[Coordinates]string) bool {
	for i := range 5 {
		for j := range 5 {
			if gameMap[Coordinates{x: i, y: j}] == "b" {
				return false
			}
		}
	}
	return true
}
