package util_opentxt

import (
	"io"
	"os"
)

func ReadTXT(fileName string) (result string) {
	fileTXT, _ := os.Open(fileName)
	defer fileTXT.Close()
	readText, _ := io.ReadAll(fileTXT)
	result = string(readText)
	return
}
