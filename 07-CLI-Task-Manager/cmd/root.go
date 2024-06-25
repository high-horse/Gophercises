package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "task is a very fast static site generator",
	// Run: func(cmd *cobra.Command, args []string) {
	//   // Do Stuff Here
	// },
}