package main

import "fmt"
import "github.com/spf13/cast"
import "github.com/spf13/viper"
import tea "github.com/charmbracelet/bubbletea"


func (m model) validateOptions() bool {
  // Check of there are any languages selected
  noLangs := true
  for _, lang := range m.session.Languages {
    if lang.Selected {
      noLangs = false
      break
    }
  }
  if noLangs {
    return false
  }
  // Check of there are any algorithms selected
  noAlgs := true
  for _, alg := range m.session.Algorithms {
    if alg.Selected {
      noAlgs = false
      break
    }
  }
  if noAlgs {
    return false
  }
  return true
}

func (m model) checkConfig() tea.Msg {
  algorithms, languages, questionCount := getConfig()
  if len(languages) == 0 {
    return errMsg{fmt.Errorf("No languages found in config file")}
  }
  if len(algorithms) == 0 {
    return errMsg{fmt.Errorf("No algorithms found in config file")}
  }
  // Check each algorithm in config file to make sure they have `name`, `type` keys.
  for _, alg := range algorithms {
    if alg.Name == "" || alg.Type == "" {
      return errMsg{fmt.Errorf("One of the Algorithms in the config.toml file is missing either name or type")}
    }
  }
  config := Config{Algorithms: algorithms, Languages: languages, QuestionCount: questionCount}
  return configMsg(config)
}

func saveConfig(alg []Algorithm, lang []Language, questionCount int) {
  viper.Set("algorithm", alg)
  viper.Set("language", lang)
  viper.Set("question_count", questionCount)
  viper.WriteConfig()
}

func getConfig() ([]Algorithm, []Language, int) {
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
				c = append(c, Algorithm{Name: cast.ToString(m["name"]), Type: cast.ToString(m["type"]), Sorted: cast.ToBool(m["sorted"]), Selected: cast.ToBool(m["selected"])})
			}
		}
	}
  var languages []Language
  lang, ok := viper.Get("language").([]interface{})
  if !ok {
    fmt.Println("Error getting languages")
    panic(1)
  } else {
    for _, table := range lang {
      if m, ok := table.(map[string]interface{}); ok {
      languages = append(languages, Language{Name: cast.ToString(m["name"]), Selected: cast.ToBool(m["selected"])})
      }
    }
  }
  questionCount := viper.GetInt("question_count")
	return c, languages, questionCount
}

