package cmd

import (
	"cli-task/db"
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all of your tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		works, err := db.ListWork()
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, work := range works {
			fmt.Printf("%d: %s \t Completed: %t \t Completed At: %s\n", work.ID, work.Name, work.Completed, work.CompletedAt)
		}
	},
}

func init(){
	RootCmd.AddCommand(listCmd)
}