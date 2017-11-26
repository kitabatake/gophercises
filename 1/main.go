package main

import (
    "fmt"
    "encoding/csv"
    "flag"
    "os"
    "strings"
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

func ask(questions []Question) int {
    var corrects int
    var answer string
    for _, question := range questions {
        fmt.Println(question.Formula)
        fmt.Scanf("%s", &answer)
        if answer == question.Answer {
            corrects++
        }
    }
    return corrects
}

func exit(message string) {
    fmt.Printf(message)
    os.Exit(1)
}

func main() {
    csvFilename := flag.String("csv", "hoge.csv", "a csv file in the format of 'question,answer'")
    flag.Parse()

    file, err := os.Open(*csvFilename)
    if err != nil {
        exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *csvFilename))
    }

    reader := csv.NewReader(file)
    lines, err := reader.ReadAll()
    if err != nil {
        exit("Failed to parse the provided CSV file.")
    }

    questions := parseLines(lines)
    corrects := ask(questions)
    fmt.Println(corrects)
}