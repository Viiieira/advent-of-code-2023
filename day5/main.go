package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FILEPATH = "example.txt"

type SourceDestination struct {
	source                string
	destination           string
	destinationRangeStart []int
	sourceRangeStart      []int
	rangeLength           []int
}

func readLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func getSeeds(lines []string) (seedsInt []int) {
	seedsStr := strings.Trim(strings.Split(lines[0], ": ")[1], " ")
	seeds := strings.Split(seedsStr, " ")

	for _, seed := range seeds {
		seedInt, _ := strconv.Atoi(seed)
		seedsInt = append(seedsInt, seedInt)
	}

	return
}

func getSourcesDestinations(lines []string) (sourcesDestinations []SourceDestination) {

	// Get the map titles index
	indexMapTitles := []int{}
	for i, line := range lines {
		if i >= 2 && line != "" {
			if strings.Contains(line, " map:") {
				indexMapTitles = append(indexMapTitles, i)
			}
		}
	}

	op := 0
	for i := 0; i < len(indexMapTitles); i++ {
		start := indexMapTitles[i]
		end := 0
		if i+1 >= len(indexMapTitles) {
			end = len(lines) - 1
		} else {
			end = indexMapTitles[i+1] - 2
		}

		op = 0
		// Print the desired range
		source, destination := "", ""
		destinationRangeStart, sourceRangeStart, rangeLength := []int{}, []int{}, []int{}
		for j := start; j <= end; j++ {
			if op == 0 {
				// Retrieve the map-title data
				titleStr := strings.Split(lines[j], " map:")[0]
				source = strings.Split(titleStr, "-to-")[0]
				destination = strings.Split(titleStr, "-to-")[1]
			} else {
				// Retrieve the map numbers
				lineNums := strings.Split(lines[j], " ")
				auxDest, _ := strconv.Atoi(lineNums[0])
				destinationRangeStart = append(destinationRangeStart, auxDest)
				auxSrc, _ := strconv.Atoi(lineNums[1])
				sourceRangeStart = append(sourceRangeStart, auxSrc)
				auxLen, _ := strconv.Atoi(lineNums[2])
				rangeLength = append(rangeLength, auxLen)
			}
			op++
		}
		newMapEntry := SourceDestination{
			source, destination, destinationRangeStart, sourceRangeStart, rangeLength,
		}
		sourcesDestinations = append(sourcesDestinations, newMapEntry)
	}

	return
}

func displaySourcesDestinations(sourcesDestinations []SourceDestination) {
	for _, sD := range sourcesDestinations {
		fmt.Printf("%s-%s:\n", sD.source, sD.destination)
		for i := 0; i <= len(sD.sourceRangeStart)-1; i++ {
			fmt.Printf("%d %d %d\n",
				sD.destinationRangeStart[i],
				sD.sourceRangeStart[i],
				sD.rangeLength[i])
		}
		fmt.Println()
		fmt.Println()
	}
}

func main() {
	lines, _ := readLinesFromFile(FILEPATH)

	seeds := getSeeds(lines)
	fmt.Printf("Seeds: %v\n\n", seeds)

	sourcesDestinations := getSourcesDestinations(lines)
	// displaySourcesDestinations(sourcesDestinations)

	for _, seed := range seeds {
		lineSeed := []int{}
		// fmt.Printf("> %d\n", seed)
		for _, sd := range sourcesDestinations {
			leftNumbers := []int{}
			rightNumbers := []int{}

			for i := 0; i <= len(sd.sourceRangeStart)-1; i++ {
				// Source Ranges
				for j := sd.sourceRangeStart[i]; j <= (sd.sourceRangeStart[i] + (sd.rangeLength[i] - 1)); j++ {
					leftNumbers = append(leftNumbers, j)
				}
				// Destination Ranges
				for j := sd.destinationRangeStart[i]; j <= (sd.destinationRangeStart[i] + (sd.rangeLength[i] - 1)); j++ {
					rightNumbers = append(rightNumbers, j)
				}
			}
			// fmt.Printf("%s, %s\n", sd.source, sd.destination)
			found := false
			lineSeed = append(lineSeed, seed)
			for j := 0; j <= len(leftNumbers)-1; j++ {
				if seed == leftNumbers[j] {
					fmt.Printf("%d\t%d\n", leftNumbers[j], rightNumbers[j])
					found = true
					lineSeed = append(lineSeed, rightNumbers[j])
				}
			}
			if !found {
				lineSeed = append(lineSeed, seed)
			}
			// fmt.Println()
		}
		fmt.Println(lineSeed)
	}
}
