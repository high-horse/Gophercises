package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "CLI for managing your tasks",
	// Run: func(cmd *cobra.Command, args []string) {
	//   // Do Stuff Here
	// },
}