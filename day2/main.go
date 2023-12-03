package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	Red   string = "red"
	Green string = "green"
	Blue  string = "blue"
)

type Game struct {
	id, redCount, blueCount, greenCount int
}

func parseGame(str string) Game {

	redCount := 0
	blueCount := 0
	greenCount := 0

	idStr := strings.ReplaceAll((strings.Split(str, " ")[1]), ":", "")
	id, _ := strconv.Atoi(idStr)

	gameStr := strings.TrimSpace(strings.Split(str, ":")[1])
	gameSets := strings.Split(gameStr, ";")

	for _, set := range gameSets {
		cubes := strings.Split(set, ",")

		for _, cube := range cubes {
			words := strings.Split(strings.TrimSpace(cube), " ")
			quantity, err := strconv.Atoi(words[0])
			if err != nil {
				panic(err)
			}

			color := words[1]

			if color == Red && quantity > redCount {
				redCount = quantity
			}
			if color == Blue && quantity > blueCount {
				blueCount = quantity
			}
			if color == Green && quantity > greenCount {
				greenCount = quantity
			}
		}
	}

	ret := Game{
		id:         id,
		redCount:   redCount,
		blueCount:  blueCount,
		greenCount: greenCount,
	}

	return ret
}

func isGamePossible(game Game) bool {
	return game.redCount > 12 || game.greenCount > 13 || game.blueCount > 14
}

func main() {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error", err)
	}
	contents := string(bytes)

	sum := 0

	lines := strings.Split(contents, "\n")
	for _, line := range lines {
		game := parseGame(line)

		if !isGamePossible(game) {
			sum += game.id
		}
	}

	fmt.Println(sum)
}
