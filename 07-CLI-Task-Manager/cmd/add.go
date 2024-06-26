package cmd

import (
	"cli-task/db"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new task to the list",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	// fmt.Println("add called")
	// 	task := strings.Join(args, " ")
	// 	fmt.Printf("Added \"%s\" to task list \n", task)
	// 	id, err := db.CreateWork(task)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}

	// 	fmt.Printf("Created Task ID: %d\n", id)
	// },
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Printf("something went wrong : \n", err.Error())
			os.Exit(1)
		}

		fmt.Printf("Added \"%s\" to task list \n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
