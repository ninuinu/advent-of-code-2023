package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CardNumbers struct{
	observerdNumbers []string
	winningNumbers []string
}

// Removes any empty strings that may exist in the input file
func deleteEmptyStrings(strarr []string) []string {
   var arr []string
   for _, str := range strarr {
      if str != "" {
         arr = append(arr, str)
      }
   }
   return arr
}

// Populates the cardMap map with information about the cards, read from the input file
func addCardToCardMap(line string, cardMap map[int]CardNumbers) map[int]CardNumbers{
	cardIdAndNumbers := strings.Split(line, ":")
	cardId := strings.Trim(cardIdAndNumbers[0], "Card ")
	numbers := strings.Split(cardIdAndNumbers[1], "|")

	observedNumbers := deleteEmptyStrings(strings.Split(string(numbers[0]), " "))
	winningNumbers := deleteEmptyStrings(strings.Split(string(numbers [1]), " "))

	cardNumbers := CardNumbers{observerdNumbers: observedNumbers, winningNumbers: winningNumbers}
	
	id, _ := strconv.Atoi(cardId)

	cardMap[id] = cardNumbers

	return cardMap
}

// Puzzle 2: Adds the occurances of card originals to the cardCopies map
func populateCardCopiesWithOriginalInstance(cardCopies map[int]int, cardId int) map[int]int{
	cardCopies[cardId] = 1
	return cardCopies
}

// Puzzle 2: Adds the occurances of card copies to the cardCopies map
func addCopies(k int, matches int, cardCopies map[int]int) map[int]int{
	
	iterations := cardCopies[k]

	for i:= 0; i < iterations; i++ {
		if matches == 1 {
			cardCopies[k+1] += 1
		} else {
			for j := 1; j <= matches; j++ {
				cardCopies[j+k] += 1
		
			}
		}
	}
	return cardCopies
}


func main() {
	file, _ := os.Open("input.txt")
	fscanner := bufio.NewScanner(file)
	
	cardMap := make(map[int]CardNumbers)
	cardCopies := make(map[int]int)

	cardId := 1

	for fscanner.Scan(){
		line := fscanner.Text()
		cardCopies = populateCardCopiesWithOriginalInstance(cardCopies, cardId)
		cardMap = addCardToCardMap(line, cardMap)
		cardId +=1
	}

	totalPointsPartOne := 0
	totalPointsPartTwo := 0

	// Extract keys
    keys := make([]int, 0, len(cardMap))
    for k := range cardMap {
        keys = append(keys, k)
    }

    // Sort the keys slice
    sort.Ints(keys)

	for _, key := range keys{
		matches := 0

		totalPointsForThisCard := 0
		for _, observedNumber := range cardMap[key].observerdNumbers{
			for _, winningNumber := range cardMap[key].winningNumbers{
				if observedNumber == winningNumber {
					matches += 1
					if matches == 1 {
						totalPointsForThisCard += 1
					} else if matches > 1{
						totalPointsForThisCard = totalPointsForThisCard * 2
					}
				}

			}			
		}

		totalPointsPartOne += totalPointsForThisCard
		cardCopies = addCopies(key, matches, cardCopies)
	}

	for _, occurances := range cardCopies{
		totalPointsPartTwo += occurances
	}

	fmt.Printf("Solution to puzzle 1: %d\n", totalPointsPartOne)	
	fmt.Printf("Solution to puzzle 2: %d\n", totalPointsPartTwo)	
}
	


