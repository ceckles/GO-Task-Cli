package cmd

import (
    "encoding/csv"
    "fmt"
    "os"
    "text/tabwriter"
    //"time"

    "github.com/ceckles/GO-Task-Cli/utils"
    "github.com/spf13/cobra"
    //"github.com/mergestat/timediff"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Long: `This command lists all tasks with their details.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		isSet := cmd.Flags().Lookup("all").Changed
		if isSet {
			file, err := utils.LoadFile("tasks.csv")
			if err != nil {
				fmt.Println("Error loading file:", err)
				return
			}
			if err := listAll(file); err != nil {
				fmt.Println("Error listing all tasks:", err)
			}
			defer utils.CloseFile(file)
		} else {
			list()
		}
	},
}

// listAll reads the CSV file and displays formatted task details
func listAll(file *os.File) error {
	fmt.Println("list flag -a called")

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

	// Create a new tabwriter
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 2, ' ', 0)

	// Print the header
	fmt.Fprintln(writer, "ID\tTask\tCreated\tDone")

	// Print each record
	for _, record := range records {
		if len(record) < 4 {
			fmt.Println("Invalid record:", record)
			continue
		}
		// Assuming fields are in the following order: ID, Task, Created, Done
		id := record[0]
		task := record[1]
		created := record[2]
		done := record[3]

		// Parse the created date
		createdTime, err := utils.ParseDate(created)
		if err != nil {
			fmt.Println("Invalid date format:", created)
			continue
		}

		// Calculate the time difference
		diff := utils.FormatTimeDiff(createdTime)

		// Print the record with alignment
		fmt.Fprintf(writer, "%s)\t%s\t%s\t%s\n", id, task, diff, done)
	}

	// Flush the writer to ensure all data is written
	if err := writer.Flush(); err != nil {
		return fmt.Errorf("failed to flush tabwriter: %w", err)
	}
	return nil
}

func list() {
	fmt.Println("list No Flag called")
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Add -a flag to the list command
	listCmd.Flags().BoolP("all", "a", false, "Show all tasks")
}