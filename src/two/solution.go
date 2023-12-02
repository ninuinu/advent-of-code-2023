package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, _ := os.Open("input.txt")
	fscanner := bufio.NewScanner(file)

	resultFirstPuzzle := 0
	resultSecondPuzzle := 0
	
	for fscanner.Scan(){
		line := fscanner.Text()
		s := strings.Split(line, ":")
		
		gameComponent := s[0]
		gameId := strings.Trim(gameComponent, "Game :")
		observations := strings.Split(s[1], ";")

		fulfillsCondition := true

		redMax := 0
		greenMax := 0
		blueMax := 0

		product := 0


		for _, observation := range observations {
			cubes := strings.Split(observation, ",")
			for _, cube := range cubes{
				digitAndColour := strings.Split(cube, " ")

				digitAsString := digitAndColour[1]
				digit, _ := strconv.Atoi(digitAsString)
				colour := digitAndColour[2]

				if colour == "red"{
					if digit > 12 {
						fulfillsCondition = false
					} 
					if digit > redMax{
						redMax = digit
					}
				} else if colour == "green"{
					if digit > 13 {
						fulfillsCondition = false
					}
					if digit > greenMax{
						greenMax = digit
					}
				} else if colour == "blue" {
					if digit > 14 {
						fulfillsCondition = false
					}
					if digit > blueMax{
						blueMax = digit
					}
				}
								
			}

		}

		product = redMax * blueMax * greenMax
		resultSecondPuzzle += product


		if fulfillsCondition{
			if id, err := strconv.Atoi(gameId); err == nil {
				resultFirstPuzzle += id
			}
		}

	}
}
