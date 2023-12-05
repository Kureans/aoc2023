package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type AgricultureMap struct {
	mapEntries []MapEntry
}

type MapEntry struct {
	source   int
	dest     int
	rangeLen int
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
func createMapEntry(dest int, source int, rangeLen int) MapEntry {
	var entry MapEntry
	entry.source = source
	entry.dest = dest
	entry.rangeLen = rangeLen
	return entry
}

func (entry MapEntry) String() string {
	return fmt.Sprintf("Source: %d, Dest: %d, Range: %d", entry.source, entry.dest, entry.rangeLen)
}

type AlmanacParser struct {
	seeds      []int
	seedRanges []SeedRange
	locations  []int
	maps       []AgricultureMap
	mapIndex   int
}

type SeedRange struct {
	start    int
	rangeLen int
}

// Part 1 function
// func (parser *AlmanacParser) getLowestLocation() int {
// 	lowestLocation := math.MaxInt32
// 	for _, seed := range parser.seeds {
// 		lowestLocation = min(lowestLocation, parser.getLocation(seed))
// 	}
// 	return lowestLocation
// }

func (parser *AlmanacParser) getLowestLocation() int {
	for lowestLocation := 0; ; lowestLocation++ {
		if parser.hasValidSeed(lowestLocation) {
			return lowestLocation
		}
	}
}

func (parser *AlmanacParser) hasValidSeed(location int) bool {
	currIdx := location
	for i := 6; i >= 0; i-- {
		currentMap := parser.maps[i]
		for _, entry := range currentMap.mapEntries {
			if currIdx < entry.source || currIdx >= entry.source+entry.rangeLen {
				continue
			}
			newIdx := entry.dest + (currIdx - entry.source)
			// fmt.Printf("Index %d, Old Idx: %d, Updated Idx: %d\n", i, currIdx, newIdx)
			currIdx = newIdx
			break
		}
	}
	return isInSeedRange(currIdx, parser.seedRanges)
}

func isInSeedRange(val int, ranges []SeedRange) bool {
	for _, r := range ranges {
		if val >= r.start && val < (r.start+r.rangeLen) {
			return true
		}
	}
	return false
}

func (parser *AlmanacParser) getLocation(seed int) int {
	currIdx := seed
	for i := 0; i < 7; i++ {
		currentMap := parser.maps[i]
		for _, entry := range currentMap.mapEntries {
			if currIdx < entry.source || currIdx >= entry.source+entry.rangeLen {
				continue
			}
			newIdx := entry.dest + (currIdx - entry.source)
			// fmt.Print(entry)
			// fmt.Printf("Index %d, Old Idx: %d, Updated Idx: %d\n", i, currIdx, newIdx)
			currIdx = newIdx

			break
		}
		// fmt.Printf("Index %d, Same Idx: %d\n", i, currIdx)
	}
	// fmt.Printf("Seed: %d, Location: %d\n", seed, currIdx)
	return currIdx
}

func (parser *AlmanacParser) parseInput(inputArr []string) {
	if len(inputArr) == 1 {
		parser.addValuesToMap(inputArr[0])
	} else {
		if inputArr[0] == "seeds" {
			parser.initialiseSeedRanges(inputArr[1])
		} else {
			parser.initialiseMap()
		}
	}
}

// Part 1 Version of initialiseSeeds
// func (parser *AlmanacParser) initialiseSeeds(seedsString string) {
// 	seedArr := strings.Fields(seedsString)
// 	parser.seeds = make([]int, len(seedArr))
// 	for i, seedStr := range seedArr {
// 		v, err := strconv.Atoi(seedStr)
// 		errorHandler(err)
// 		parser.seeds[i] = v
// 	}
// }

func (parser *AlmanacParser) initialiseSeeds(seedsString string) {
	seedArr := strings.Fields(seedsString)
	parser.seeds = make([]int, 0)
	var currSeed int
	for i, seedStr := range seedArr {
		v, err := strconv.Atoi(seedStr)
		errorHandler(err)
		if i%2 == 0 {
			currSeed = v
		} else {
			for seed := currSeed; seed < (currSeed + v); seed++ {
				parser.seeds = append(parser.seeds, seed)
			}
		}
	}
}

func (parser *AlmanacParser) initialiseSeedRanges(seedsString string) {
	seedArr := strings.Fields(seedsString)
	parser.seedRanges = make([]SeedRange, len(seedArr)/2)
	for i := 0; i < len(seedArr); i += 2 {
		start, err := strconv.Atoi(seedArr[i])
		rangeLen, err := strconv.Atoi(seedArr[i+1])
		errorHandler(err)
		var seedRange SeedRange
		seedRange.start = start
		seedRange.rangeLen = rangeLen
		parser.seedRanges[i/2] = seedRange
	}
}

func (parser *AlmanacParser) initialiseMap() {
	parser.mapIndex++
	parser.maps[parser.mapIndex].mapEntries = make([]MapEntry, 5)
}

func (parser *AlmanacParser) addValuesToMap(mappingsString string) {
	valueArr := strings.Fields(mappingsString)
	if len(valueArr) != 3 {
		panic(len(valueArr))
	}
	dest, err := strconv.Atoi(valueArr[0])
	source, err := strconv.Atoi(valueArr[1])
	rangeLen, err := strconv.Atoi(valueArr[2])

	errorHandler(err)

	parser.maps[parser.mapIndex].mapEntries = append(parser.maps[parser.mapIndex].mapEntries, createMapEntry(source, dest, rangeLen))
}

func (parser *AlmanacParser) initParser() {
	parser.mapIndex = -1
	parser.maps = make([]AgricultureMap, 7)
}

func errorHandler(e error) {
	if e != nil {
		fmt.Println("Error Encountered")
		panic(e)
	}
}

func part1() {
	f, err := os.Open("./input.txt")
	errorHandler(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	var parser AlmanacParser
	parser.initParser()
	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		inputArr := strings.Split(scanner.Text(), ":")
		parser.parseInput(inputArr)
	}
	// for i, j := 0, len(parser.maps)-1; i < j; i, j = i+1, j-1 {
	// 	parser.maps[i], parser.maps[j] = parser.maps[j], parser.maps[i]
	// }
	lowestLocation := parser.getLowestLocation()
	fmt.Println("Lowest Location: ", lowestLocation)
}

func main() {
	part1()
}
