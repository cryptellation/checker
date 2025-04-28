package main

import (
	"context"
	"errors"
	"os"

	"github.com/spf13/cobra"
)

var (
	pathFlag string
)

var rootCmd = &cobra.Command{
	Use:          "checker",
	Version:      "1.0.0",
	SilenceUsage: true,
	Short:        "Check the code based on custom requirements for Cryptellation",
	RunE: func(cmd *cobra.Command, args []string) error {
		return errors.Join(
			checkInvalidTodos(cmd, args),
			checkTestTags(cmd, args),
		)
	},
}

func main() {
	// Set flags
	rootCmd.PersistentFlags().StringVarP(&pathFlag, "path", "p", ".", "Set the path to check for invalid todos")
	rootCmd.AddCommand(checkTodosCmd)
	rootCmd.AddCommand(checkTestTagsCmd)

	// Execute command
	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		os.Exit(1)
	}
}
