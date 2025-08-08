package misc

import (
	"log"
	"os"
)

func DefineError(v error, msg string) {
	log.Fatalf("%s: %v\n", msg, v)
}

func GetFileSize(filepath string) int64 {
	fileInfo, err := os.Stat(filepath)
	if err != nil {
		DefineError(err, "Error in getting file size")
	}
	return fileInfo.Size()
}
