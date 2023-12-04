package main

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

func checkNeighboursForFirstAndLastLine(i int, j int, m [][]int, coordinates []Coordinate) []Coordinate {
	// corner cases on first row
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

	return coordinates

}

func checkNeighboursForLeftAndRightSides(i int, j int, m [][]int, coordinates []Coordinate) []Coordinate {
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
	return coordinates
}

func checkNeighbours(i int, j int, m [][]int, coordinates []Coordinate) []Coordinate {
	coordinates = checkNorth(i, j, m, coordinates)
	coordinates = checkNorthEast(i, j, m, coordinates)
	coordinates = checkEast(i, j, m, coordinates)
	coordinates = checkSouthEast(i, j, m, coordinates)
	coordinates = checkSouth(i, j, m, coordinates)
	coordinates = checkSouthWest(i, j, m, coordinates)
	coordinates = checkWest(i, j, m, coordinates)
	coordinates = checkNorthWest(i, j, m, coordinates)
	return coordinates
}
