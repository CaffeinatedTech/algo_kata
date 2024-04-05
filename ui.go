package main

import "fmt"

func welcomMessage(c []Algorithm, languages []string) {
	languages_text, algorithms_text := configText(c, languages)
	fmt.Println("--== Algorithm Kata ==--")
	fmt.Println("Lets practice some algorithms.")
	fmt.Println("\nLanguages: ", languages_text)
	fmt.Println("Algorithms: ", algorithms_text)
}

func askForNumber() int {
	fmt.Println("\nHow many would you like to do this session?")
	var num int
	_, err := fmt.Scan(&num)
	if err != nil {
		fmt.Println("Error reading number of algorithms")
		panic(1)
	}
	return num
}

func printResults(results []Result) {
	fmt.Println("Results:")
	for i, result := range results {
		thisResultCorrect := "Incorrect"
		if result.Correct {
			thisResultCorrect = "Correct"
		}
		fmt.Printf("Algorithm %d: %s in %s %s %s\n", i+1, result.Algorithm, result.Language, result.Time, thisResultCorrect)
	}
}
