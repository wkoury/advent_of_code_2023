package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/hashicorp/go-set/v2"
)

func main() {
	lines := getLines("input.txt")

	partOneResult := partOne(lines)
	fmt.Println(partOneResult)

	partTwoResult := partTwo(lines)
	fmt.Println(partTwoResult)
}

func partOne(lines []string) int {
	ret := 0

	for _, line := range lines {
		cardValue := getCardValue(getWinningNumbers(line), getLotteryNumbers(line))
		ret += cardValue
	}

	return ret
}

func partTwo(lines []string) int {
	cards := []int{}

	for ii := 0; ii < len(lines); ii++ {
		cards = append(cards, 1)
	}

	for ii, line := range lines {
		for kk := 0; kk < cards[ii]; kk++ {
			// fmt.Printf("Processing line %d, card %d\n", ii, ii+1)
			// fmt.Printf("Cards: %d\n", cards[ii])
			cardValue := getCardValuePartTwo(getWinningNumbers(line), getLotteryNumbers(line))
			// fmt.Printf("Card value: %d\n", cardValue)
			for jj := 1; jj <= cardValue; jj++ {
				cards[ii+jj]++
			}
		}
	}

	sum := 0
	for _, card := range cards {
		sum += card
	}

	return sum
}

func getCardValuePartTwo(winners *set.Set[int], numbers *set.Set[int]) int {
	ret := 0

	numbers.ForEach(func(item int) bool {
		if winners.Contains(item) {
			ret++
		}
		return true
	})

	return ret
}

func getCardValue(winners *set.Set[int], numbers *set.Set[int]) int {
	ret := 0.5
	numbers.ForEach(func(item int) bool {
		if winners.Contains(item) {
			ret *= 2
		}
		return true
	})

	if ret == 0.5 {
		return 0
	}

	return int(ret)
}

func getWinningNumbers(line string) *set.Set[int] {
	ret := set.New[int](10)

	firstHalf := strings.Split(line, "|")[0]
	postComma := strings.Split(firstHalf, ":")[1]
	words := strings.Split(postComma, " ")

	for _, word := range words {
		num, err := strconv.Atoi(word)
		if err == nil {
			ret.Insert(num)
		}
	}

	return ret
}

func getLotteryNumbers(line string) *set.Set[int] {
	ret := set.New[int](25)

	secondHalf := strings.Split(line, "|")[1]
	words := strings.Split(secondHalf, " ")

	for _, word := range words {
		num, err := strconv.Atoi(word)
		if err == nil {
			ret.Insert(num)
		}
	}

	return ret
}

func getLines(filename string) []string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("error", err)
	}
	contents := string(bytes)

	return strings.Split(contents, "\n")
}
