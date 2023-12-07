package day5

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type MapDefinition struct {
	Start      int
	End        int
	Difference int
}

func Part1() {
	file, err := os.Open("day5/Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	allMaps := make([][]MapDefinition, 7)
	var seeds []int

	mapIndex := -1
	scanner := bufio.NewScanner(file)
	var tempMaps []MapDefinition
	for scanner.Scan() {
		line := scanner.Text()

		seedsRegex := regexp.MustCompile("^seeds: ")
		if seedsRegex.MatchString(line) {
			seedsString := seedsRegex.ReplaceAllString(line, "")
			seedsStringSlice := strings.Split(seedsString, " ")
			for _, seedString := range seedsStringSlice {
				seed, _ := strconv.Atoi(seedString)
				seeds = append(seeds, seed)
			}
			continue
		}

		mapLabelRegex := regexp.MustCompile("^.*-.*-.* map:")
		if mapLabelRegex.MatchString(line) {
			if mapIndex >= 0 {
				allMaps[mapIndex] = tempMaps
				tempMaps = []MapDefinition{}
			}
			mapIndex++
			continue
		}

		numberRegex := regexp.MustCompile("[0-9]+")
		if numberRegex.MatchString(line) {
			nums := strings.Split(line, " ")
			destinationStart := nums[0]
			sourceStart := nums[1]
			rangeLength := nums[2]

			destinationStartInt, _ := strconv.Atoi(destinationStart)
			sourceStartInt, _ := strconv.Atoi(sourceStart)
			rangeLengthInt, _ := strconv.Atoi(rangeLength)

			tempMaps = append(tempMaps, MapDefinition{
				Start:      sourceStartInt,
				End:        sourceStartInt + rangeLengthInt - 1,
				Difference: destinationStartInt - sourceStartInt,
			})
		}
	}
	allMaps[mapIndex] = tempMaps

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	minLocation := math.MaxUint32
	for _, seed := range seeds {
		location := seed
		for _, maps := range allMaps {
			for _, m := range maps {
				if location >= m.Start && location <= m.End {
					location += m.Difference
					break
				}
			}
		}
		if location < minLocation {
			minLocation = location
		}

		//println("Seed:", seed, "Location:", location)
	}

	println("(Part 1) Minimum location:", minLocation)
}
func Part2() {
	file, err := os.Open("day5/Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	allMaps := make([][]MapDefinition, 7)
	var seeds []int

	mapIndex := -1
	scanner := bufio.NewScanner(file)
	var tempMaps []MapDefinition
	for scanner.Scan() {
		line := scanner.Text()

		seedsRegex := regexp.MustCompile("^seeds: ")
		if seedsRegex.MatchString(line) {
			seedsString := seedsRegex.ReplaceAllString(line, "")
			seedsStringSlice := strings.Split(seedsString, " ")
			for i, seedString := range seedsStringSlice {
				if i%2 == 0 {
					startSeed, _ := strconv.Atoi(seedString)
					rangeLength, _ := strconv.Atoi(seedsStringSlice[i+1])
					for j := 0; j < rangeLength; j++ {
						seed := startSeed + j
						seeds = append(seeds, seed)
					}
				}
			}
			continue
		}

		mapLabelRegex := regexp.MustCompile("^.*-.*-.* map:")
		if mapLabelRegex.MatchString(line) {
			if mapIndex >= 0 {
				allMaps[mapIndex] = tempMaps
				tempMaps = []MapDefinition{}
			}
			mapIndex++
			continue
		}

		numberRegex := regexp.MustCompile("[0-9]+")
		if numberRegex.MatchString(line) {
			nums := strings.Split(line, " ")
			destinationStart := nums[0]
			sourceStart := nums[1]
			rangeLength := nums[2]

			destinationStartInt, _ := strconv.Atoi(destinationStart)
			sourceStartInt, _ := strconv.Atoi(sourceStart)
			rangeLengthInt, _ := strconv.Atoi(rangeLength)

			tempMaps = append(tempMaps, MapDefinition{
				Start:      sourceStartInt,
				End:        sourceStartInt + rangeLengthInt - 1,
				Difference: destinationStartInt - sourceStartInt,
			})
		}
	}
	allMaps[mapIndex] = tempMaps

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	minLocation := math.MaxUint32
	for _, seed := range seeds {
		location := seed
		for _, maps := range allMaps {
			for _, m := range maps {
				if location >= m.Start && location <= m.End {
					location += m.Difference
					break
				}
			}
		}
		if location < minLocation {
			minLocation = location
		}

		//println("Seed:", seed, "Location:", location)
	}

	println("(Part 2) Minimum location:", minLocation)
}
