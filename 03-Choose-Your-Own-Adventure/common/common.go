package common

import (
	"log"
	"os"
)

func Check(err error, msg string) {
	if err != nil {
		log.Fatalf("ERROR: %s \n%v", msg, err)
		os.Exit(1)
	}
}