package main

import (
	"context"
	"errors"
	"os"

	"github.com/spf13/cobra"
)

var (
	pathFlag              string
	checkInvalidTodosFlag bool
	checkTestTagsFlag     bool
)

var rootCmd = &cobra.Command{
	Use:          "checker",
	Version:      "1.0.0",
	SilenceUsage: true,
	Short:        "Check the code based on custom requirements for Cryptellation",
	RunE: func(cmd *cobra.Command, args []string) error {
		var errInvalidTodos error
		if checkInvalidTodosFlag {
			errInvalidTodos = checkInvalidTodos(cmd, args)
		}

		var errTestTags error
		if checkTestTagsFlag {
			errTestTags = checkTestTags(cmd, args)
		}

		return errors.Join(
			errInvalidTodos,
			errTestTags,
		)
	},
}

func main() {
	// Set flags
	rootCmd.PersistentFlags().StringVarP(&pathFlag, "path", "p", ".", "set the path to check for invalid todos")
	rootCmd.PersistentFlags().BoolVar(&checkInvalidTodosFlag, "check-invalid-todos", true, "check invalid todos")
	rootCmd.PersistentFlags().BoolVar(&checkTestTagsFlag, "check-test-tags", true, "check test tags")
	rootCmd.AddCommand(checkTodosCmd)
	rootCmd.AddCommand(checkTestTagsCmd)

	// Execute command
	if err := rootCmd.ExecuteContext(context.Background()); err != nil {
		os.Exit(1)
	}
}
