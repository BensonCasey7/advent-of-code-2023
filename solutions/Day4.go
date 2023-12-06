package solutions

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

type ScratchCard struct {
	WinningNumbers []string
	ActualNumbers  []string
	count          int
}

func NewScratchCard() ScratchCard {
	scratchCard := ScratchCard{}
	scratchCard.WinningNumbers = []string{}
	scratchCard.ActualNumbers = []string{}
	scratchCard.count = 1
	return scratchCard
}

func Day4Part1() {
	file, err := os.Open("solutions/Day4Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		cardNumberRegex := regexp.MustCompile("Card.*[0-9]+: ")
		strippedLine := cardNumberRegex.ReplaceAllString(line, "")
		numberGroups := strings.Split(strippedLine, " | ")
		scratchCard := NewScratchCard()

		numberRegex := regexp.MustCompile("[0-9]+")

		scratchCard.WinningNumbers = numberRegex.FindAllString(numberGroups[0], -1)
		scratchCard.ActualNumbers = numberRegex.FindAllString(numberGroups[1], -1)
		score := calculateCardScore(scratchCard)

		sum += score
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	println("(Part 1) Sum of all scores:", sum)
}

func calculateCardScore(scratchCard ScratchCard) int {
	score := 0
	for _, number := range scratchCard.WinningNumbers {
		for _, actualNumber := range scratchCard.ActualNumbers {
			if number == actualNumber {
				if score == 0 {
					score = 1
				} else {
					score *= 2
				}
			}
		}
	}
	return score
}

func Day4Part2() {
	file, err := os.Open("solutions/Day4Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)

	scratchCards := []ScratchCard{}
	for scanner.Scan() {
		line := scanner.Text()

		cardNumberRegex := regexp.MustCompile("Card.*[0-9]+: ")
		strippedLine := cardNumberRegex.ReplaceAllString(line, "")
		numberGroups := strings.Split(strippedLine, " | ")
		scratchCard := NewScratchCard()

		numberRegex := regexp.MustCompile("[0-9]+")

		scratchCard.WinningNumbers = numberRegex.FindAllString(numberGroups[0], -1)
		scratchCard.ActualNumbers = numberRegex.FindAllString(numberGroups[1], -1)

		scratchCards = append(scratchCards, scratchCard)
	}

	for i, scratchCard := range scratchCards {
		for times := 0; times < scratchCard.count; times++ {
			wins := 0
			for _, number := range scratchCard.WinningNumbers {
				for _, actualNumber := range scratchCard.ActualNumbers {
					if number == actualNumber {
						wins++
						if i+wins < len(scratchCards) {
							scratchCards[i+wins].count++
						}
					}
				}
			}
		}
	}

	for _, scratchCard := range scratchCards {
		sum += scratchCard.count
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	println("(Part 2) Sum of all cards:", sum)
}
