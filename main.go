package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

func solvePart1(pathsMap map[string][]string) {
	pathsCount := 0
	traverse1(pathsMap, "start", []string{}, &pathsCount)

	fmt.Println("PART 1 ANSWER")
	fmt.Println(pathsCount)
}

func traverse1(pathsMap map[string][]string, currCave string, visited []string, pathsCount *int) {
	if !canVisit1(visited, currCave) {
		return
	}

	visited = append(visited, currCave)
	if currCave == "end" {
		*pathsCount++
		return
	}

	nextPositions, ok := pathsMap[currCave]
	if ok {
		for _, nextPosition := range nextPositions {
			traverse1(pathsMap, nextPosition, visited, pathsCount)
		}
	}
}

// canVisit1 checks if a cave can be visited if it has never been visited or if it has been visited, it must be a big cave.
func canVisit1(visited []string, cave string) bool {
	for _, pos := range visited {
		if pos == cave && isSmallCave(cave) {
			return false
		}
	}

	return true
}

func solvePart2(pathsMap map[string][]string) {
	pathsCount := 0
	traverse2(pathsMap, "start", []string{}, &pathsCount)

	fmt.Println("PART 2 ANSWER")
	fmt.Println(pathsCount)
}

func traverse2(pathsMap map[string][]string, currCave string, visited []string, pathsCount *int) {
	if !canVisit2(visited, currCave) {
		return
	}

	visited = append(visited, currCave)
	if currCave == "end" {
		*pathsCount++
		return
	}

	nextPositions, ok := pathsMap[currCave]
	if ok {
		for _, nextPosition := range nextPositions {
			traverse2(pathsMap, nextPosition, visited, pathsCount)
		}
	}
}

// canVisit2 checks if a cave can be visited if it has never been visited or if it has been visited, it must be a big cave.
func canVisit2(visited []string, cave string) bool {
	hasVisitedSmallCaveTwice := false
	caveVisitCounts := make(map[string]int)
	for _, pos := range visited {
		_, ok := caveVisitCounts[pos]

		if ok {
			caveVisitCounts[pos]++
		} else {
			caveVisitCounts[pos] = 1
		}

		if isSmallCave(pos) && caveVisitCounts[pos] == 2 {
			hasVisitedSmallCaveTwice = true
		}
	}

	for _, pos := range visited {
		if pos == cave && isSmallCave(cave) {
			if cave == "start" || cave == "end" {
				// start and end cave can only be visited once.
				return false
			}

			if hasVisitedSmallCaveTwice {
				return false
			}
		}
	}

	return true
}

// isSmallCave checks if the cave is a small cave. It does this by ensuring that all the characters are lowercase.
func isSmallCave(cave string) bool {
	for _, chr := range cave {
		if unicode.IsUpper(chr) {
			return false
		}
	}

	return true
}

func main() {
	content, err := os.ReadFile("input.txt")
	if err != nil {
		log.Printf("Failed to read input file: %s\n", err)
		return
	}

	// Parse the content, split every newline
	paths := strings.Split(string(content), "\n")

	// Create a map of paths.
	pathsMap := make(map[string][]string)
	for _, path := range paths {
		points := strings.Split(path, "-")
		origin := points[0]
		dest := points[1]

		_, ok := pathsMap[origin]
		if !ok {
			pathsMap[origin] = make([]string, 0)
		}
		pathsMap[origin] = append(pathsMap[origin], dest)

		_, ok = pathsMap[dest]
		if !ok {
			pathsMap[dest] = make([]string, 0)
		}
		pathsMap[dest] = append(pathsMap[dest], origin)
	}

	solvePart1(pathsMap)
	solvePart2(pathsMap)
}
