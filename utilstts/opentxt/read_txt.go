package utilopentxt

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

func CheckTXTFile(fileName string) bool {
	_, err := os.Stat(fileName)
	return !os.IsExist(err)
}
