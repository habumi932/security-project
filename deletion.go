package main

import (
	"fmt"
	"path/filepath"
	"os"
)

func deleteFiles(root string) {
	// looping through target files
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// skip if directory
		if !info.IsDir() {
			// delete the file
			fmt.Println("Deleting " + path + "...")
			// remove files
			err := os.Remove(path)
			if err != nil {
				fmt.Println("error while deleting files")
			}
		}
		return nil
	})
	fmt.Println("All files on your device have been deleted.")
}
