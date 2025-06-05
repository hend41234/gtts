package utilopenxml

import (
	"io"
	"log"
	"os"
)

func ReadXML(filename string) (result string) {
	file, _ := os.Open(filename)
	defer file.Close()
	readFile, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("error read xml file")
	}
	result = string(readFile)
	return
}

func CheckXMLFile(fileName string) bool {
	_, err := os.Stat(fileName)
	return os.IsExist(err)
}
