package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func errorHandler(e error) {
	if e != nil {
		fmt.Println("Error Encountered")
		panic(e)
	}
}

func parseCardInput(cardInput string) int {
	gamePoints := 0
	inputArr := strings.Split(cardInput, "|")
	cardValuesString := strings.Split(inputArr[0], ":")[1]
	cardValues := strings.Fields(cardValuesString)
	playerValues := strings.Fields(inputArr[1])
	cardValueMap := make(map[int]bool)

	for _, vString := range cardValues {
		v, err := strconv.Atoi(vString)
		if err != nil {
			errorHandler(err)
		}
		cardValueMap[v] = true
	}
	for _, vString := range playerValues {
		v, err := strconv.Atoi(vString)
		if err != nil {
			errorHandler(err)
		}
		if cardValueMap[v] == true {
			if gamePoints == 0 {
				gamePoints += 1
			} else {
				gamePoints *= 2
			}
		}
	}

	return gamePoints
}

func addNewCardValues(idx int, cardInput string, cardsCountArr []int) int {
	inputArr := strings.Split(cardInput, "|")
	cardValuesString := strings.Split(inputArr[0], ":")[1]
	cardValues := strings.Fields(cardValuesString)
	playerValues := strings.Fields(inputArr[1])
	winCounts := 0
	currCopies := cardsCountArr[idx]
	cardValueMap := make(map[int]bool)

	for _, vString := range cardValues {
		v, err := strconv.Atoi(vString)
		if err != nil {
			errorHandler(err)
		}
		cardValueMap[v] = true
	}
	for _, vString := range playerValues {
		v, err := strconv.Atoi(vString)
		if err != nil {
			errorHandler(err)
		}
		if cardValueMap[v] == true {
			winCounts++
		}
	}

	for i := idx + 1; i < idx+1+winCounts; i++ {
		cardsCountArr[i] += currCopies
	}

	return cardsCountArr[idx]
}

func part1() {
	f, err := os.Open("./input.txt")
	errorHandler(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	points := 0
	for scanner.Scan() {
		points += parseCardInput(scanner.Text())
	}
	fmt.Println("Total Points: ", points)
}

func part2() {
	contentBytes, err := os.ReadFile("./input.txt")
	errorHandler(err)
	cardsArr := strings.Split(string(contentBytes), "\n")
	cardCountsArr := make([]int, len(cardsArr)+1)
	for i := 0; i < len(cardCountsArr); i++ {
		cardCountsArr[i] = 1
	}
	points := 0
	for i, v := range cardsArr {
		points += addNewCardValues(i+1, v, cardCountsArr)
	}
	// for i, v := range cardCountsArr {
	// 	fmt.Println(i, v)
	// }
	fmt.Println("Total Points: ", points)
}

func main() {
	part2()
}
