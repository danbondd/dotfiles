package main

import (
	"fmt"
	"os"
)

func main() {
	homeDir := os.Getenv("HOME")
	if len(homeDir) == 0 {
		fmt.Fprintf(os.Stderr, "error getting $HOME env variable")
		return
	}

	files, err := newConfig(fileReader, newJSONDecoder, homeDir)
	if err != nil {
		return
	}

	for file, loc := range files {
		if err := createSymLink(file, loc, homeDir); err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			return
		}
	}

	fmt.Println("Setup complete.")
}

func createSymLink(file, loc, homeDir string) error {
	if err := fileExists(loc); err != nil {
		return err
	}

	return os.Symlink(loc, fmt.Sprintf("%s/.ditfiles/files/%s", homeDir, file))
}

func fileExists(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return fmt.Errorf("file does not exist: %s", path)
	}

	return nil
}
