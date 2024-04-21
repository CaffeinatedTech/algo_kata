package main

import "fmt"
import "time"
import "encoding/json"
import "github.com/fatih/color"

var green = color.New(color.FgGreen).Add(color.Bold).SprintFunc()
var red = color.New(color.FgRed).Add(color.Bold).SprintFunc()

func (m model) welcomMessage() string {
	msg := "Welcome to Algorithm Kata\n\n"

	msg += "Please select the languages, and algorithms that you would like to practice.\n\n"
	msg += "Languages:\n"
	for i, lang := range m.session.Languages {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checked := " "
		if lang.Selected {
			checked = "x"
		}
		msg += fmt.Sprintf("%s [%s] %s\n", cursor, checked, lang.Name)
	}

	numLangs := len(m.session.Languages)
	msg += "\nAlgorithms:\n"
	for i, alg := range m.session.Algorithms {
		cursor := " "
		if m.cursor == i+numLangs {
			cursor = ">"
		}
		checked := " "
		if alg.Selected {
			checked = "x"
		}
		msg += fmt.Sprintf("%s [%s] %s\n", cursor, checked, alg.Name)
	}
	return msg
}

func (m model) askForNumber() string {
	maxQuestions := m.countChecked()
	msg := fmt.Sprintf("How many questions would you like to answer? (max: %d)\n", maxQuestions)
	return msg
}

func (m model) countChecked() int {
	count := 0
	for _, lang := range m.session.Languages {
		if lang.Selected {
			count++
		}
	}
	for _, alg := range m.session.Algorithms {
		if alg.Selected {
			count++
		}
	}
	return count
}

func (m model) intermissionMessage() string {
	msg := ""
	if m.lastAnswerCorrect != "" {
		msg += fmt.Sprintf("\nYour answer was %s\n", m.lastAnswerCorrect)
	}
	return fmt.Sprintf("%s\nThis round will be in %s, press Enter when ready.\n", msg, green(m.nextQuestion.language))
}

func (m model) questionMessage() string {
	msg := fmt.Sprintf("\nThis round will be in %s\n", green(m.nextQuestion.language))
	msg += fmt.Sprintf("\nAlgorithm: %s\n", green(m.nextQuestion.algorithm.Name))

	if m.nextQuestion.algorithm.Sorted {
		msg += "This array is already sorted.\n"
	}

	if m.nextQuestion.algorithm.Type == "search" {
		msg += fmt.Sprintf("Find the index of %d in the following array:\n", m.nextQuestion.expectedResult)
	}

	questionJson, _ := json.Marshal(m.nextQuestion.array)
	questionJsonString := string(questionJson)
	msg += fmt.Sprintf("\nArray: %v\n", questionJsonString)
	return msg
}

func printResults(results []Result) string {
	msg := "Results:\n"
	for i, result := range results {
		thisResultCorrect := red("Incorrect")
		if result.Correct {
			thisResultCorrect = green("Correct")
		}
		msg += fmt.Sprintf("Algorithm %d: %s in %s %s %s\n", i+1, result.Algorithm, result.Language, result.Time.Round(time.Second), thisResultCorrect)
	}
	msg += "\nPress any key to exit.\n"
	return msg
}
