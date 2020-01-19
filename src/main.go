package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type problem struct {
	ques string
	ans  string
}

func main() {
	// create a flag and parse it.
	csvFileName := flag.String("csv", "problems.csv.txt", "A csv file in the format question, answer.")
	timer := flag.Int("timer", 5, "Timer in seconds that will stop the quiz if answer doesn't come.")
	flag.Parse()

	lines, err := readFile(csvFileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while reading the file : %s \n %v", *csvFileName, err)
	}
	problems := parseFile(lines)

	finalScore := showQuiz(problems, timer)
	fmt.Printf("\n%s", finalScore)
}

// readFile takes the name of the file and returns the string slice of all the columns.
func readFile(csvFileName *string) ([][]string, error) {
	file, err := os.Open(*csvFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		return nil, err

	}
	return lines, nil
}

//parseFile takes the input as [][]string and parses it to array of struct problem.
func parseFile(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			ques: strings.TrimSpace(line[0]),
			ans:  strings.TrimSpace(line[1]),
		}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ret), func(i, j int) { ret[i], ret[j] = ret[j], ret[i] })

	return ret
}

//showQuiz displays the problem one by one and returns the final score.
func showQuiz(problems []problem, timer *int) string {
	correct := 0
	for i, p := range problems {
		timeLimit := time.NewTimer(time.Duration(*timer) * time.Second)
		fmt.Printf("Problem #%d: %s\t", i+1, p.ques)
		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timeLimit.C:
			return fmt.Sprintf("Final Score is %d of %d", correct, len(problems))
		case answer := <-answerCh:
			if answer == p.ans {
				correct++
			}
		}
	}
	return fmt.Sprintf("Final Score is %d of %d", correct, len(problems))
}
