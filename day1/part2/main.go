package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func readFile(filePath string) ([]string, error) {
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

func containsMapKey(s string, m map[string]int) (int, bool) {
	for key, value := range m {
		if strings.Contains(s, key) {
			return value, true
		}
	}

	return 0, false
}

func findNumber(s string, mode int, m map[string]int) int {

	if mode == 1 {
		newString := ""

		for i := 0; i <= len(s)-1; i++ {
			char := rune(s[i])
			if unicode.IsDigit(char) {
				return int(char - '0')
			} else {
				newString += string(s[i])

				digit, exists := containsMapKey(newString, m)
				if exists {
					return digit
				}
			}
		}
	}

	if mode == 2 {
		newString := ""

		for i := len(s) - 1; i >= 0; i-- {
			char := rune(s[i])
			if unicode.IsDigit(char) {
				return int(char - '0')
			} else {
				newString = string(s[i]) + newString

				digit, exists := containsMapKey(newString, m)
				if exists {
					return digit
				}
			}
		}
	}

	return 0
}

func main() {
	allDigits := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"zero":  0,
	}

	filePath := "../input.txt"
	lines, err := readFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	calibration_sum := 0
	for _, line := range lines {
		calibration_value :=
			strconv.Itoa(findNumber(line, 1, allDigits)) +
				strconv.Itoa(findNumber(line, 2, allDigits))
		num, err := strconv.Atoi(calibration_value)
		if err != nil {
			log.Fatalf("Error converting string to integer: %v", err)
		}

		calibration_sum += num
	}

	fmt.Println("Calibration Sum: ", calibration_sum)
}
