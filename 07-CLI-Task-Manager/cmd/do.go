package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "mark your task done",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		fmt.Printf("Added \"%s\" to task list \n", task)
	},
}

func init(){
	RootCmd.AddCommand(doCmd)
}