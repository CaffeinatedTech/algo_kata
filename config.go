package main

import "fmt"
import "github.com/spf13/cast"
import "github.com/spf13/viper"

func checkConfig() {
  algorithms, languages := getConfig()
  if len(languages) == 0 {
    fmt.Println("No languages found in config file")
    panic(1)
  }
  if len(algorithms) == 0 {
    fmt.Println("No algorithms found in config file")
    panic(1)
  }
  // Check each algorithm in config file to make sure they have `name`, `type` keys.
  for _, alg := range algorithms {
    if alg.Name == "" || alg.Type == "" {
      fmt.Println("One of the Algorithms in the config.toml file is missing either name or type")
      panic(1)
    }
  }
}

func getConfig() ([]Algorithm, []string) {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file", err)
		panic(err)
	}
	var c []Algorithm
	alg, ok := viper.Get("algorithm").([]interface{})
	if !ok {
		fmt.Println("Error getting algorithms")
		panic(1)
	} else {
		for _, table := range alg {
			if m, ok := table.(map[string]interface{}); ok {
				c = append(c, Algorithm{Name: cast.ToString(m["name"]), Type: cast.ToString(m["type"]), Sorted: cast.ToBool(m["sorted"])})
			}
		}
	}
	languages := viper.GetStringSlice("languages")
	return c, languages
}

func configText(c []Algorithm, languages []string) (string, string) {
	languages_text := ""
	for _, lang := range languages {
		if languages_text != "" {
			languages_text += ", "
		}
		languages_text += lang + " "
	}
	algorithms_text := ""
	for _, alg := range c {
		if algorithms_text != "" {
			algorithms_text += ", "
		}
		algorithms_text += alg.Name
	}
	return languages_text, algorithms_text
}

func checkPracticeNumer(s Session) Session {
	maxNum := len(s.Algorithms) * len(s.Languages)
  if s.Num < 1 {
    fmt.Println("You must do at least one question.")
    s.Num = 1
  }
	if s.Num > maxNum {
		fmt.Println("You can only do up to", maxNum, "questions.")
		s.Num = maxNum
	}
	return s
}
