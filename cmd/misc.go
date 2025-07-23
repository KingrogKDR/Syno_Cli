package cmd

import (
	"fmt"
	"os"
)

func defineError(v error, msg string) {
	fmt.Println(msg)
	panic(v)
}

func getFileSize(filepath string) int64 {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		defineError(err, "Error in getting file size")
	}
	return fileInfo.Size()
}
