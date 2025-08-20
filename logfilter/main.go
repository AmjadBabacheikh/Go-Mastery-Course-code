package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var logFilter = &cobra.Command{
	Use:   "logFilter",
	Short: "Filter log file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This runs if no subcommand is provided")
	},
}

var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "Filter log file",
	Long:  "Filter log file by level and keyword",
	Run: func(cmd *cobra.Command, args []string) {
		file, _ := cmd.Flags().GetString("file")
		keyword, _ := cmd.Flags().GetString("keyword")
		level, _ := cmd.Flags().GetString("level")
		if file == "" {
			log.Fatal("Please provide a log file using --file")
		}

		lines, err := readLines(file)
		if err != nil {
			log.Fatalf("Could not read file: %v", err)
		}
		filterd := filterLogs(lines, level, keyword)

		for _, line := range filterd {
			fmt.Println(line)
		}
	},
}

func init() {
	filterCmd.Flags().StringP("file", "f", "sample.log", "Path to log file")
	filterCmd.Flags().StringP("keyword", "k", "", "Keyword to search in log line")
	filterCmd.Flags().StringP("level", "l", "", "Log level to match")
	filterCmd.MarkFlagRequired("level")
}

func readLines(path string) ([]string, error) {
	// Create a slice to store the lines
	var lines []string

	// open the file
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Loop through the lines and add them to the slice
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines, nil
}

func filterLogs(lines []string, level string, keyword string) []string {
	var result []string
	for _, line := range lines {
		if level != "" && !strings.Contains(line, level) {
			continue
		}
		if keyword != "" && !strings.Contains(strings.ToLower(line), strings.ToLower(keyword)) {
			continue
		}
		result = append(result, line)
	}
	return result
}

func main() {
	logFilter.AddCommand(filterCmd)
	logFilter.Execute()
}
