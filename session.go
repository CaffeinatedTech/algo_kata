package main

import "bufio"
import "fmt"
import "github.com/fatih/color"
import "github.com/spf13/cast"
import "encoding/json"
import "math/rand"
import "os"
import "slices"
import "time"

func randomAlgorithmAndLang(s Session, results []Result) (Algorithm, string) {
	thisAlgorithm := Algorithm{}
	thisLanguage := ""
	for {
		thisAlgorithm = s.Algorithms[rand.Intn(len(s.Algorithms))]
		thisLanguage = s.Languages[rand.Intn(len(s.Languages))]
		// Check if this algorithm, and language are already in the results array.
		contains := slices.ContainsFunc(results, func(r Result) bool {
			return r.Algorithm == thisAlgorithm.Name && r.Language == thisLanguage
		})
		if !contains {
			break
		}
	}
	return thisAlgorithm, thisLanguage
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

// Check if the answer is correct.  the answer is taken from bufio.NewScanner.Scan
func answerCheck(arr []int, expectedResultIndex int, algoType string, answerScanner *bufio.Scanner) bool {
	if algoType == "sort" {
		// Unmarshal answer into a slice of ints
		var answer []int
		json.Unmarshal([]byte(answerScanner.Text()), &answer)
		return compareArrays(arr, answer)
	} else {
		return cast.ToInt(answerScanner.Text()) == expectedResultIndex
	}
}

func runTheSession(s Session) []Result {
	green := color.New(color.FgGreen).Add(color.Bold).SprintFunc()
	results := make([]Result, s.Num)

	for i := 0; i < s.Num; i++ {
		// Get random algorithm, and language - make sure we don't repeat the same pair.
		thisAlgorithm, thisLanguage := randomAlgorithmAndLang(s, results)

		// Generate an array of random numbers
		questionArr := randomArray(10, thisAlgorithm.Sorted)
		// Get the expected result, and the index of the expected result.
		expectedResult, expectedResultIndex := expectedResult(questionArr)

		// Turn the question array into a JSON string for better printing.
		questionJson, _ := json.Marshal(questionArr)
		questionJsonString := string(questionJson)

		// Announce the language of the round
		fmt.Printf("\nThis round will be in %s.  Press Enter when you are ready...\n", green(thisLanguage))
		var input string
		fmt.Scanln(&input)

		// Announce the algorithm of the round, and start a timer.
		fmt.Printf("\nAlgorithm %d: %v\n", i+1, green(thisAlgorithm.Name))
		if thisAlgorithm.Sorted {
			fmt.Println("This array is already sorted.")
		}

		if thisAlgorithm.Type == "search" {
			fmt.Printf("Find the index of %d in the following array:\n", expectedResult)
		}

		fmt.Printf("\nArray: %v\n", questionJsonString)

		start := time.Now()
		fmt.Println("\nEnter your answer when you are done.  Starting timer...")
		// If this is a sorting algorithm, then sort the array, ready to compare to the user's input later.
		if thisAlgorithm.Type == "sort" {
			shellSort(questionArr)
			questionJson, _ = json.Marshal(questionArr)
			questionJsonString = string(questionJson)
		}

		answerScanner := bufio.NewScanner(os.Stdin)
		answerScanner.Scan()
		end := time.Now()
		fmt.Println("Time taken: ", end.Sub(start).Round(time.Second))

		answerCorrect := answerCheck(questionArr, expectedResultIndex, thisAlgorithm.Type, answerScanner)

		if answerCorrect {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect!")
		}

		fmt.Print("\n")
		// Save the results for this round
		results[i] = Result{Algorithm: thisAlgorithm.Name, Language: thisLanguage, Time: end.Sub(start).Round(time.Second), Correct: answerCorrect}
	}
	return results
}
