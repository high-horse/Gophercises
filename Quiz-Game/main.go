package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	// func String(name string, value string, usage string) *string
	filename := flag.String("file", "questions.csv", "question, answers csv")
	timelimit := flag.Int("limit", 60, "time limit for quiz in seconds")
	shuffle := flag.Bool("shuffle", false, "shuffle questions (bool)")

	flag.Parse()

	file, err := os.Open(*filename)
	checkErr(err, "error opening file")

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	checkErr(err, "error reading csv")

	if *shuffle {
		shuffleRecord(records)
	}

	// shuffle the records

	timer := time.NewTimer(time.Duration(*timelimit) * time.Second)

	correct := 0
	// problemloop is a lebel
problemloop:
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
			// return
			break problemloop
		case ans := <-ansCh:
			if strings.EqualFold(ans, answer) {
				correct++
				fmt.Println("correct!")
			} else {
				fmt.Println("incorrect")
			}
		}
	}
	fmt.Printf("Points earned: %d/%d", correct, len(records))
}

func checkErr(err error, cause string) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR \n%s\n%v\n", cause, err)
		os.Exit(1)
	}
}

func shuffleRecord(records [][]string) {
	rand.Seed(time.Now().UnixNano())
	for i := len(records) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		records[i], records[j] = records[j], records[i]
	}
}
