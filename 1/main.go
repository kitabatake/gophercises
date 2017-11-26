package main

import (
    "fmt"
    "encoding/csv"
    "log"
    "strings"
    "io"
    "io/ioutil"
    "strconv"
)

type Question struct {
    Formula string
    Answer int
}

func parse(s string) []Question {
    var questions []Question
    reader := csv.NewReader(strings.NewReader(s))
    for {
        record, err := reader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }

        answer, error := strconv.Atoi(record[1])
        if error != nil {
            panic(error)
        }
        questions = append(questions, Question{record[0], answer})
    }
    return questions
}

func ask(questions []Question) int {
    var corrects int
    var answer int
    for _, question := range questions {
        fmt.Println(question.Formula)
        fmt.Scanf("%d", &answer)
        if answer == question.Answer {
            corrects++
        }
    }
    return corrects
}

func main() {
    file, err := ioutil.ReadFile("./hoge.csv")
    if err != nil {
        panic(err)
    }
    
    questions := parse(string(file))
    corrects := ask(questions)
    fmt.Println(corrects)
}