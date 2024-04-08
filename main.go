package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
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
	// drawMap()
	gameMap = make(map[Coordinates]string, 25)
	fillMap(&gameMap)
	var tries int = 5
	var hit bool = false
	for tries > 0 {
		hit = userInput(&gameMap, hit, tries)
		if !hit {
			tries--
		}
	}
}

func clearConsole() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	}
}

func drawEmptyMap(hit bool, tries int) {
	fmt.Printf("tries left: %d   hit: %t\n", tries, hit)
	fmt.Println(" A  B  C  D  E")
	fmt.Println("----------------")
	fmt.Println("|  |  |  |  |  | 1")
	fmt.Println("----------------")
	fmt.Println("|  |  |  |  |  | 2")
	fmt.Println("----------------")
	fmt.Println("|  |  |  |  |  | 3")
	fmt.Println("----------------")
	fmt.Println("|  |  |  |  |  | 4")
	fmt.Println("----------------")
	fmt.Println("|  |  |  |  |  | 5")
	fmt.Println("---------------------")
}

func drawMap(hit bool, tries int) {
	fmt.Printf("tries left: %d   hit: %t\n", tries, hit)
	x := []string{"A", "B", "C", "D", "E"}
	fmt.Println("  A   B   C   D   E")

	gameMap = make(map[Coordinates]string, 25)

	fillMap(&gameMap)

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

func userInput(gameMap *map[Coordinates]string, prevHit bool, tries int) bool {
	// clearConsole()
	drawMap(prevHit, tries)
	var hit bool = false
	var x string
	var y int

	fmt.Println("Enter the coordinates in this format (A 1): ")
	fmt.Scanf("%s %d", x, y)
	fmt.Println(x, y, (*gameMap)[Coordinates{x: x, y: y}])
	if (*gameMap)[Coordinates{x: x, y: y}] == "b" {
		hit = true
	}
	return hit
}
