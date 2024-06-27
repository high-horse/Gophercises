package cmd

import (
	"cli-task/db"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var listWorkCmd = &cobra.Command{
	Use:   "list-work",
	Short: "list all of your work",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list work called")
		works, err := db.ListWork()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if len(works) == 0 {
			fmt.Println("You have no works")
			return
		}
		// fmt.Fprintf(os.Stdout, "%d: %s \t Completed: %t \t Completed At: %s\n", works[0].ID, works[0].Name, works[0].Completed, works[0].CompletedAt)
		for _, work := range works {
			fmt.Printf("%d: %s \t Completed: %t \t Completed At: %s\n", work.ID, work.Name, work.Completed, work.CompletedAt)
		}
	},
}

func init(){
	RootCmd.AddCommand(listWorkCmd)
}