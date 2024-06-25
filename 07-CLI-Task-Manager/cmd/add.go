package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add a new task to the list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("add called")
	},
}

func init(){
	RootCmd.AddCommand(addCmd)
}