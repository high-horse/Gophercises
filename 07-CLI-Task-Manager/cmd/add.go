package cmd

import (
	"cli-task/db"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("add called")
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to task list \n", task)
		id, err := db.CreateWork(task)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Created Task ID: %d\n", id)
	},
}

func init(){
	RootCmd.AddCommand(addCmd)
}