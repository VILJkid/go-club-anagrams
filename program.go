/*
	Problem statement:
	You are given a slice of strings.
	input = []string{string1, string2, string3, ...}

	Sort and club the strings together if they are anagrams.

	Eg:
		input = []string{"fed", "bca", "def", "cab", "xyz"}

		output = [][]string{
			{"bca", "cab"},
			{"def", "fed"},
			{"xyz"},
		}
*/

package main

import (
	"sort"
	"strings"

	"slices"
)

// Clean the input (edge cases)
func filterAndValidateWord(inputString string) string {
	inputString = strings.TrimSpace(inputString)
	inputString = strings.ToLower(inputString)
	return inputString
}

// Sort the inputString
func sortString(inputString string) string {
	inputBytes := []byte(inputString)
	sort.SliceStable(inputBytes, func(i, j int) bool {
		return inputBytes[i] < inputBytes[j]
	})
	return string(inputBytes)
}

// List anagrams in map
func listAllAnagramsInMap(wordList []string) map[string][]string {
	mapOfSliceOfString := make(map[string][]string, len(wordList))

	for _, word := range wordList {
		word = filterAndValidateWord(word)
		sortedWord := sortString(word)

		// Check if the word-key already exists in map
		if existingSlice, found := mapOfSliceOfString[sortedWord]; found {
			mapOfSliceOfString[sortedWord] = append(existingSlice, word)
			slices.Sort(mapOfSliceOfString[sortedWord])
			continue
		}

		// Instantiate a slice with word
		mapOfSliceOfString[sortedWord] = []string{word}
	}
	return mapOfSliceOfString
}

// Sort map and return the sorted keys
func getsortedMapKeys(inputMap map[string][]string) []string {
	keys := make([]string, 0, len(inputMap))
	for key := range inputMap {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}

// Convert map[string] to slice
func convertMapToSlice(inputMap map[string][]string) [][]string {
	keys := getsortedMapKeys(inputMap)
	outputSlice := make([][]string, 0, len(keys))
	for _, key := range keys {
		outputSlice = append(outputSlice, inputMap[key])
	}
	return outputSlice
}

// Input a slice of words and return the list of words, bunched together
// if they are made up of same characters of string
func ListAllAnagramsInSlice(wordList []string) [][]string {
	anagramMap := listAllAnagramsInMap(wordList)
	return convertMapToSlice(anagramMap)
}
