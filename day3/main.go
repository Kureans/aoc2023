package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func getNumEnd(start int, bpline string) int {
	for i := start; i < len(bpline); i++ {
		if bpline[i] < '0' || bpline[i] > '9' {
			return i - 1
		}
	}
	return len(bpline) - 1
}

func getNumStart(start int, bpline string) int {
	for i := start; i >= 0; i-- {
		if bpline[i] < '0' || bpline[i] > '9' {
			return i + 1
		}
	}
	return 0
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAdjacent(start int, end int, rowNum int, blueprintLines []string, symbol byte) bool {
	height := len(blueprintLines)
	width := len(blueprintLines[0])
	borderColStart := max(start-1, 0)
	borderColEnd := min(end+1, width-1)
	borderRowStart := max(rowNum-1, 0)
	borderRowEnd := min(rowNum+1, height-1)

	// check top & bottom border
	for i := borderColStart; i <= borderColEnd; i++ {
		if rowNum != 0 && blueprintLines[borderRowStart][i] != symbol {
			return true
		}
		if rowNum != height-1 && blueprintLines[borderRowEnd][i] != symbol {
			return true
		}
	}

	// check left and right border
	for i := borderRowStart; i <= borderRowEnd; i++ {
		if start != 0 && blueprintLines[i][borderColStart] != symbol {
			return true
		}
		if end != width-1 && blueprintLines[i][borderColEnd] != symbol {
			return true
		}
	}

	return false
}

func getTwoAdjacentNumbers(colNum int, rowNum int, blueprintLines []string, idxToStringMp map[int]NumStringIdx) int {
	firstNumber := NumStringIdx{-1, -1, -1}
	secondNumber := NumStringIdx{-1, -1, -1}
	height := len(blueprintLines)
	width := len(blueprintLines[0])
	borderColStart := max(colNum-1, 0)
	borderColEnd := min(colNum+1, width-1)
	borderRowStart := max(rowNum-1, 0)
	borderRowEnd := min(rowNum+1, height-1)

	// check top & bottom border
	for i := borderColStart; i <= borderColEnd; i++ {
		if rowNum != 0 && isDigit(blueprintLines[borderRowStart][i]) {
			isValid := tryUpdateNumbers(&firstNumber, &secondNumber, borderRowStart, i, width, idxToStringMp)
			if !isValid {
				return 0
			}
		}

		if rowNum != height-1 && isDigit(blueprintLines[borderRowEnd][i]) {
			isValid := tryUpdateNumbers(&firstNumber, &secondNumber, borderRowEnd, i, width, idxToStringMp)
			if !isValid {
				return 0
			}
		}
	}

	// check left and right border
	for i := borderRowStart; i <= borderRowEnd; i++ {
		if colNum != 0 && isDigit(blueprintLines[i][borderColStart]) {
			isValid := tryUpdateNumbers(&firstNumber, &secondNumber, i, borderColStart, width, idxToStringMp)
			if !isValid {
				return 0
			}
		}
		if colNum != width-1 && isDigit(blueprintLines[i][borderColEnd]) {
			isValid := tryUpdateNumbers(&firstNumber, &secondNumber, i, borderColEnd, width, idxToStringMp)
			if !isValid {
				return 0
			}
		}
	}

	if firstNumber.isEmpty() || secondNumber.isEmpty() {
		return 0
	}
	firstNumberInt, err := firstNumber.getInt(blueprintLines)
	if err != nil {
		fmt.Println("Error converting NumStringIdx to Number")
	}
	secondNumberInt, err := secondNumber.getInt(blueprintLines)
	if err != nil {
		fmt.Println("Error converting NumStringIdx to Number")
	}

	return firstNumberInt * secondNumberInt
}

func tryUpdateNumbers(firstNumber *NumStringIdx, secondNumber *NumStringIdx, rowNum int, colNum int, width int, idxToNumStringMp map[int]NumStringIdx) bool {
	targetNumber := idxToNumStringMp[(rowNum*width)+colNum]
	if firstNumber.isEmpty() {

		*firstNumber = targetNumber
		// fmt.Println("First Number Empty, assigning ", targetNumber)
		// fmt.Println("First Number: ", *firstNumber)
	} else if secondNumber.isEmpty() && !firstNumber.isEqual(targetNumber) {
		*secondNumber = targetNumber
		// fmt.Println("Second Number Empty, assigning ", targetNumber)
		// fmt.Println("Second Number: ", *secondNumber)
	} else if !firstNumber.isEqual(targetNumber) && !secondNumber.isEqual(targetNumber) {
		// fmt.Println("too many numbers")
		// fmt.Println("First Number: ", *firstNumber)
		// fmt.Println("Second Number: ", *secondNumber)
		// fmt.Println("Third Number: ", targetNumber)
		return false
	}
	return true
}

func tryAddPartNum(rowNum int, start int, blueprintLines []string) (int, int) {
	end := getNumEnd(start, blueprintLines[rowNum])
	isAdjacent := isAdjacent(start, end, rowNum, blueprintLines, '.')
	if isAdjacent {
		numValue, err := strconv.Atoi(blueprintLines[rowNum][start : end+1])
		if err != nil {
			fmt.Println(err)
		}
		return numValue, end
	}
	return 0, end
}

type NumStringIdx struct {
	rowNum int
	start  int
	end    int
}

func (obj NumStringIdx) isEmpty() bool {
	return obj.rowNum == -1
}

func (a NumStringIdx) isEqual(b NumStringIdx) bool {
	return a.rowNum == b.rowNum && a.start == b.start && a.end == b.end
}

func (obj NumStringIdx) getInt(blueprintLines []string) (int, error) {
	return strconv.Atoi(blueprintLines[obj.rowNum][obj.start : obj.end+1])
}

func (obj NumStringIdx) String() string {
	return fmt.Sprintf("Start [%d], End [%d], rowNum [%d]\n", obj.start, obj.end, obj.rowNum)
}

func part2() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error")
	}

	sum := 0
	blueprintString := string(data)
	blueprintLines := strings.Split(blueprintString, "\n")
	height := len(blueprintLines)
	width := len(blueprintLines[0])
	idxToNumStringMp := make(map[int]NumStringIdx)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if blueprintLines[i][j] < '0' || blueprintLines[i][j] > '9' {
				continue
			}
			end := getNumEnd(j, blueprintLines[i])
			start := getNumStart(j, blueprintLines[i])
			idx := (i * width) + j
			idxToNumStringMp[idx] = NumStringIdx{i, start, end}
		}
	}
	// for k, v := range idxToNumStringMp {
	// 	fmt.Println("Key: ", k)
	// 	fmt.Println("Value: ", v)
	// }

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if blueprintLines[i][j] != '*' {
				continue
			}
			sum += getTwoAdjacentNumbers(j, i, blueprintLines, idxToNumStringMp)
		}
	}

	fmt.Println("Sum: ", sum)
}

func main() {
	part2()
}

func part1() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error")
	}

	sum := 0
	blueprintString := string(data)
	blueprintLines := strings.Split(blueprintString, "\n")
	height := len(blueprintLines)
	width := len(blueprintLines[0])
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if blueprintLines[i][j] < '0' || blueprintLines[i][j] > '9' {
				continue
			}
			value, nextIdx := tryAddPartNum(i, j, blueprintLines)
			sum += value
			j = nextIdx
		}
	}
	fmt.Println("Sum: ", sum)
}
