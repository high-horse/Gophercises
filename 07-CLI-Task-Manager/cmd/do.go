package cmd

import (
	"cli-task/db"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "mark your task done",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	var ids []int
	// 	for _, arg := range args {
	// 		id, err := strconv.Atoi(arg)
	// 		if err != nil {
	// 			fmt.Println("Failed to parse the argument: ", arg)
	// 			continue
	// 		}
	// 		ids = append(ids, id)
	// 	}
	// 	fmt.Println(ids)

	// 	for _, id := range ids {
	// 		err := db.UpdateWork(id, true)
	// 		if err != nil {
	// 			fmt.Println("Failed to mark the task as done: ", id)
	// 			continue
	// 		}
	// 		fmt.Printf("Marked task %d as done\n", id)
	// 	}
	// },
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
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("something went wrong : error: ", err.Error())
			return
		}
		
		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number: ", id)
				continue
			}

			task := tasks[id-1] 
			err := db.DeleteTask(task.Key)
			if err != nil {
				fmt.Printf("failed to mark the task as done: %s\n something went wrong : \n error: \n", task.Value, err.Error())
				continue
			}
			fmt.Printf("Marked task %d as done\n", task.Key)
		}

	},
}

func init(){
	RootCmd.AddCommand(doCmd)
}