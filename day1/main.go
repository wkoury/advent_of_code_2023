package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error", err)
	}
	contentString := string(content)
	contentString = strings.ReplaceAll(contentString, "one", "one1one")
	contentString = strings.ReplaceAll(contentString, "two", "two2two")
	contentString = strings.ReplaceAll(contentString, "three", "three3three")
	contentString = strings.ReplaceAll(contentString, "four", "four4four")
	contentString = strings.ReplaceAll(contentString, "five", "five5five")
	contentString = strings.ReplaceAll(contentString, "six", "six6six")
	contentString = strings.ReplaceAll(contentString, "seven", "seven7seven")
	contentString = strings.ReplaceAll(contentString, "eight", "eight8eight")
	contentString = strings.ReplaceAll(contentString, "nine", "nine9nine")

	lines := strings.Split(contentString, "\n")

	sum := 0

	for ii, line := range lines {
		if ii < len(lines)-1 {

			// Sentinel values
			firstNumber := -1
			lastNumber := -1

			re := regexp.MustCompile("[0-9]")

			// Find the first number
			for ii := range line {
				if firstNumber == -1 && re.Match([]byte{line[ii]}) {
					firstNumber, _ = strconv.Atoi(string(line[ii]))
				}
			}

			// Find the last number
			for ii := len(line) - 1; ii >= 0; ii-- {
				if lastNumber == -1 && re.Match([]byte{line[ii]}) {
					lastNumber, _ = strconv.Atoi(string(line[ii]))
				}
			}

			sum += firstNumber*10 + lastNumber
		}
	}

	fmt.Println(sum)
}
