package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"sort"
	"strconv"
)

var clear map[string]func()

type Coordinates struct {
	x string
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
	clearConsole()
	gameMap = make(map[Coordinates]string, 25)
	fillMap(&gameMap)
	var tries int = 5
	var hit bool
	drawEmptyMap(hit, tries, nil)
	for tries > 0 {
		hit, err := userInput(&gameMap)
		if !hit && err == nil {
			tries--
		}
		drawEmptyMap(hit, tries, err)
	}
	if tries <= 0 {
		drawMap(hit, tries, gameMap)
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
	x := []string{"A", "B", "C", "D", "E"}
	for i := range 5 {
		fmt.Println("---------------------")
		for _, j := range x {
			if gameMap[Coordinates{x: j, y: i}] != "b" {
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
	x := []string{"A", "B", "C", "D", "E"}
	fmt.Println("  A   B   C   D   E")

	for i := range 5 {
		fmt.Println("---------------------")
		for _, j := range x {
			fmt.Printf("| %s ", gameMap[Coordinates{x: j, y: i}])
		}
		fmt.Printf("| %d\n", i+1)
	}
	fmt.Println("---------------------")
}

// fill map randomly
func fillMap(gameMap *map[Coordinates]string) {
	// TODO: increase or decrease the number of empty slots to increase or decrese difficulty
	ch := []string{" ", "b"}
	x := []string{"A", "B", "C", "D", "E"}

	for i := range 5 {
		for _, j := range x {

			(*gameMap)[Coordinates{x: j, y: i}] = ch[rand.Intn(2)]
		}
	}
}

// accept user input and validate it
func userInput(gameMap *map[Coordinates]string) (bool, error) {
	xin := []string{"A", "B", "C", "D", "E"}

	var hit bool = false
	coordinates := Coordinates{}

	r := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the coordinates in this format (A 1): ")
	inp, _, err := r.ReadLine()
	if err != nil {
		return false, errors.New("invalid coordinates")
	}
	coordinates.x = string(inp[0:1])
	coordinates.y, err = strconv.Atoi(string(inp[2:]))
	if err != nil {
		return false, errors.New("invalid y coordinate")
	}

	if i := sort.SearchStrings(xin, coordinates.x); i > len(xin)-1 {
		return hit, errors.New("invalid x coordinate")
	}

	if coordinates.y > 5 {
		return hit, errors.New("invalid y coordinate")
	}

	coordinates.y = coordinates.y - 1

	if (*gameMap)[coordinates] == "b" {
		(*gameMap)[coordinates] = "*"
		hit = true
	}
	return hit, nil
}
