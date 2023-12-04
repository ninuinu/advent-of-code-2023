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

type NumberAndCoordinate struct {
	startCoordinate Coordinate
	endCoordinate   Coordinate
	number          int
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

func setDigitFlag(i int, j int, m [][]int) [][]int {
	if i < 0 || i >= len(m) || j < 0 || j >= len(m[i]) {
		// if out of bounds
		return m
	}
	m[i][j] = 1
	return m
}

func setMatrixFlag(i int, j int, m [][]int) [][]int {
	m[i][j] = 2
	return m
}

func isColumnOutOfBounds(columnIndex int, line string) bool{
	return columnIndex == len(line)-1

}

func extractNumbers(i int, line string, numbers []NumberAndCoordinate) []NumberAndCoordinate {
	numberAsString := ""
	for j, character := range line {
		if unicode.IsDigit(character) {
			numberAsString += string(character)
			if isColumnOutOfBounds(j, line) {
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

func getDigitsFromCoordinates(coordinates []Coordinate, numbers []NumberAndCoordinate) []Coordinate {
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

	return finalTuples
}

func isFirstOrLastRow(i int, lastIndex int) bool{
	return (i == 0) || (i == lastIndex)
}

func isLeftOrRightBorder(i int, j int, lastIndex int) bool {
	return (i != 0 || i != lastIndex) && (j == 0 || j == lastIndex)
}

func isSymbol(c string) bool{
 return (c == "*") || (c == "$") || (c == "-") || (c == "+") || (c == "#") || (c == "=") || (c == "/") || (c == "@") || (c == "%") || (c == "&") 
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

		numbers = extractNumbers(lineIndex, line, numbers)

		for index, c := range line {

			char := string(c)

			if unicode.IsDigit(c) {
				m = setDigitFlag(lineIndex, index, m)
			}
			if isSymbol(char){
				m = setMatrixFlag(lineIndex, index, m)
				totalNumberOfSymbols += 1
			}
		}
		lineIndex += 1
	}

	lastIndex := len(m) - 1

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			// if char stored at index is a symbol
			if m[i][j] == 2 {
				if isFirstOrLastRow(i, lastIndex) {
					coordinates = checkNeighboursForFirstAndLastLine(i, j, m, coordinates)
				} else if isLeftOrRightBorder(i, j, lastIndex) { 
					coordinates = checkNeighboursForLeftAndRightSides(i, j, m, coordinates)
				} else { 
					// all cells that do not lie on a border
					coordinates = checkNeighbours(i, j, m, coordinates)
				}
			}
		}
	}

	coordinatesForDigitsToBeIncludedInFinalSum := getDigitsFromCoordinates(coordinates, numbers)
	sum := 0

	for _, coordinate := range coordinatesForDigitsToBeIncludedInFinalSum {
		for _, number := range numbers {
			if number.endCoordinate.X == coordinate.X && number.endCoordinate.Y == coordinate.Y {
				sum += number.number
			}
		}
	}

	fmt.Printf("Sum for puzzle 1 is: %d\n", sum)
}
