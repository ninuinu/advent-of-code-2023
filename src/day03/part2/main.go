package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode"
)

type Coordinate struct {
	X, Y int
}

type GearAndCoordinate struct {
	coordinate Coordinate
	matches    []int
}

var gearMatches []GearAndCoordinate

type NumberAndCoordinate struct {
	startCoordinate Coordinate
	endCoordinate   Coordinate
	number          int
}

func checkWest(i int, j int, m [][]int, coordinates []Coordinate) []Coordinate {
	if m[i][j-1] == 1 {
		coordinates = append(coordinates, Coordinate{X: i, Y: j - 1})
	}
	return coordinates
}

func checkEast(i int, j int, m [][]int, coordinates []Coordinate) []Coordinate {
	if m[i][j+1] == 1 {
		coordinates = append(coordinates, Coordinate{X: i, Y: j + 1})
	}
	return coordinates
}

func checkSouth(i int, j int, m [][]int, coordinates []Coordinate) []Coordinate {
	if m[i+1][j] == 1 {
		coordinates = append(coordinates, Coordinate{X: i + 1, Y: j})
	}
	return coordinates
}

func checkSouthWest(i int, j int, m [][]int, coordinates []Coordinate) []Coordinate {
	if m[i+1][j-1] == 1 {
		coordinates = append(coordinates, Coordinate{X: i + 1, Y: j - 1})
	}
	return coordinates
}

func checkSouthEast(i int, j int, m [][]int, coordinates []Coordinate) []Coordinate {
	if m[i+1][j+1] == 1 {
		coordinates = append(coordinates, Coordinate{X: i + 1, Y: j + 1})
	}
	return coordinates
}

func checkNorth(i int, j int, m [][]int, coordinates []Coordinate) []Coordinate {
	if m[i-1][j] == 1 {
		coordinates = append(coordinates, Coordinate{X: i - 1, Y: j})
	}
	return coordinates
}

func checkNorthWest(i int, j int, m [][]int, coordinates []Coordinate) []Coordinate {
	if m[i-1][j-1] == 1 {
		coordinates = append(coordinates, Coordinate{X: i - 1, Y: j - 1})
	}
	return coordinates
}

func checkNorthEast(i int, j int, m [][]int, coordinates []Coordinate) []Coordinate {
	if m[i-1][j+1] == 1 {
		coordinates = append(coordinates, Coordinate{X: i - 1, Y: j + 1})
	}
	return coordinates
}

func checkNeighboursForFirstAndLastLine(i int, j int, m [][]int, coordinates []Coordinate, numbers []NumberAndCoordinate) []Coordinate {
	// corner cases on first row
	coordinatesBefore := coordinates

	lastIndex := len(m) - 1
	if i == 0 && (j == 0 || j == lastIndex) {
		if j == 0 {
			coordinates = checkEast(i, j, m, coordinates)
			coordinates = checkSouthEast(i, j, m, coordinates)

		} else if j == lastIndex {
			coordinates = checkWest(i, j, m, coordinates)
			coordinates = checkSouthWest(i, j, m, coordinates)
		}
		// both first and last element in row want to check value beneath them
		coordinates = checkSouth(i, j, m, coordinates)

	} else if i == 0 { // all other cases on first row
		coordinates = checkWest(i, j, m, coordinates)
		coordinates = checkEast(i, j, m, coordinates)
		coordinates = checkSouthWest(i, j, m, coordinates)
		coordinates = checkSouth(i, j, m, coordinates)
		coordinates = checkSouthEast(i, j, m, coordinates)
	}

	if i == lastIndex && (j == 0 || j == lastIndex) {
		if j == 0 {
			coordinates = checkEast(i, j, m, coordinates)
			coordinates = checkNorthEast(i, j, m, coordinates)

		} else if j == lastIndex {
			coordinates = checkWest(i, j, m, coordinates)
			coordinates = checkNorthWest(i, j, m, coordinates)
		}
		// both first and last element in row want to check value beneath them
		coordinates = checkNorth(i, j, m, coordinates)

	} else if i == lastIndex { // all other cases on first row
		coordinates = checkWest(i, j, m, coordinates)
		coordinates = checkEast(i, j, m, coordinates)
		coordinates = checkNorthWest(i, j, m, coordinates)
		coordinates = checkNorth(i, j, m, coordinates)
		coordinates = checkNorthEast(i, j, m, coordinates)
	}

	var digits []int
	// coordinates

	unique := getUniqueSet(coordinatesBefore, coordinates)
	digits = getDigitsFromCoordinates(unique, numbers, digits)

	if len(digits) == 2 {
		coordinate := Coordinate{X: i, Y: j}
		gearMatches = append(gearMatches, GearAndCoordinate{coordinate: coordinate, matches: digits})
	}
	return coordinates

}

