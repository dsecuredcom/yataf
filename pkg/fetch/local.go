package fetch

import (
	"fmt"
	"os"
)

func GetContentOfLocalFile(File string) string {
	fmt.Printf("\033[34m[i] Loading file from local storage: %s\033[0m\n", File)

	file, err := os.ReadFile(File)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(file)
}
