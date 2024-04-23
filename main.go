package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
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
	for tries > 0 {
		drawEmptyMap(hit, tries)
		hit, e := userInput(&gameMap)
		if !hit && !e {
			tries--
		}
	}
	if tries == 0 || tries < 0 {
		drawMap(hit, tries, gameMap)
	}
}

func clearConsole() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	}
}

func drawEmptyMap(hit bool, tries int) {
	clearConsole()
	fmt.Printf("tries left: %d   hit: %t\n", tries, hit)
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

func userInput(gameMap *map[Coordinates]string) (bool, bool) {

	var hit bool = false
	var x string
	var y int

	r := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the coordinates in this format (A 1): ")
	inp, _, err := r.ReadLine()
	if err != nil {
		fmt.Printf("Invalid coordinates.\n")
		return false, true
	}
	x = string(inp[0:1])
	y, err = strconv.Atoi(string(inp[2:]))
	if err != nil {
		fmt.Printf("Invalid x coordinate")
		return false, true
	}

	fmt.Println(x, y, (*gameMap)[Coordinates{x: x, y: y}])
	if (*gameMap)[Coordinates{x: x, y: (y - 1)}] == "b" {
		(*gameMap)[Coordinates{x: x, y: (y - 1)}] = "*"
		hit = true
	}
	return hit, false
}