func checkNeighboursForLeftAndRightSides(i int, j int, m [][]int, coordinates []Coordinate, numbers []NumberAndCoordinate) []Coordinate {

	coordinatesBefore := coordinates

	if j == 0 {
		coordinates = checkNorthEast(i, j, m, coordinates)
		coordinates = checkEast(i, j, m, coordinates)
		coordinates = checkSouthEast(i, j, m, coordinates)
	} else {
		coordinates = checkNorthWest(i, j, m, coordinates)
		coordinates = checkWest(i, j, m, coordinates)
		coordinates = checkSouthWest(i, j, m, coordinates)
	}
	coordinates = checkNorth(i, j, m, coordinates)
	coordinates = checkSouth(i, j, m, coordinates)

	var digits []int

	unique := getUniqueSet(coordinatesBefore, coordinates)
	digits = getDigitsFromCoordinates(unique, numbers, digits)

	if len(digits) == 2 {
		coordinate := Coordinate{X: i, Y: j}
		gearMatches = append(gearMatches, GearAndCoordinate{coordinate: coordinate, matches: digits})
	}

	return coordinates
}

func getUniqueSet(a []Coordinate, b []Coordinate) []Coordinate {

	tupleCounts := make(map[Coordinate]int)

	// Count tuples in both slices
	for _, tuple := range a {
		tupleCounts[tuple]++
	}
	for _, tuple := range b {
		tupleCounts[tuple]++
	}

	// Find unique tuples
	var uniqueTuples []Coordinate
	for tuple, count := range tupleCounts {
		if count == 1 { // Tuple is unique to one of the slices
			uniqueTuples = append(uniqueTuples, tuple)
		}
	}
	return uniqueTuples
}

func checkNeighbours(i int, j int, m [][]int, coordinates []Coordinate, numbers []NumberAndCoordinate) []Coordinate {

	coordinatesBefore := coordinates

	coordinates = checkNorth(i, j, m, coordinates)
	coordinates = checkNorthEast(i, j, m, coordinates)
	coordinates = checkEast(i, j, m, coordinates)
	coordinates = checkSouthEast(i, j, m, coordinates)
	coordinates = checkSouth(i, j, m, coordinates)
	coordinates = checkSouthWest(i, j, m, coordinates)
	coordinates = checkWest(i, j, m, coordinates)
	coordinates = checkNorthWest(i, j, m, coordinates)

	var digits []int
	// coordinates

	unique := getUniqueSet(coordinatesBefore, coordinates)
	digits = getDigitsFromCoordinates(unique, numbers, digits)

	if len(digits) == 2 {
		coordinate := Coordinate{X: i, Y: j}
		gearMatches = append(gearMatches, GearAndCoordinate{coordinate: coordinate, matches: digits})
	}

	return coordinates
}

func initMatrix(size int) [][]int {
	m := make([][]int, size)
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			m[i] = append(m[i], 0)
		}
	}
	return m
}

func setMatrixCellToDigit(i int, j int, m [][]int) [][]int {
	if i < 0 || i >= len(m) || j < 0 || j >= len(m[i]) {
		// if out of bounds
		return m
	}
	m[i][j] = 1
	return m
}

func setMatrixCellToSymbol(i int, j int, m [][]int) [][]int {
	m[i][j] = 2
	return m
}

