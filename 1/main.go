package main

import (
    "fmt"
    "encoding/csv"
    "flag"
    "os"
    "strings"
    "time"
)

type Question struct {
    Formula string
    Answer string
}

func parseLines(lines [][]string) []Question {
    questions := make([]Question, len(lines))
    for i, line := range lines {
        questions[i] = Question{
            Formula: line[0],
            Answer: strings.TrimSpace(line[1]),
        }
    }
    return questions
}

func ask(questions []Question, corrects *int, finish chan int) {
    var answer string
    for _, question := range questions {
        fmt.Println(question.Formula)
        fmt.Scanf("%s", &answer)
        if answer == question.Answer {
            *corrects++
        }
    }
    finish <- 0
}

func exit(message string) {
    fmt.Printf(message)
    os.Exit(1)
}

func main() {
    csvFilename := flag.String("csv", "hoge.csv", "a csv file in the format of 'question,answer'")
    timeLimit := flag.Int("limit", 30, "time limitation of ask duration")
    flag.Parse()

    file, err := os.Open(*csvFilename)
    if err != nil {
        exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
    }

    timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

    reader := csv.NewReader(file)
    lines, err := reader.ReadAll()
    if err != nil {
        exit("Failed to parse the provided CSV file.")
    }
    questions := parseLines(lines)

    corrects := 0
    answerCh := make(chan string)

questionloop:
    for _, question := range questions {
        fmt.Println(question.Formula)
        go func() {
            answer := ""
            fmt.Scanf("%s", &answer)
            answerCh <- answer
        }()

        select {
        case <- timer.C:
            fmt.Println("time over.")
            break questionloop
        case answer := <- answerCh:
            if answer == question.Answer {
                corrects++
            }
        }
    }

    fmt.Printf("%d / %d. \n", corrects, len(questions))
}