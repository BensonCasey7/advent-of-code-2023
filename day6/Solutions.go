package day6

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Record struct {
	Time     int
	Distance int
}

func Part1() {
	file, err := os.Open("day6/Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	records := []Record{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		timesRegex := regexp.MustCompile("^Time:")
		numberRegex := regexp.MustCompile("[0-9]+")
		if timesRegex.MatchString(line) {
			matches := numberRegex.FindAllString(line, -1)
			for _, match := range matches {
				num, _ := strconv.Atoi(match)
				records = append(records, Record{
					Time: num,
				})
			}
			continue
		}

		distancesRegex := regexp.MustCompile("^Distance:")
		if distancesRegex.MatchString(line) {
			matches := numberRegex.FindAllString(line, -1)
			for i, match := range matches {
				num, _ := strconv.Atoi(match)
				records[i].Distance = num
			}
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	product := 1
	for _, record := range records {
		winCount := 0
		for i := 1; i < record.Time; i++ {
			distance := (record.Time - i) * i
			if distance > record.Distance {
				winCount++
			}
		}
		product *= winCount
	}

	println("(Part 1) product of winning scenarios:", product)
}

func Part2() {
	file, err := os.Open("day6/Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	record := Record{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		timesRegex := regexp.MustCompile("^Time:")
		numberRegex := regexp.MustCompile("[0-9]+")
		if timesRegex.MatchString(line) {
			matches := numberRegex.FindAllString(line, -1)
			joinedMatches := strings.Join(matches, "")
			num, _ := strconv.Atoi(joinedMatches)
			record.Time = num
			continue
		}

		distancesRegex := regexp.MustCompile("^Distance:")
		if distancesRegex.MatchString(line) {
			matches := numberRegex.FindAllString(line, -1)
			joinedMatches := strings.Join(matches, "")
			num, _ := strconv.Atoi(joinedMatches)
			record.Distance = num
			continue
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(record)

	winCount := 0
	for i := 1; i < record.Time; i++ {
		distance := (record.Time - i) * i
		if distance > record.Distance {
			winCount++
		}
	}

	println("(Part 2) number of winning scenarios:", winCount)
}
