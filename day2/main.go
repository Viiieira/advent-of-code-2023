package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

// TODO: convert cubeID to cubeAmount

func main() {
	filePath := "example.txt"
	games := make(map[int]map[int]map[int]string)
	lines, err := readLinesFromFile(filePath)

	if err != nil {
		log.Fatal(err)
	}

	games = fillGamesMap(games, lines)

	printGames(games)

	fmt.Printf("Part 1: %d\n", part1(games))
	fmt.Printf("Part 2: %d\n", part2(games))
}

// Função para inserir os valores para o mapa
func insertIntoMap(games map[int]map[int]map[int]string, gameID, setID, cubeID int, cubeColor string) map[int]map[int]map[int]string {
	if games[gameID] == nil {
		games[gameID] = make(map[int]map[int]string)
	}

	if games[gameID][setID] == nil {
		games[gameID][setID] = make(map[int]string)
	}

	games[gameID][setID][cubeID] = cubeColor
	return games
}

func fillGamesMap(games map[int]map[int]map[int]string, lines []string) map[int]map[int]map[int]string {
	for _, line := range lines {
		// Get the Game ID
		indexOfTwoPoints := strings.Index(line, ":")
		indexOfGameLabel := strings.Index(line, "Game ")
		gameID, _ := strconv.Atoi(line[indexOfGameLabel+5 : indexOfTwoPoints])

		// Get the Number of Sets
		sets := strings.Split(line[indexOfTwoPoints+2:], ";")
		cubes := []string{}
		for setID, set := range sets {
			cubes = strings.Split(set, ",")
			for _, cube := range cubes {
				cube = strings.Trim(cube, " ")
				cube_info := strings.Split(cube, " ")
				cubeID, _ := strconv.Atoi(cube_info[0])
				cubeColor := cube_info[1]

				games = insertIntoMap(games, gameID, setID, cubeID, cubeColor)
			}
		}
	}

	return games
}

func printGames(games map[int]map[int]map[int]string) {
	for gameID, sets := range games {
		fmt.Printf("Game %d:\n", gameID)
		for set, cubes := range sets {
			fmt.Printf("\tSet %d:\n", set)
			for cubeID, cubeColor := range cubes {
				fmt.Printf("\t\t%d Cubes: %s\n", cubeID, cubeColor)
			}
		}
	}
}

func part1(games map[int]map[int]map[int]string) (sum int) {
	for gameID, sets := range games {

		max_red, max_green, max_blue := 0, 0, 0
		for _, cubes := range sets {
			for cubeID, cubeColor := range cubes {
				if cubeColor == "red" && cubeID > max_red {
					max_red = cubeID
				}
				if cubeColor == "green" && cubeID > max_green {
					max_green = cubeID
				}
				if cubeColor == "blue" && cubeID > max_blue {
					max_blue = cubeID
				}
			}
		}
		if max_red <= 12 && max_green <= 13 && max_blue <= 14 {
			// Jogo válido
			sum += gameID
		}
	}
	return
}

func part2(games map[int]map[int]map[int]string) (sum int) {
	for _, sets := range games {
		max_red, max_green, max_blue := 0, 0, 0
		for _, cubes := range sets {
			for cubeID, cubeColor := range cubes {
				fmt.Printf("> %d %s\n", cubeID, cubeColor)
				if cubeColor == "red" && cubeID > max_red {
					max_red = cubeID
				}
				if cubeColor == "green" && cubeID > max_green {
					max_green = cubeID
				}
				if cubeColor == "blue" && cubeID > max_blue {
					max_blue = cubeID
				}
			}
		}
		fmt.Printf("Max_red %d, Max_green %d, Max_blue %d\n", max_red, max_green, max_blue)
		sum += max_red * max_green * max_blue
	}
	return
}
