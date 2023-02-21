package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	var answer string
	var score int

	// Test value
	csvtext := "5+5,10 7+3,10 1+1,2 8+3,11 1+2,3 8+6,14 3+1,4 1+4,5 5+1,6 2+3,5 3+3,6 2+4,6 5+2,7"

	// Read in csv text. Set delimiter to spaces.
	r := csv.NewReader(strings.NewReader(csvtext))
	r.Comma = ' '
	records, err := r.ReadAll()
	if err != nil{
		fmt.Println(err)
	}

	// Take read-in data and send to quiz struct
	quiz := createStruct(records[0])
	
	// Generate quiz
	for i, questions := range quiz {
		fmt.Printf("Question %d :- %s? ",i+1,questions.Question)
		fmt.Scan(&answer)
		if strings.ToLower(strings.Trim(answer," ")) == strings.ToLower(strings.Trim(questions.Answer, " ")){
			score++;
		}
	}

	// Score
	fmt.Printf("You answered %d questions correctly out of a total of %d questions.\n",score, len(quiz))
}

func createStruct(data []string) []Quiz {
    var quiz []Quiz
    for _, line := range data {
		var question Quiz
		s := strings.Split(line, ",")
		question.Question, question.Answer = s[0], s[1]
		quiz = append(quiz, question)
	}
	return quiz
}
