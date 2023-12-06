package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const FILEPATH = "input.txt"

type ScratchCard struct {
	cardID        int
	winningValues []int
	myValues      []int
	matches       int
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

func getWinningValuesCard(line string) []int {
	leftSide := strings.Trim(strings.Split(line, "|")[0], " ")
	leftNumbers := strings.Split(leftSide, " ")

	winningValues := []int{}
	for _, value := range leftNumbers {
		num, _ := strconv.ParseInt(strings.Trim(value, " "), 10, 64)
		winningValues = append(winningValues, int(num))
	}

	return winningValues
}

func getMyValuesCard(line string) []int {
	rightSide := strings.Trim(strings.Split(line, "|")[1], " ")
	rightNumbers := strings.Split(rightSide, " ")

	myValues := []int{}
	for _, value := range rightNumbers {
		num, _ := strconv.ParseInt(strings.Trim(value, " "), 10, 64)
		if int(num) != 0 {
			myValues = append(myValues, int(num))
		}
	}

	return myValues
}

func getPart1Result(winningValuesCard []int, myValuesCard []int) (result int) {
	pointsMatch := 0
	for _, winningValue := range winningValuesCard {
		for _, myValue := range myValuesCard {
			if winningValue == myValue {
				if pointsMatch == 0 {
					pointsMatch++
				} else {
					pointsMatch *= 2
				}
			}
		}
	}
	result += pointsMatch
	return
}

// Part 2 stuff
func getPart2Result(scratchCards []ScratchCard) (result int) {
	copies := make([][]int, len(scratchCards))
	score := make([]int, len(scratchCards))

	// Contar matches por carta
	for i, card := range scratchCards {
		matchesPerCard := 0

		// Get the matches per card
		for _, winningValue := range card.winningValues {
			for _, myValue := range card.myValues {
				if winningValue == myValue {
					matchesPerCard++
				}
			}
		}

		for j := i + 1; j <= i+matchesPerCard; j++ {
			copies[i] = append(copies[i], j)
		}
	}

	for i := range score {
		score[i] = 1
	}

	for i := len(scratchCards) - 1; i >= 0; i-- {
		for _, j := range copies[i] {
			score[i] += score[j]
		}
	}

	totalScore := 0
	for _, s := range score {
		totalScore += s
	}

	return totalScore
}

func main() {
	lines, _ := readLinesFromFile(FILEPATH)

	winningValuesCard, myValuesCard := []int{}, []int{}
	part1_result := 0
	scratchCards := []ScratchCard{}

	for cardID, line := range lines {
		line = strings.Trim(strings.Split(line, ":")[1], " ")
		winningValuesCard = getWinningValuesCard(line)
		myValuesCard = getMyValuesCard(line)
		part1_result += getPart1Result(winningValuesCard, myValuesCard)

		scratchCards = append(scratchCards, ScratchCard{cardID + 1, winningValuesCard, myValuesCard, 0})
	}

	part2_result := getPart2Result(scratchCards)

	fmt.Printf("Part 1: %d\n", part1_result)
	fmt.Printf("Part 2: %d\n", part2_result)

}
