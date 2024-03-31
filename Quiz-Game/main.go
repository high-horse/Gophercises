package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// func String(name string, value string, usage string) *string
	filename := flag.String("file", "questions.csv", "question, answers csv")
	flag.Parse()

	file, err := os.Open(*filename)
	checkErr(err, "error opening file")

	defer file.Close()



	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	checkErr(err, "error reading csv")

	correct := 0
	incorrect := 0
	for index, rec := range records {

		question, answer := rec[0], rec[1]
		fmt.Printf("QN %d . %s\n", index+1, question)

		reader := bufio.NewReader(os.Stdin)
		ans, err := reader.ReadString('\n')
		checkErr(err, "error reading answer")
		ans = strings.TrimSpace(ans)

		if strings.ToLower(ans) == strings.ToLower(answer) {
			correct++
			fmt.Println("correct!")
		} else {
			incorrect ++
			fmt.Println("incorrect")
		}
	}
	println("Points earned: ", correct)
	println("Incorrect Points: ", incorrect)

}

func checkErr(err error, cause string) {
	if err != nil {
		exitGracefully(err, cause)
	}
}

func exitGracefully(err error, cause string) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n%v\n", cause, err)
	os.Exit(1)
}
