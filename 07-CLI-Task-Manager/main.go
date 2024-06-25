package main

import (
	"cli-task/cmd"
	"cli-task/db"
	"fmt"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

func main() {
	home, err := homedir.Dir()  
	if err != nil {
		panic(err)
	}
	dbPath  := filepath.Join(home, "tasks.db")
	err = db.Init(dbPath)
	if err != nil {
		panic(err)
	}

	// defer db.Close()
	fmt.Println("db success")

	cmd.RootCmd.Execute()
}