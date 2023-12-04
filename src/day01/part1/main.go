package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func main() {
	file, _ := os.Open("../input.txt")
	fscanner := bufio.NewScanner(file)
	
	var totalSum int = 0
	for fscanner.Scan(){
		line := fscanner.Text()

		var left string
		var right string

		
		for _, character := range line{
			if unicode.IsDigit(character){
				left = string(character)
				break
			}
			fmt.Print(left)
		}

		runes := []rune(line) 
		for i := len(runes) - 1; i >= 0; i-- {
			if unicode.IsDigit(runes[i]){
				right = string(runes[i])
				break
			}
			fmt.Print(right)
		}

		num := left + right
		if n, err := strconv.Atoi(num); err == nil {
			totalSum += n
		}
	}
	fmt.Printf("%d", totalSum)
}

