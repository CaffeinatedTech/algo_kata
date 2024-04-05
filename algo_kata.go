package main

import "flag"
import "fmt"
import "strings"
import "time"

type languages []string

func (l *languages) String() string {
	return fmt.Sprint(*l)
}
func (l *languages) Set(value string) error {
	for _, lang := range strings.Split(value, ",") {
		*l = append(*l, lang)
	}
	return nil
}

type Algorithm struct {
	Name   string `mapstructure:"name"`
	Type   string `mapstructure:"type"`
	Sorted bool   `mapstructure:"sorted"`
}

type Session struct {
	Algorithms []Algorithm
	Languages  []string
	Num        int
}

type Result struct {
	Algorithm string
	Language  string
	Time      time.Duration
	Correct   bool
}

func main() {
	var languagesFlag languages
	flag.Var(&languagesFlag, "languages", "A comma separated list of languages to practice.")
	flag.Var(&languagesFlag, "l", "A comma separated list of languages to practice.")
	flag.Parse()

  checkConfig()
	algos, languages := getConfig()
	if len(languagesFlag) > 0 {
		languages = languagesFlag
	}
	welcomMessage(algos, languages)
	num := askForNumber()
	s := Session{Algorithms: algos, Languages: languages, Num: num}
	s = checkPracticeNumer(s)
	fmt.Println("Starting session with", s.Num, "questions.")
	results := runTheSession(s)

	// Print out the results
	printResults(results)
}
