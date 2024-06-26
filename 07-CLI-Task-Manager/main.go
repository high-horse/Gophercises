package main

import (
	"cli-task/cmd"
	"cli-task/db"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, err := homedir.Dir()
	must(err)
	println(home)
	dbPath := filepath.Join(home, "tasks.db")
	must(db.Init(dbPath))
	// must(db.InitBucket(dbPath))

	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
