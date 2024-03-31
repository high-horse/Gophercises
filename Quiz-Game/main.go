package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// func String(name string, value string, usage string) *string
	filename := flag.String("file", "questions.csv", "question, answers csv")
	timelimit := flag.Int("limit", 30, "time limit for quiz in seconds")
	flag.Parse()

	file, err := os.Open(*filename)
	checkErr(err, "error opening file")

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	checkErr(err, "error reading csv")

	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)

	correct := 0
	for index, rec := range records {
		question, answer := rec[0], rec[1]
		fmt.Printf("Question #%d: %s = ", index+1, question)

		ansCh := make(chan string)
		go func() {
			reader := bufio.NewReader(os.Stdin)
			ans, err := reader.ReadString('\n')
			checkErr(err, "error reading answer")
			ans = strings.TrimSpace(ans)
			ansCh <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nPoints earned: %d/%d\n", correct, len(records))
			return
		case ans := <-ansCh:
			if strings.EqualFold(ans, answer) {
				correct++
				fmt.Println("correct!")
			} else {
				fmt.Println("incorrect")
			}
		}

	}
	fmt.Printf("\nPoints earned: %d/%d", correct, len(records))

}

func checkErr(err error, cause string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR \n%s\n%v\n", cause, err)
		os.Exit(1)
	}
}
