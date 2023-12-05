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

// ---- Part 2 Functions ----
// func displayGearsPositions(gearsCoordinates []GearsCordinates) {
// 	for _, gear := range gearsCoordinates {
// 		fmt.Printf("A gear was found at [%d][%d]!\n", gear.X, gear.Y)
// 	}
// }

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

func part2(lines []string,
	gearsPositions []GearsCordinates, maxRows int, maxCols int) (part2 int) {

	digitsNearGears := make([]map[string]GearsCordinates, 3)

	// Inicializar cada mapa
	for i := 0; i < len(digitsNearGears); i++ {
		digitsNearGears[i] = make(map[string]GearsCordinates)
	}

	aux := 0
	for i, line := range lines {
		for j, char := range line {
			// Gear foi detetada
			if string(char) == "*" {
				// Guardar em array todas as posições adjacentes ao * encontrado
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

				// Ver se as posições do array correspondem a um dígito
				for _, pos := range positionsToCheck {
					// Ver se a posição está dentro das dimensões da matriz
					if pos.X >= 0 && pos.Y >= 0 && pos.X <= maxRows-1 && pos.Y <= maxCols-1 {
						// É um número?
						if unicode.IsDigit(int32(lines[pos.X][pos.Y])) {
							key, _ := strconv.Atoi(string(lines[pos.X][pos.Y]))
							digitsNearGears[aux][strconv.Itoa(key)] = GearsCordinates{pos.X, pos.Y}
						}
					}
				}
				fmt.Println()
				aux++
			}
		}
	}

	for _, m := range digitsNearGears {
		if len(m) > 1 {
			for _, value := range m {
				// Procurar resto dos dígitos para formar o número completo
				for _, line := range lines {
					for i := value.Y; i >= 0; i-- {
						char := rune(line[i])
						if unicode.IsDigit(char) {
							// TODO: Find motivation to complete this one
						}
					}
				}
			}
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
	part2Numbers := part2(lines, gearsPositions, numRows, numCols)
	part2Numbers = part2Numbers
	// Obter resultado da Parte 2

	fmt.Printf("Part 1: %d\n", part1_result)
	// fmt.Printf("Part 2: %d\n", part2_result)
}
