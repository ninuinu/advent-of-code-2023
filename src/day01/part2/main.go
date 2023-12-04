package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)


func main() {
	file, _ := os.Open("../input.txt")
	fscanner := bufio.NewScanner(file)

	numbers := [10]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	
	var totalSum int = 0
	for fscanner.Scan(){
		line := fscanner.Text()

		var left string
		var right string

		leftAlreadySet := false
		for i, character := range line{
			if leftAlreadySet {
				break
			}
			if unicode.IsDigit(character){
				left = string(character)
				break
			} else {

				for j, number := range numbers {
					end := i+len(number)
					lineIndexRange := len(line) 
					
					if end > lineIndexRange {
						end = lineIndexRange
					}
					potentialMatch := line[i:end]
					if(number == strings.TrimRight(potentialMatch, "\n") && j+1 != 10){
						left = strconv.Itoa(j+1)
						leftAlreadySet = true
						break
					}	
				}
			}
		}

		rightAlreadySet := false
		runes := []rune(line) 
		for i := len(runes) - 1; i >= 0; i-- {
			if rightAlreadySet {
				break
			}
			if unicode.IsDigit(runes[i]){
				right = string(runes[i])
				break
			} else {
				for j, number := range numbers {

					start := i-(len(number)-1)
					if start < 0 {
						start = 0
					}

					potentialMatch := line[start:i+1]

					if(number == strings.TrimRight(potentialMatch, "\n") && j+1 != 10){

						right = strconv.Itoa(j+1)
						rightAlreadySet = true
						break
					}	
				}
			}
		}

		numAsString := left + right
		if numAsDigit, err := strconv.Atoi(numAsString); err == nil {
			totalSum += numAsDigit
		}
	}
	fmt.Printf("%d", totalSum)
}
