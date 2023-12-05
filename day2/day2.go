package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type GameInfo struct {
	id    int
	red   []int
	blue  []int
	green []int
}

func errorHandler(e error) {
	if e != nil {
		fmt.Println("Error Encountered")
		panic(e)
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func parseGameInput(input string) int {
	arr := strings.FieldsFunc(input, func(r rune) bool {
		return r == ':'
	})

	// gameId, err := strconv.Atoi(strings.Fields(arr[0])[1])
	// errorHandler(err)

	redMinCount, blueMinCount, greenMinCount := 0, 0, 0

	sets := strings.FieldsFunc(arr[1], func(r rune) bool {
		return r == ';'
	})

	for _, set := range sets {
		colors := strings.FieldsFunc(set, func(r rune) bool {
			return r == ','
		})

		for _, colorString := range colors {
			colorCountPair := strings.Fields(colorString)
			colorCountValue, err := strconv.Atoi(colorCountPair[0])
			errorHandler(err)
			if colorCountPair[1] == "blue" {
				blueMinCount = max(blueMinCount, colorCountValue)
			}
			if colorCountPair[1] == "red" {
				redMinCount = max(redMinCount, colorCountValue)
			}
			if colorCountPair[1] == "green" {
				greenMinCount = max(greenMinCount, colorCountValue)
			}
		}
	}

	return redMinCount * blueMinCount * greenMinCount
}

func main() {
	f, err := os.Open("./input.txt")
	errorHandler(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	validGameCount := 0
	for scanner.Scan() {
		validGameCount += parseGameInput(scanner.Text())
	}
	fmt.Println("Valid Game Count: ", validGameCount)
}
