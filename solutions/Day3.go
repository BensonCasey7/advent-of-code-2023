package solutions

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type EnginePart struct {
	Number              int
	AffectedCoordinates []string
	Coordinates         []string
	IsCounted           bool
}

func NewEnginePart() EnginePart {
	enginePart := EnginePart{}
	enginePart.Number = 0
	enginePart.AffectedCoordinates = []string{}
	enginePart.Coordinates = []string{}
	enginePart.IsCounted = false
	return enginePart
}

func Day3Part1() {
	file, err := os.Open("solutions/Day3Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var coordinatesToCheck []string
	symbolRegex := regexp.MustCompile("[^0-9.\\n]")
	var engineParts []EnginePart
	partNumberRegex := regexp.MustCompile("[0-9]+")

	x := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		symbolMatchesIndex := symbolRegex.FindAllStringIndex(line, -1)
		for _, matchIndex := range symbolMatchesIndex {
			// Add all surrounding indices to coordinatesToCheck
			coordinatesToCheck = append(coordinatesToCheck, fmt.Sprintf("%d,%d", x-1, matchIndex[0]-1))
			coordinatesToCheck = append(coordinatesToCheck, fmt.Sprintf("%d,%d", x-1, matchIndex[0]))
			coordinatesToCheck = append(coordinatesToCheck, fmt.Sprintf("%d,%d", x-1, matchIndex[0]+1))
			coordinatesToCheck = append(coordinatesToCheck, fmt.Sprintf("%d,%d", x, matchIndex[0]-1))
			coordinatesToCheck = append(coordinatesToCheck, fmt.Sprintf("%d,%d", x, matchIndex[0]+1))
			coordinatesToCheck = append(coordinatesToCheck, fmt.Sprintf("%d,%d", x+1, matchIndex[0]-1))
			coordinatesToCheck = append(coordinatesToCheck, fmt.Sprintf("%d,%d", x+1, matchIndex[0]))
			coordinatesToCheck = append(coordinatesToCheck, fmt.Sprintf("%d,%d", x+1, matchIndex[0]+1))
		}

		partMatchesIndex := partNumberRegex.FindAllStringIndex(line, -1)
		for _, matchIndex := range partMatchesIndex {
			part := NewEnginePart()
			part.Number, _ = strconv.Atoi(line[matchIndex[0]:matchIndex[1]])
			for idx := matchIndex[0]; idx < matchIndex[1]; idx++ {
				part.AffectedCoordinates = append(part.AffectedCoordinates, fmt.Sprintf("%d,%d", x, idx))
			}
			engineParts = append(engineParts, part)
		}

		x++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for _, coordinate := range coordinatesToCheck {
		for i := range engineParts {
			for _, affectedCoordinate := range engineParts[i].AffectedCoordinates {
				if affectedCoordinate == coordinate {
					engineParts[i].IsCounted = true
				}
			}
		}
	}

	partNumberSum := 0
	for _, part := range engineParts {
		if part.IsCounted {
			partNumberSum += part.Number
		}
	}

	println("(Part 1) Sum of all part numbers:", partNumberSum)
}

type Gear struct {
	CoordinatesToCheck []string
	Parts              []EnginePart
}

func NewGear() Gear {
	gear := Gear{}
	gear.CoordinatesToCheck = []string{}
	gear.Parts = []EnginePart{}
	return gear
}

func Day3Part2() {
	file, err := os.Open("solutions/Day3Input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var gears []Gear
	gearRegex := regexp.MustCompile("\\*")
	var engineParts []EnginePart
	partNumberRegex := regexp.MustCompile("[0-9]+")

	x := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		gearMatchesIndex := gearRegex.FindAllStringIndex(line, -1)
		for _, matchIndex := range gearMatchesIndex {
			gear := NewGear()
			gear.CoordinatesToCheck = append(gear.CoordinatesToCheck, fmt.Sprintf("%d,%d", x-1, matchIndex[0]-1))
			gear.CoordinatesToCheck = append(gear.CoordinatesToCheck, fmt.Sprintf("%d,%d", x-1, matchIndex[0]))
			gear.CoordinatesToCheck = append(gear.CoordinatesToCheck, fmt.Sprintf("%d,%d", x-1, matchIndex[0]+1))
			gear.CoordinatesToCheck = append(gear.CoordinatesToCheck, fmt.Sprintf("%d,%d", x, matchIndex[0]-1))
			gear.CoordinatesToCheck = append(gear.CoordinatesToCheck, fmt.Sprintf("%d,%d", x, matchIndex[0]+1))
			gear.CoordinatesToCheck = append(gear.CoordinatesToCheck, fmt.Sprintf("%d,%d", x+1, matchIndex[0]-1))
			gear.CoordinatesToCheck = append(gear.CoordinatesToCheck, fmt.Sprintf("%d,%d", x+1, matchIndex[0]))
			gear.CoordinatesToCheck = append(gear.CoordinatesToCheck, fmt.Sprintf("%d,%d", x+1, matchIndex[0]+1))
			gears = append(gears, gear)
		}

		partMatchesIndex := partNumberRegex.FindAllStringIndex(line, -1)
		for _, matchIndex := range partMatchesIndex {
			part := NewEnginePart()
			part.Number, _ = strconv.Atoi(line[matchIndex[0]:matchIndex[1]])
			for idx := matchIndex[0]; idx < matchIndex[1]; idx++ {
				part.AffectedCoordinates = append(part.AffectedCoordinates, fmt.Sprintf("%d,%d", x, idx))
			}
			engineParts = append(engineParts, part)
		}

		x++
	}

	for i := range gears {
		for _, coordinate := range gears[i].CoordinatesToCheck {
			for j := range engineParts {
				for _, affectedCoordinate := range engineParts[j].AffectedCoordinates {
					if affectedCoordinate == coordinate {
						shouldAdd := true
						for _, part := range gears[i].Parts {
							if part.Number == engineParts[j].Number {
								shouldAdd = false
							}
						}
						if shouldAdd {
							gears[i].Parts = append(gears[i].Parts, engineParts[j])
						}
					}
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	sum := 0
	for _, gear := range gears {
		if len(gear.Parts) == 2 {
			sum += gear.Parts[0].Number * gear.Parts[1].Number
		}
	}

	println("(Part 2) Sum of all gear ratios:", sum)
}
