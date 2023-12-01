package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

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

func findNumber(s string, mode int) int {

	// Ler linha do início
	if mode == 1 {
		for i := 0; i <= len(s)-1; i++ {
			char := rune(s[i])
			if unicode.IsDigit(char) {
				return int(char - '0')
			}
		}
	}

	// Ler linha desde o início
	if mode == 2 {
		for i := len(s) - 1; i >= 0; i-- {
			char := rune(s[i])
			if unicode.IsDigit(char) {
				return int(char - '0')
			}
		}
	}

	return 0
}

func main() {
	filePath := "input.txt"
	lines, err := readLinesFromFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	calibration_sum := 0
	for _, line := range lines {
		calibration_value :=
			strconv.Itoa(findNumber(line, 1)) +
				strconv.Itoa(findNumber(line, 2))
		num, err := strconv.Atoi(calibration_value)
		if err != nil {
			log.Fatalf("Error converting string to integer: %v", err)
		}

		calibration_sum += num
	}

	fmt.Println("Calibration Sum: ", calibration_sum)
}
