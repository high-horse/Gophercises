package cmd

import (
	"cli-task/db"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doWorkCmd = &cobra.Command{
	Use:   "do-work",
	Short: "mark your work done",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument: ", arg)
				continue
			}
			ids = append(ids, id)
		}
		fmt.Println(ids)

		for _, id := range ids {
			err := db.UpdateWork(id, true)
			if err != nil {
				fmt.Println("Failed to mark the task as done: ", id)
				continue
			}
			fmt.Printf("Marked task %d as done\n", id)
		}
	},
}

func init(){
	RootCmd.AddCommand(doWorkCmd)
}