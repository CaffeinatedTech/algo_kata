package main

import "encoding/json"
import "math/rand"
import "slices"
import "github.com/spf13/cast"

func (m model) randomAlgorithmAndLang() (Algorithm, string) {

	thisAlgorithm := Algorithm{}
	var thisLanguage Language
	for {
		thisAlgorithm = m.session.Algorithms[rand.Intn(len(m.session.Algorithms))]
		thisLanguage = m.session.Languages[rand.Intn(len(m.session.Languages))]
		// Check if this algorithm, and language are already in the results array.
		contains := slices.ContainsFunc(m.results, func(r Result) bool {
			return r.Algorithm == thisAlgorithm.Name && r.Language == thisLanguage.Name
		})
		if !contains && thisAlgorithm.Selected && thisLanguage.Selected {
			break
		}
	}
	return thisAlgorithm, thisLanguage.Name
}

func randomArray(num int, sorted bool) []int {
	var arr []int
	for i := 0; i < num; i++ {
		arr = append(arr, rand.Intn(100))
	}
	if sorted {
		shellSort(arr)
	}
	return arr
}

func expectedResult(arr []int) (int, int) {
	expected := arr[rand.Intn(len(arr))]
	return expected, slices.Index(arr, expected)
}

// Check if the answer is correct.
func answerCheck(arr []int, expectedResultIndex int, algoType string, answerString string) bool {
	if algoType == "sort" {
		shellSort(arr)
		// Convert the answerString to an array of integers
		var answer []int
		err := json.Unmarshal([]byte(answerString), &answer)
		if err != nil {
			return false
		}
		return compareArrays(arr, answer)
	} else {
		return cast.ToInt(answerString) == expectedResultIndex
	}
}
