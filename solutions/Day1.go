package solutions

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Day1Part1() {
	println("(Part 1) Sum of all numbers:", sumNumbers(false))
}

func Day1Part2() {
	println("(Part 2) Sum of all numbers:", sumNumbers(true))
}

func sumNumbers(translateStrings bool) int {
	file, err := os.Open("solutions/Day1Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if translateStrings {
			line = translateNumbers(line)
		}
		re := regexp.MustCompile("[0-9]+")
		matches := re.FindAllString(line, -1)
		allNums := ""
		for _, match := range matches {
			allNums += match
		}

		var stringResult = string(allNums[0]) + string(allNums[len(allNums)-1])
		intResult, _ := strconv.Atoi(stringResult)

		sum += intResult
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func translateNumbers(input string) string {
	output := ""
	for i := 0; i < len(input); i++ {
		substring := input[i:]
		match, _ := regexp.MatchString("^one", substring)
		if match {
			output += "1"
			continue
		}
		match, _ = regexp.MatchString("^two", substring)
		if match {
			output += "2"
			continue
		}
		match, _ = regexp.MatchString("^three", substring)
		if match {
			output += "3"
			continue
		}
		match, _ = regexp.MatchString("^four", substring)
		if match {
			output += "4"
			continue
		}
		match, _ = regexp.MatchString("^five", substring)
		if match {
			output += "5"
			continue
		}
		match, _ = regexp.MatchString("^six", substring)
		if match {
			output += "6"
			continue
		}
		match, _ = regexp.MatchString("^seven", substring)
		if match {
			output += "7"
			continue
		}
		match, _ = regexp.MatchString("^eight", substring)
		if match {
			output += "8"
			continue
		}
		match, _ = regexp.MatchString("^nine", substring)
		if match {
			output += "9"
			continue
		}
		match, _ = regexp.MatchString("^zero", substring)
		if match {
			output += "0"
			continue
		}

		output += string(input[i])
	}

	return output
}
