package main

import "fmt"
import "strconv"
import "time"
import "github.com/charmbracelet/bubbles/textinput"
import "github.com/charmbracelet/bubbles/stopwatch"
import tea "github.com/charmbracelet/bubbletea"

type Language struct {
	Name     string `mapstructure:"name"`
	Selected bool   `mapstructure:"selected"`
}

type Algorithm struct {
	Name     string `mapstructure:"name"`
	Type     string `mapstructure:"type"`
	Sorted   bool   `mapstructure:"sorted"`
	Selected bool   `mapstructure:"selected"`
}

type Session struct {
	Algorithms []Algorithm
	Languages  []Language
	Num        int
}

type Result struct {
	Algorithm string
	Language  string
	Time      time.Duration
	Correct   bool
}

type State int64

const (
	CHECK_CONFIG   State = iota // Check the configuration toml file
	WELCOME                     // Show the welcome screen, and allow selecting languages, and algorithms
	QUESTION_COUNT              // Ask the user for the number of questions they want to practice
	INTERMISSION                // Present the next question's language, and wait for user to be ready
	QUESTION                    // Show the question, and await the response
	COMPLETE                    // Show the final score table, and quit
)

type Config struct {
	Algorithms    []Algorithm
	Languages     []Language
	QuestionCount int
}

type statusMsg int
type stateMsg State
type configMsg Config
type errMsg struct{ err error }

func (e errMsg) Error() string { return e.err.Error() }

type nextQuestion struct {
  algorithm Algorithm
  language string
  array []int
  expectedResult int
  expectedResultIndex int
}

type model struct {
	state   State
	session Session
	results []Result
  nextQuestion nextQuestion
  lastAnswerCorrect string
	cursor  int
	num_ti  textinput.Model
  answer_ti textinput.Model
  stopwatch stopwatch.Model
}

