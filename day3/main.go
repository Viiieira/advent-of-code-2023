package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

const FILEPATH = "example.txt"

type Coordinates struct {
	Entries []struct {
		StartX, StartY, EndX, EndY int
	}
}

type GearsCordinates struct {
	X, Y int
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

func getAllGearsPositions(lines []string) (gearsPositions []GearsCordinates) {
	for i, line := range lines {
		for j, char := range line {
			if string(char) == "*" {
				gearsPositions = append(gearsPositions, GearsCordinates{i, j})
			}
		}
	}

	return gearsPositions
}

func part2(lines []string, gearsPositions []GearsCordinates, maxRows int, maxCols int) (part2 int) {
	digitsNearGears, aux := make([]map[string]GearsCordinates, 3), 0

	for i := 0; i < len(digitsNearGears); i++ {
		digitsNearGears[i] = make(map[string]GearsCordinates)
	}

	for i, line := range lines {
		for j, char := range line {
			if string(char) == "*" {
				positionsToCheck := []GearsCordinates{
					{i, j - 1},     // Esquerda
					{i, j + 1},     // Direita
					{i + 1, j},     // Baixo
					{i - 1, j},     // Cima
					{i + 1, j - 1}, // Baixo + Esquerda
					{i + 1, j + 1}, // Baixo + Direita
					{i - 1, j - 1}, // Cima + Esquerda
					{i - 1, j + 1}, // Cima + Direita
				}

				for _, pos := range positionsToCheck {
					if pos.X >= 0 && pos.Y >= 0 && pos.X <= maxRows-1 && pos.Y <= maxCols-1 {
						if unicode.IsDigit(int32(lines[pos.X][pos.Y])) {
							key, _ := strconv.Atoi(string(lines[pos.X][pos.Y]))
							digitsNearGears[aux][strconv.Itoa(key)] = GearsCordinates{pos.X, pos.Y}
						}
					}
				}
				aux++
			}
		}
	}

	for _, m := range digitsNearGears {
		// Apenas contar estrelas com pelo menos 2 números adjacentes
		if len(m) >= 2 {
			// Ver todos os dígitos
			twoNumbers := []int{}
			for _, coords := range m {
				completeNumberStr := ""
				for i := range lines {
					// Linha encontrada para encontrar o resto do número
					if i == coords.X {
						for j := coords.Y; j >= 0; j-- {
							char := rune(lines[i][j])
							if unicode.IsDigit(char) {
								completeNumberStr = string(lines[i][j]) + completeNumberStr
							} else {
								break
							}
						}
						for j := coords.Y + 1; j <= len(lines[i])-1; j++ {
							char := rune(lines[i][j])
							if unicode.IsDigit(char) {
								completeNumberStr += string(lines[i][j])
							} else {
								break
							}
						}
						completeNumber, _ := strconv.Atoi(completeNumberStr)

						// Check if the number doesnt already exist
						exists := false
						for _, num := range twoNumbers {
							if num == completeNumber {
								exists = true
								break
							}
						}
						if !exists {
							twoNumbers = append(twoNumbers, completeNumber)
						}

					}
				}
			}
			part2 += (twoNumbers[0] * twoNumbers[1])
		}
	}

	return
}

func main() {
	lines, err := readLinesFromFile(FILEPATH)

	if err != nil {
		log.Fatal(err)
	}

	numRows, numCols := getMatrixSize(lines)

	nums := findNumbers(lines)
	part1ValidNumbers := returnValidNumbers(nums, lines, numRows, numCols)
	part1_result := sumAllValidNumbers(part1ValidNumbers)

	// Encontrar as posições das gears
	gearsPositions := getAllGearsPositions(lines)
	// Obter resultado da Parte 2
	part2_result := part2(lines, gearsPositions, numRows, numCols)

	fmt.Printf("Part 1: %d\n", part1_result)
	fmt.Printf("Part 2: %d\n", part2_result)
}
