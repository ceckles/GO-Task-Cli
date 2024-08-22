package utils

import (
	"os"
	"syscall"
)

// LoadFile opens a file, creates it if it doesn't exist, and locks it exclusively.
// It returns a pointer to the file and an error if the operation fails.
//
// Parameters:
// - filepath: The path to the file to open or create.
//
// Returns:
// - *os.File: The file pointer to the opened file.
// - error: An error if the operation fails, otherwise nil.
func LoadFile(filepath string) (*os.File, error) {
	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}

	// Obtain an exclusive lock on the file descriptor
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_EX); err != nil {
		_ = f.Close()
		return nil, err
	}

	return f, nil
}

// closeFile unlocks and closes the file.
// This function ensures the file lock is released before closing the file.
func CloseFile(f *os.File) error {
	if err := syscall.Flock(int(f.Fd()), syscall.LOCK_UN); err != nil {
		return err
	}
	return f.Close()
}
