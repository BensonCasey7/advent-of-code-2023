package day2

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	Blue  int
	Green int
	Red   int
}

func NewGame() Game {
	game := Game{}
	game.Blue = 0
	game.Green = 0
	game.Red = 0
	return game
}

func Part1() {
	gameMaximums := calculateGameMaximums()

	gameLimits := NewGame()
	gameLimits.Red = 12
	gameLimits.Green = 13
	gameLimits.Blue = 14

	validGameIdsSum := 0
	for index, gameMaximum := range gameMaximums {
		if gameMaximum.Red <= gameLimits.Red && gameMaximum.Green <= gameLimits.Green && gameMaximum.Blue <= gameLimits.Blue {
			id := index + 1
			validGameIdsSum += id
		}
	}

	println("(Part 1) Sum of all game IDs:", validGameIdsSum)
}
func Part2() {
	gameMaximums := calculateGameMaximums()

	gamePowerSums := 0
	for _, gameMaximum := range gameMaximums {
		gamePower := gameMaximum.Red * gameMaximum.Green * gameMaximum.Blue
		gamePowerSums += gamePower
	}

	println("(Part 2) Sum of all game powers:", gamePowerSums)
}

func calculateGameMaximums() []Game {
	file, err := os.Open("day2/Input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []Game{}
	}
	defer file.Close()

	lineCount := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineCount++
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return []Game{}
	}

	// Reopen the file to process each line
	file, err = os.Open("day2/Input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []Game{}
	}
	defer file.Close()

	gameMaximums := make([]Game, lineCount)

	// Create a new scanner to process each line and populate the array
	scanner = bufio.NewScanner(file)
	index := 0

	for scanner.Scan() {
		line := scanner.Text()

		gameMaximum := NewGame()

		regex := regexp.MustCompile("Game \\d+: ")
		strippedString := regex.ReplaceAllString(line, "")

		gameRounds := strings.Split(strippedString, "; ")
		for _, gameRound := range gameRounds {
			round := parseGameRound(gameRound)
			if round.Blue > gameMaximum.Blue {
				gameMaximum.Blue = round.Blue
			}
			if round.Green > gameMaximum.Green {
				gameMaximum.Green = round.Green
			}
			if round.Red > gameMaximum.Red {
				gameMaximum.Red = round.Red
			}
		}
		gameMaximums[index] = gameMaximum
		index++
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return []Game{}
	}

	return gameMaximums
}

func parseGameRound(input string) Game {
	game := NewGame()
	blueRegex := regexp.MustCompile("\\d+ blue")
	greenRegex := regexp.MustCompile("\\d+ green")
	redRegex := regexp.MustCompile("\\d+ red")

	blueMatch := blueRegex.FindString(input)
	greenMatch := greenRegex.FindString(input)
	redMatch := redRegex.FindString(input)

	blueMatch = strings.Replace(blueMatch, " blue", "", -1)
	greenMatch = strings.Replace(greenMatch, " green", "", -1)
	redMatch = strings.Replace(redMatch, " red", "", -1)

	blueInt, _ := strconv.Atoi(blueMatch)
	greenInt, _ := strconv.Atoi(greenMatch)
	redInt, _ := strconv.Atoi(redMatch)

	game.Blue = blueInt
	game.Green = greenInt
	game.Red = redInt

	return game
}
