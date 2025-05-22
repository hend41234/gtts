package util_openxml

import (
	"io"
	"log"
	"os"
)

func ReadXML() (result string) {
	file, _ := os.Open("ssml_sample.xml")
	defer file.Close()
	readFile, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("error read xml file")
	}
	result = string(readFile)
	return
}
