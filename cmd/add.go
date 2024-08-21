package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/ceckles/GO-Task-Cli/utils"
	"github.com/spf13/cobra"
)

var csvFilePath = "tasks.csv" // Path to your CSV file

var addCmd = &cobra.Command{
	Use:   "add [task details]",
	Short: "Add a new task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskDetails := args[0]

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

		// Get the next task ID
		nextID, err := getNextID(file)
		if err != nil {
			fmt.Println("Error generating task ID:", err)
			return
		}

		// Get the current timestamp
		timestamp := time.Now().Format("2006-01-02 15:04:05.999999999 -0700 MST")

		// Write the new task to the CSV file
		writer := csv.NewWriter(file)
		err = writer.Write([]string{
			strconv.Itoa(nextID),
			taskDetails,
			timestamp,
			"false",
		})
		if err != nil {
			fmt.Println("Error writing to CSV file:", err)
			return
		}
		writer.Flush()

		fmt.Printf("Task added: %s\n", taskDetails)
	},
}

// Modified getNextID to take *os.File as a parameter
func getNextID(file *os.File) (int, error) {
	// Reset file pointer to the beginning
	if _, err := file.Seek(0, 0); err != nil {
		return 0, err
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}

	if len(records) == 0 {
		return 1, nil
	}

	lastRecord := records[len(records)-1]
	lastID, err := strconv.Atoi(lastRecord[0])
	if err != nil {
		return 0, err
	}

	return lastID + 1, nil
}

func init() {
	rootCmd.AddCommand(addCmd)
}
