package cmd

import (
	"encoding/csv"
	"fmt"

	"github.com/ceckles/GO-Task-Cli/utils"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [task ID]",
	Short: "Delete a task by its ID and re-sequence task IDs",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID := args[0]

		// Open and lock the CSV file using utils.LoadFile
		file, err := utils.LoadFile(csvFilePath)
		if err != nil {
			fmt.Println("Error opening CSV file:", err)
			return
		}
		defer func() {
			if err := utils.CloseFile(file); err != nil {
				fmt.Println("Error closing CSV file:", err)
			}
		}()

		// Read all records from the CSV file
		reader := csv.NewReader(file)
		records, err := reader.ReadAll()
		if err != nil {
			fmt.Println("Error reading CSV file:", err)
			return
		}

		// Find and remove the record with the given task ID
		var newRecords [][]string
		found := false
		for _, record := range records {
			if record[0] == taskID {
				found = true
				continue
			}
			newRecords = append(newRecords, record)
		}

		if !found {
			fmt.Printf("Task ID %s not found.\n", taskID)
			return
		}

		// Re-sequence task IDs
		for i, record := range newRecords {
			record[0] = fmt.Sprintf("%d", i+1)
			newRecords[i] = record
		}

		// Truncate the file and write the updated records
		if err := file.Truncate(0); err != nil {
			fmt.Println("Error truncating CSV file:", err)
			return
		}
		if _, err := file.Seek(0, 0); err != nil {
			fmt.Println("Error seeking in CSV file:", err)
			return
		}

		writer := csv.NewWriter(file)
		err = writer.WriteAll(newRecords)
		if err != nil {
			fmt.Println("Error writing to CSV file:", err)
			return
		}
		writer.Flush()

		fmt.Printf("Task ID %s deleted and remaining IDs re-sequenced.\n", taskID)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
