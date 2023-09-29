package main

import (
	"reflect"
	"strconv"
	"testing"
)

type testInputList struct {
	inputList    []string
	expectedList [][]string
}

var testInputLists = []testInputList{
	{
		inputList: []string{"cerg ", "sample", " cegr", "", " ", "🚲ab", "ba🚲", "cgre", "cger", "samlep", "samepl", "samelp", "sapmle", "anurag", "gnuraA"},
		expectedList: [][]string{
			{"", ""},
			{"anurag", "gnuraa"},
			{"ba🚲", "🚲ab"},
			{"samelp", "samepl", "samlep", "sample", "sapmle"},
			{"cegr", "cerg", "cger", "cgre"},
		},
	},
}

func TestListAllAnagramsInSlice(t *testing.T) {
	for index, test := range testInputLists {
		testName := "Test" + strconv.Itoa(index)
		t.Run(testName, func(t *testing.T) {
			if result := ListAllAnagramsInSlice(test.inputList); !reflect.DeepEqual(result, test.expectedList) {
				t.Errorf("expected:\n\t%#+v\ngot\n\t%#+v", test.expectedList, result)
			}
		})
	}
}

func BenchmarkListAllAnagramsInSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ListAllAnagramsInSlice(testInputLists[0].inputList)
	}
}
