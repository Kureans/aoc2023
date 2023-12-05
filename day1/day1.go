package main

import (
	"bufio"
	"fmt"
	"os"
)

func errorHandler(e error) {
	if e != nil {
		fmt.Println("Error Encountered")
		panic(e)
	}
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isStartChar(c byte) bool {
	return c == 'o' || c == 't' || c == 'f' || c == 's' || c == 'e' || c == 'n'
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isSpelledOutDigit(s string, start int, mp map[string]byte) byte {
	spelledOut := s[start:min(start+3, len(s))]
	if mp[spelledOut] != 0 {
		return mp[spelledOut]
	}
	spelledOut = s[start:min(start+4, len(s))]
	if mp[spelledOut] != 0 {
		return mp[spelledOut]
	}
	spelledOut = s[start:min(start+5, len(s))]
	if mp[spelledOut] != 0 {
		return mp[spelledOut]
	}
	return '0'
}

//Assumes that at least 2 digits exist per string passed in
func extractCalibrationValue(s string, mp map[string]byte) int {
	var leftDigit byte
	var rightDigit byte

	for left := 0; left < len(s); left++ {
		if isDigit(s[left]) {
			leftDigit = s[left] - '0'
			break
		}
		if isStartChar(s[left]) && isSpelledOutDigit(s, left, mp) != '0' {
			leftDigit = isSpelledOutDigit(s, left, mp) - '0'
			break
		}

	}

	for right := len(s) - 1; right >= 0; right-- {
		if isDigit(s[right]) {
			rightDigit = s[right] - '0'
			break
		}
		if isStartChar(s[right]) && isSpelledOutDigit(s, right, mp) != '0' {
			rightDigit = isSpelledOutDigit(s, right, mp) - '0'
			break
		}
	}

	value := (int(leftDigit) * 10) + int(rightDigit)
	return value
}

//apparently maps are reference types so idh to pass by ptr?
func initDigitMap(mp map[string]byte) {
	mp["one"] = '1'
	mp["two"] = '2'
	mp["three"] = '3'
	mp["four"] = '4'
	mp["five"] = '5'
	mp["six"] = '6'
	mp["seven"] = '7'
	mp["eight"] = '8'
	mp["nine"] = '9'
}

func main() {
	f, err := os.Open("./input.txt")
	errorHandler(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	totalCalibrationValue := 0
	stringDigitMap := make(map[string]byte)
	initDigitMap(stringDigitMap)
	for scanner.Scan() {
		totalCalibrationValue += extractCalibrationValue(scanner.Text(), stringDigitMap)
		// fmt.Println("Current Value: ", totalCalibrationValue)
	}
	fmt.Println("Done. Total Value is: ", totalCalibrationValue)
}
