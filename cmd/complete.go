/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/ceckles/GO-Task-Cli/utils"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "Mark a task as complete",
	Long: `Mark a task as complete by providing its ID. This command updates the task
status in the CSV file to 'true'. 

Example usage:
complete 1`,
	Run: func(cmd *cobra.Command, args []string) {
		// Check if the required argument is provided
		if len(args) < 1 {
			fmt.Println("Error: missing required argument (task ID)")
			cmd.Help() // Optionally display help if arguments are missing
			return
		}

		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error converting argument to integer: %v\n", err)
			return
		}

		// Load the CSV file
		file, err := utils.LoadFile("tasks.csv")
		if err != nil {
			fmt.Println("Error loading file:", err)
			return
		}
		defer utils.CloseFile(file) // Ensure file is closed even if an error occurs

		// Complete the task
		if err := complete(id, file); err != nil {
			fmt.Println("Error completing task:", err)
			return
		}

		fmt.Println("Task completed successfully")
	},
}

func complete(id int, file *os.File) error {
	fmt.Println("complete task called")

	// Seek to the beginning of the file in case it's not at the start
	if _, err := file.Seek(0, os.SEEK_SET); err != nil {
		return fmt.Errorf("failed to seek to beginning of file: %w", err)
	}

	// Create a new CSV reader
	reader := csv.NewReader(file)
	// Read all records from the CSV
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("failed to read all records from CSV: %w", err)
	}

	// Check if the CSV has at least one record (the header row)
	if len(records) == 0 {
		fmt.Println("No records found in CSV.")
		return nil
	}

	// Use a buffer to write records back to the file
	var updatedRecords [][]string
	taskFound := false

	// Process records
	for _, record := range records {
		if len(record) > 0 {
			// Assuming the ID is in the first column (index 0) and the status in the fourth column (index 3)
			if record[0] == strconv.Itoa(id) {
				if record[3] == "true" {
					return fmt.Errorf("task %d is already marked as complete", id)
				}
				record[3] = "true" // Mark as complete
				taskFound = true
			}
		}
		updatedRecords = append(updatedRecords, record)
	}

	if !taskFound {
		return fmt.Errorf("task %d not found", id)
	}

	// Seek to the beginning of the file to overwrite
	if _, err := file.Seek(0, os.SEEK_SET); err != nil {
		return fmt.Errorf("failed to seek to beginning of file: %w", err)
	}

	// Truncate the file to remove any existing content
	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("failed to truncate file: %w", err)
	}

	// Create a new CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write all updated records to the file
	for _, record := range updatedRecords {
		if err := writer.Write(record); err != nil {
			return fmt.Errorf("failed to write record to CSV: %w", err)
		}
	}

	return nil
}
func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
