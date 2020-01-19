package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

type problem struct {
	ques string
	ans  string
}

func main() {
	// create a flag and parse it.
	csvFileName := flag.String("csv", "problems.csv.txt", "A csv file in the format question, answer.")
	flag.Parse()

	file, err := os.Open(*csvFileName)
	if err != nil {
		fmt.Printf("Error while opening %s\n", *csvFileName)
		os.Exit(1)
	}
	defer file.Close()
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while reading the file: %v", err)
		os.Exit(1)
	}
	problems := parseFile(lines)

	finalScore := showQuiz(problems)
	fmt.Println(finalScore)
}

//parseFile takes the input as [][]string and parses it to array of struct problem.
func parseFile(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			ques: line[0],
			ans:  line[1],
		}
	}
	return ret
}

//showQuiz displays the problem one by one and returns the final score.
func showQuiz(problems []problem) string {
	correct := 0
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s\t", i+1, p.ques)
		var answer string
		fmt.Scanf("%s\n", &answer)
		if answer == p.ans {
			correct++
		}
	}
	return fmt.Sprintf("Final Score is %d of %d", correct, len(problems))
}
