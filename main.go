package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	inputList := []string{"cerg ", "sample", " cegr", "", " ", "🚲ab", "ba🚲", "cgre", "cger", "samlep", "samepl", "samelp", "sapmle", "anurag", "gnuraA"}

	fmt.Printf("Input:\n%#+v\n", inputList)
	fmt.Printf("Output:\n%#+v\n", ListAllAnagramsInSlice(inputList))
}

// Input a slice of words and return the list of words, bunched together
// if they are made up of same characters of string
func ListAllAnagramsInSlice(wordList []string) (anagramList [][]string) {

	anagramMap := ListAllAnagramsInMap(wordList)
	anagramList = convertMapToSlice(anagramMap)

	return
}

// List anagrams in map
func ListAllAnagramsInMap(wordList []string) map[string][]string {
	mapOfSliceOfString := make(map[string][]string)

	for _, word := range wordList {
		word = filterAndValidateWord(word)
		sortedWord := sortString(word)

		// Check if the word-key already exists in map
		if _, found := mapOfSliceOfString[sortedWord]; found {
			mapOfSliceOfString[sortedWord] = append(mapOfSliceOfString[sortedWord], word)
			continue
		}

		// Instantiate a slice with word
		newSlice := []string{word}
		mapOfSliceOfString[sortedWord] = newSlice
	}
	return mapOfSliceOfString
}

// Clean the input (edge cases)
func filterAndValidateWord(inputString string) string {
	inputString = strings.TrimSpace(inputString)
	inputString = strings.ToLower(inputString)
	return inputString
}

// Sort the inputString
func sortString(inputString string) string {
	splittedString := strings.Split(inputString, "")
	sort.Strings(splittedString)
	sortedString := strings.Join(splittedString, "")
	return sortedString
}

// Convert map[string] to slice
func convertMapToSlice(inputMap map[string][]string) (outputSlice [][]string) {
	for _, eachSlice := range inputMap {
		outputSlice = append(outputSlice, eachSlice)
	}
	return
}
