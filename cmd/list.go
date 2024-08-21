/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"syscall"
	//"strings"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
		//check if -a is called if so display all tasks and value if not show short list.
		isSet:=cmd.Flags().Lookup("all").Changed
		if(isSet){
			file, err := loadFile("tasks.csv")
			if err != nil {
				fmt.Println("Error loading file:", err)
				return
			}
			if err := listAll(file); err != nil {
				fmt.Println("Error listing all tasks:", err)
			}
			defer closeFile(file)
		} else {
			list()
		}
	},
}

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

	// Print the header
	// fmt.Printf("%-5s %-50s %-20s %-5s\n", "ID", "Task", "Created", "Done")
	// fmt.Println(strings.Repeat("-", 80)) // Adjusted the length based on column widths

	// Print each record
	for _, record := range records {
		if len(record) < 4 {
			fmt.Println("Invalid record:", record)
			continue
		}
		fmt.Printf("%-5s %-50s %-20s %-5s\n", record[0], record[1], record[2], record[3])
	}

	return nil
}
func list() {
	fmt.Println("list No Flag called")
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

    //flag to -a
	listCmd.Flags().BoolP("all", "a", false, "Show all tasks")
}


func loadFile(filepath string) (*os.File, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf("failed to open file for reading")
	}

    // Exclusive lock obtained on the file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

func closeFile(f *os.File) error {
	syscall.Flock(int(f.Fd()), syscall.LOCK_UN)
	return f.Close()
}