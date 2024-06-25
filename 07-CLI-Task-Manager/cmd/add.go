package cmd

import (
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
	},
}

func init(){
	RootCmd.AddCommand(addCmd)
}