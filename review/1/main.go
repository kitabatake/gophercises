package main

import (
	"fmt"
	"flag"
	"os"
	"encoding/csv"
	"time"
)

func main() {
	csvFilePath := flag.String("csv", "problems.csv", "problems csv file path")
	timeLimit := flag.Int("limit", 10, "time limit for answers problems")
	flag.Parse()

	csvFile, err := os.Open(*csvFilePath)
	if err != nil {
		fmt.Printf("%s is not exists.", csvFile)
		os.Exit(1)
	}

	csvReader := csv.NewReader(csvFile)
	csvLines, err := csvReader.ReadAll()

	problems := make([]problem, len(csvLines))
	for i, line := range csvLines {
		problems[i] = problem{
			question: line[0],
			answer: line[1],
		}
	}

	corrects := 0
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	var scannedAnswer string
	answerCh := make(chan string)
	problemLoop:
	for _, problem := range problems {
		fmt.Println(problem.question)
		go func() {
			fmt.Scanf("%s", &scannedAnswer)
			answerCh <- scannedAnswer
		}()

		select {
		case <-timer.C:
			fmt.Println("time out!")
			break problemLoop
		case answer := <-answerCh:
			if answer == problem.answer {
				corrects++
			}
		}

	}
	fmt.Printf("result is %d/%d\n", corrects, len(problems))
}

type problem struct {
	question string
	answer string
}
