package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filePath := "input.txt"
	lines, err := readLinesFromFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	power := 0

	// Percorrer todas as linhas
	for _, line := range lines {
		max_red, max_green, max_blue := 0, 0, 0

		indexSepSets := strings.Index(line, ":")
		sets := strings.Split(line[indexSepSets+2:], ";")
		for _, set := range sets {
			set = strings.Trim(set, " ")

			cubes := strings.Split(set, ",")
			for _, cube := range cubes {
				cube = strings.Trim(cube, " ")

				cube_data := strings.Split(cube, " ")
				cubeAmount, _ := strconv.Atoi(cube_data[0])
				cubeColor := cube_data[1]
				if cubeColor == "red" && cubeAmount > max_red {
					max_red = cubeAmount
				}
				if cubeColor == "green" && cubeAmount > max_green {
					max_green = cubeAmount
				}
				if cubeColor == "blue" && cubeAmount > max_blue {
					max_blue = cubeAmount
				}
			}
		}
		power += (max_red * max_green * max_blue)
	}

	fmt.Printf("Part 2: %d\n", power)
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
