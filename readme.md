# Club Anagram Problem

## Problem Statement
You are given a slice of strings.

    input = []string{string1, string2, string3, ...}

Club all the `anagrams` together in slice of slice of strings.

## Constraints

- Handle edge cases of input strings.

## Example

    input = []string{"fed", "bca", "def", "cab", "xyz"}

Slice1 = []string{"bca", "cab"} \
Slice2 = []string{"def", "fed"} \
Slice3 = []string{"xyz"}

    output = [][]string{{"bca", "cab"},{"def", "fed"},{"xyz"}}

## Solution

```golang
func ListAllAnagramsInSlice(wordList []string) [][]string
```
Function [ListAllAnagramsInSlice](https://github.com/VILJkid/go-club-anagrams/blob/1c1118d44b9d84b6e8c7931c68f52c7c7ec49a12/program.go#L86-L89) solves the given problem using hashmap and iterating over each string and updating the map. \
Finally, it sorts the map and returns the result as a slice of slice of strings.

## Execution

Run the tests already provided.

```shell
go test -v
```