func getNumbers(i int, line string, numbers []NumberAndCoordinate) []NumberAndCoordinate {
	numberAsString := ""
	for j, c := range line {
		if unicode.IsDigit(c) {
			numberAsString += string(c)
			if j == len(line)-1 {

				lengthOfNumber := len(numberAsString)
				startCoordinate := Coordinate{X: i, Y: j - lengthOfNumber + 1}

				endCoordinate := Coordinate{X: i, Y: j}
				number, _ := strconv.Atoi(numberAsString)
				numbers = append(numbers, NumberAndCoordinate{startCoordinate: startCoordinate, endCoordinate: endCoordinate, number: number})
			}
		} else {
			if numberAsString != "" {
				lengthOfNumber := len(numberAsString)
				startCoordinate := Coordinate{X: i, Y: j - lengthOfNumber}
				endCoordinate := Coordinate{X: i, Y: j - 1}
				number, _ := strconv.Atoi(numberAsString)
				numbers = append(numbers, NumberAndCoordinate{startCoordinate: startCoordinate, endCoordinate: endCoordinate, number: number})
				numberAsString = ""

			}
		}
	}
	return numbers
}

func getDigitsFromCoordinates(coordinates []Coordinate, numbers []NumberAndCoordinate, digits []int) []int {
	sort.Slice(coordinates, func(i, j int) bool {
		if coordinates[i].X == coordinates[j].X {
			return coordinates[i].Y < coordinates[j].Y
		}
		return coordinates[i].X < coordinates[j].X
	})

	var uniqueCoordinates []Coordinate

	for i, coordinate := range coordinates {
		if i == len(coordinates)-1 {
			uniqueCoordinates = append(uniqueCoordinates, coordinate)
		} else {

			nextCoordinate := coordinates[i+1]

			if coordinate.X == nextCoordinate.X {
				if coordinate.Y+1 != nextCoordinate.Y {
					uniqueCoordinates = append(uniqueCoordinates, coordinate)
				}
			} else {
				uniqueCoordinates = append(uniqueCoordinates, coordinate)
			}
		}
	}

	var uniqueTuples []Coordinate
	for i, coordinate := range uniqueCoordinates {
		if i == 0 || coordinate != uniqueCoordinates[i-1] {
			uniqueTuples = append(uniqueTuples, coordinate)
		}
	}

	var finalTuples []Coordinate

	for _, tuple := range uniqueTuples {
		hasBeenAdded := false
		for _, number := range numbers {
			if tuple.X == number.startCoordinate.X {
				if tuple.Y >= number.startCoordinate.Y && tuple.Y < number.endCoordinate.Y {
					finalTuples = append(finalTuples, number.endCoordinate)
					hasBeenAdded = true
				}
			}
		}
		if !hasBeenAdded {
			finalTuples = append(finalTuples, tuple)
		}
	}

	for _, coordinate := range finalTuples {
		for _, number := range numbers {
			if number.endCoordinate.X == coordinate.X && number.endCoordinate.Y == coordinate.Y {
				digits = append(digits, number.number)
			}
		}
	}

	return digits

}

func main() {

	file, _ := os.Open("../input.txt")
	fscanner := bufio.NewScanner(file)
	totalNumberOfSymbols := 0
	lineIndex := 0

	size := 140

	m := initMatrix(size)

	var coordinates []Coordinate
	var numbers []NumberAndCoordinate

	for fscanner.Scan() {
		line := fscanner.Text()

		numbers = getNumbers(lineIndex, line, numbers)

		for index, c := range line {

			char := string(c)

			if unicode.IsDigit(c) {
				m = setMatrixCellToDigit(lineIndex, index, m)
			}
			if char == "*" {
				m = setMatrixCellToSymbol(lineIndex, index, m)
				totalNumberOfSymbols += 1
			}
		}

		lineIndex += 1

	}

	lastIndex := len(m) - 1

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if m[i][j] == 2 {
				if (i == 0) || (i == lastIndex) {
					coordinates = checkNeighboursForFirstAndLastLine(i, j, m, coordinates, numbers)
				} else if (i != 0 || i != lastIndex) && (j == 0 || j == lastIndex) {
					coordinates = checkNeighboursForLeftAndRightSides(i, j, m, coordinates, numbers)
				} else {
					coordinates = checkNeighbours(i, j, m, coordinates, numbers)
				}
			}
		}
	}

	sumOfGearRatios := 0
	for _, match := range gearMatches {
		gearRatio := 1
		for _, number := range match.matches {
			gearRatio = gearRatio * number
		}
		sumOfGearRatios += gearRatio
	}

	fmt.Printf("Sum of gear ratios for puzzle 2 is: %d\n", sumOfGearRatios)

}
