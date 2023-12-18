package main

import (
	"fmt"
	"os"
)

func main() {
	lines := getLines("input.txt")
	fmt.Println(lines)
}

func getLines(filename string) string {
	bytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error", err)
	}
	contents := string(bytes)

	return contents
}