func (m model) Init() tea.Cmd {
	return m.checkConfig
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" || msg.String() == "esc" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		}
		if m.state == WELCOME {
			optionCount := len(m.session.Languages) + len(m.session.Algorithms)
			switch msg.String() {
			case "j", "down":
				m.cursor++
				if m.cursor >= optionCount {
					m.cursor = 0
				}
			case "k", "up":
				m.cursor--
				if m.cursor < 0 {
					m.cursor = 0
				}
			case " ":
				if m.cursor < len(m.session.Languages) {
					m.session.Languages[m.cursor].Selected = !m.session.Languages[m.cursor].Selected
				} else {
					m.session.Algorithms[m.cursor-len(m.session.Languages)].Selected = !m.session.Algorithms[m.cursor-len(m.session.Languages)].Selected
				}
			case "enter":
				// Validate the options.  Must have at least one language, and at least one algorithm.
				if m.validateOptions() {
					m.state = QUESTION_COUNT
				}
			}
		} else if m.state == QUESTION_COUNT {
			switch msg.String() {
			case "enter":
				m.session.Num, _ = strconv.Atoi(m.num_ti.Value())
				if m.session.Num > m.countChecked() {
					m.session.Num = m.countChecked()
				} else if m.session.Num < 1 {
					m.session.Num = 1
				}
				saveConfig(m.session.Algorithms, m.session.Languages, m.session.Num)
				// Decide the next question and language, and generate the random array
        nextAlgorithm, nextLanguage := m.randomAlgorithmAndLang()
        m.nextQuestion = nextQuestion{algorithm: nextAlgorithm, language: nextLanguage}
        m.nextQuestion.array = randomArray(10, nextAlgorithm.Sorted)
        m.nextQuestion.expectedResult, m.nextQuestion.expectedResultIndex = expectedResult(m.nextQuestion.array)
				m.state = INTERMISSION
			}
		} else if m.state == INTERMISSION {
      switch msg.String() {
        case "enter":
          m.state = QUESTION
          return m, m.stopwatch.Start()
      }
    } else if m.state == QUESTION {
      switch msg.String() {
        case "enter":
          // Check the answer
          correct := answerCheck(m.nextQuestion.array, m.nextQuestion.expectedResultIndex, m.nextQuestion.algorithm.Type, m.answer_ti.Value())
          m.lastAnswerCorrect = green("Correct")
          if !correct {
            m.lastAnswerCorrect = red("Incorrect")
          }
          elapsedTime := m.stopwatch.Elapsed()
          m.results = append(m.results, Result{
            Algorithm: m.nextQuestion.algorithm.Name,
            Language: m.nextQuestion.language,
            Time: elapsedTime,
            Correct: correct,
          })
          if len(m.results) >= m.session.Num {
            m.state = COMPLETE
          } else {
            nextAlgorithm, nextLanguage := m.randomAlgorithmAndLang()
            m.nextQuestion = nextQuestion{algorithm: nextAlgorithm, language: nextLanguage}
            m.nextQuestion.array = randomArray(10, nextAlgorithm.Sorted)
            m.nextQuestion.expectedResult, m.nextQuestion.expectedResultIndex = expectedResult(m.nextQuestion.array)
            m.answer_ti.SetValue("")
            m.state = INTERMISSION
          }
          return m, tea.Batch(m.stopwatch.Stop(), m.stopwatch.Reset())
      }
    } else if m.state == COMPLETE {
      return m, tea.Quit
    }
	case errMsg:
		return m, tea.Quit
	case configMsg:
		m.session.Languages = msg.Languages
		m.session.Algorithms = msg.Algorithms
		m.session.Num = msg.QuestionCount
		m.state = WELCOME
		return m, nil
	}
	if m.state == QUESTION_COUNT {
		var cmd tea.Cmd
		m.num_ti, cmd = m.num_ti.Update(msg)
		return m, cmd
	}
  if m.state == INTERMISSION {
    // Have to update the stopwatch to carry out the reset after a question.
    var cmdStopwatch tea.Cmd
    m.stopwatch, cmdStopwatch = m.stopwatch.Update(msg)
    return m, cmdStopwatch
  }
  if m.state == QUESTION {
    var cmdAnswer tea.Cmd
    var cmdStopwatch tea.Cmd
    m.answer_ti, cmdAnswer = m.answer_ti.Update(msg)
    m.stopwatch, cmdStopwatch = m.stopwatch.Update(msg)
    return m, tea.Batch(cmdAnswer, cmdStopwatch)
  }
  // var cmdStopwatch tea.Cmd
  // m.stopwatch, cmdStopwatch = m.stopwatch.Update(msg)
	return m, nil
}

func (m model) View() string {
	switch m.state {
	case WELCOME:
		return m.welcomMessage()
	case QUESTION_COUNT:
		max_count := m.countChecked()
		return fmt.Sprintf(
			"How many questions will you do? (max: %d)\n\n%s\n\n%s",
			max_count,
			m.num_ti.View(),
			"(esc to quit)",
		) + "\n"
	case INTERMISSION:
    return m.intermissionMessage()
  case QUESTION:
    msg := m.questionMessage()
    return fmt.Sprintf(
      "%s\nTime Taken: %s\n\nYour answer: %s\n\n%s",
      msg,
      m.stopwatch.View(),
      m.answer_ti.View(),
      "(esc to quit)",
    ) + "\n"
  case COMPLETE:
    return printResults(m.results)
	}
	return ""
}

func main() {
	num_ti := textinput.New()
	num_ti.Focus()
	num_ti.CharLimit = 2
	num_ti.Width = 5
  answer_ti := textinput.New()
  answer_ti.Focus()
  answer_ti.Width = 50

	initialModel := model{
		state:  CHECK_CONFIG,
		cursor: 0,
		num_ti: num_ti,
    answer_ti: answer_ti,
    stopwatch: stopwatch.NewWithInterval(time.Second),
	}
	p := tea.NewProgram(initialModel)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error starting the program", err)
	}

}
