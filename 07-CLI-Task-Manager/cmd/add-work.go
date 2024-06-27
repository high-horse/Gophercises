package cmd

import (
	"cli-task/db"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var addWorkCmd = &cobra.Command{
	Use:   "add-work",
	Short: "add a new work to the list",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("add called")
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to task list \n", task)
		id, err := db.CreateWork(task)
		println("id created : ", id)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("Created Task ID: %d\n", id)
	},
}

func init() {
	RootCmd.AddCommand(addWorkCmd)
}
