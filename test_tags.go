package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var checkTestTagsCmd = &cobra.Command{
	Use:   "test-tags",
	Short: "Check test tags",
	RunE:  checkTestTags,
}

func checkTestTags(cmd *cobra.Command, _ []string) error {
	invalidFiles, err := checkTestTagsOnDir(pathFlag)
	if err != nil {
		return err
	}

	if len(invalidFiles) > 0 {
		for _, line := range invalidFiles {
			cmd.Printf(" - %s\n", line)
		}
		return fmt.Errorf("found %d file with invalid tags", len(invalidFiles))
	}

	return nil
}

func checkTestTagsOnDir(path string) ([]string, error) {
	TestTags := make([]string, 0)
	err := filepath.Walk(path, func(path string, _ os.FileInfo, err error) error {
		// Return if there is an error
		if err != nil {
			return err
		}

		// Check if the file is a file to test
		if !strings.HasSuffix(path, "_test.go") {
			// Check if the file is a test file
			return nil
		}

		// Check if the file has a test tags
		invalid, err := checkFileHasTags(path)
		if err != nil {
			return err
		} else if invalid {
			TestTags = append(TestTags, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	return TestTags, nil
}

func checkFileHasTags(path string) (bool, error) {
	r := regexp.MustCompile(`// \+build`)

	// Open the file
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Check if the line contains a test tag
		if r.MatchString(line) {
			return false, nil
		}

		// Check if the line start with 'package'
		if strings.HasPrefix(line, "package") {
			// If the line starts with 'package', it means that the file is a test file
			// and it doesn't have a test tag
			return true, nil
		}
	}

	// Should not reach here
	return true, nil
}
