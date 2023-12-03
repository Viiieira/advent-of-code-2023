package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
	"unicode"
)

const FILEPATH = "input.txt"

type Coordinates struct {
	Entries []struct {
		StartX int
		StartY int
		EndX   int
		EndY   int
	}
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

func getMatrixSize(matrix []string) (numRows int, numCols int) {
	numRows = len(matrix)

	if numRows > 0 {
		numCols = len(matrix[0])
	}

	return numRows, numCols
}

// func printMatrix(matrix []string) {
// 	for i := range matrix {
// 		for j := range matrix[i] {
// 			fmt.Printf("[%c]", matrix[i][j])
// 		}
// 		fmt.Println()
// 	}
// }

// func printCoordinatesMap(coordsMap map[string]Coordinates) {
// 	for key, value := range coordsMap {
// 		for _, entry := range value.Entries {
// 			fmt.Printf("%s --> (%d,%d) to (%d,%d)\n", key, entry.StartX, entry.StartY, entry.EndX, entry.EndY)
// 		}
// 	}
// }

func findNumbers(s []string) (validNumbers map[string]Coordinates) {
	validNumbers = make(map[string]Coordinates)

	for i, line := range s {
		var currentNumber string
		var startCoord struct{ X, Y int }

		for j, char := range line {
			char = rune(char)
			if unicode.IsDigit(char) {
				if currentNumber == "" {
					startCoord = struct{ X, Y int }{i, j}
				}
				currentNumber += string(char)
			} else if currentNumber != "" {
				number, err := strconv.Atoi(currentNumber)
				if err == nil { // Append to the slice associated with the key
					coords := validNumbers[strconv.Itoa(number)]
					coords.Entries = append(coords.Entries, struct{ StartX, StartY, EndX, EndY int }{startCoord.X, startCoord.Y, i, j - 1})
					validNumbers[strconv.Itoa(number)] = coords
				}
				currentNumber = ""
			}
		}

		if currentNumber != "" {
			number, err := strconv.Atoi(currentNumber)
			if err == nil {
				coords := validNumbers[strconv.Itoa(number)]
				coords.Entries = append(coords.Entries, struct{ StartX, StartY, EndX, EndY int }{startCoord.X, startCoord.Y, i, len(line) - 1})
				validNumbers[strconv.Itoa(number)] = coords
			}
		}

	}

	return
}

func returnValidNumbers(numbers map[string]Coordinates, matrix []string, maxRows int, maxCols int) []int {
	validNumbers := []int{}

	for key, value := range numbers {
		for _, entry := range value.Entries {
			posToCheck := [][]int{
				{entry.StartX, entry.StartY - 1},     // Esquerda
				{entry.StartX, entry.EndY + 1},       // Direita
				{entry.StartX - 1, entry.StartY - 1}, // Cima+Esquerda
				{entry.StartX - 1, entry.EndY + 1},   // Cima+Direita
				{entry.StartX + 1, entry.StartY - 1}, // Baixo+Esquerda
				{entry.StartX + 1, entry.EndY + 1},   // Baixo+Direita
			}
			// Adicionar posições em cima e em baixo
			for i := entry.StartY; i <= entry.EndY; i++ {
				posToCheck = append(posToCheck, []int{entry.StartX - 1, i})
				posToCheck = append(posToCheck, []int{entry.StartX + 1, i})
			}

			for _, pos := range posToCheck {
				if pos[0] >= 0 && pos[1] >= 0 && pos[0] <= maxRows-1 && pos[1] <= maxCols-1 {
					char := rune(matrix[pos[0]][pos[1]])
					if !unicode.IsDigit(char) && char != '.' {
						validNumber, _ := strconv.Atoi(key)
						validNumbers = append(validNumbers, validNumber)
					}
				}
			}
		}
	}
	return validNumbers
}

func sumAllValidNumbers(validNumbers []int) int {
	sum := 0
	for _, num := range validNumbers {
		sum += num
	}
	return sum
}

func trackTime(name string) func() {
	startTime := time.Now()
	fmt.Printf("[%s] Started...\n", name)

	return func() {
		elapsed := time.Since(startTime)
		fmt.Printf("[%s] Took %s to run.\n", name, elapsed)
	}
}

func main() {
	defer trackTime("Part 1")()

	lines, err := readLinesFromFile(FILEPATH)

	if err != nil {
		log.Fatal(err)
	}

	numRows, numCols := getMatrixSize(lines)

	nums := findNumbers(lines)
	validNumbers := returnValidNumbers(nums, lines, numRows, numCols)
	part1_result := sumAllValidNumbers(validNumbers)

	fmt.Printf("Part 1: %d\n", part1_result)
}
