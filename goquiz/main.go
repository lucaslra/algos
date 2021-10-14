package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

type QuizRound struct {
	Question string
	Answer   string
}

func main() {
	csvData := readProblemFile()
	defer func(csvData *os.File) {
		err := csvData.Close()
		if err != nil {
			panic(err)
		}
	}(csvData)

	problems := parseProblems(csvData)

	shouldStart := confirmStart(problems)

	if shouldStart {
		score := runQuiz(problems)

		fmt.Println("Score:", score)
	}
}

func confirmStart(problems []*QuizRound) bool {
	fmt.Printf("> The quiz is about to start. \n> There will be %0d questions.\n> Your time limit is of %s\n> Start? [y/N]\n", len(problems), timeLimit(1))
	ch, _, err := keyboard.GetSingleKey()
	if err != nil {
		log.Fatal(err)
	}

	inp := string(ch)
	fmt.Println(inp)
	if inp == "Y" || inp == "y" {
		return true
	}
	return false
}

func timeLimit(duration time.Duration) string {
	return "30 seconds"
}

func runQuiz(problems []*QuizRound) (score int) {
	fmt.Println("\n> Starting Quiz!")
	for i, round := range problems {
		fmt.Printf("\n> Question %0d: %s\n", i, round.Question)
		fmt.Print("> Answer: ")

		var answer string
		_, err := fmt.Scanln(&answer)
		if err != nil {

			log.Fatal("Could not read answer", err)
		}

		if answer = strings.TrimSpace(answer); answer == round.Answer {
			fmt.Println("Right answer!")
			score++
		} else {
			fmt.Println("Wrong answer!")
			score--
		}
	}

	return
}

func parseProblems(data *os.File) []*QuizRound {
	csvReader := csv.NewReader(data)

	quizRounds := []*QuizRound{}
	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		quizRounds = append(quizRounds, &QuizRound{
			Question: row[0],
			Answer:   row[1],
		})
	}

	return quizRounds
}

func readProblemFile() *os.File {
	args := os.Args[1:]
	var problemFile string
	if len(args) <= 0 {
		problemFile = "problems.csv"
	} else {
		problemFile = args[0]
	}

	data, err := os.Open(problemFile)
	if err != nil {
		log.Fatalf("Failed to read problem file: %q", problemFile)
	}

	return data
}
