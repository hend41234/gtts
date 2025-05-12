package main

import (
	"gctts/text"
)

func main() {
	// file, _ := os.Open("sample1.json")
	// defer file.Close()
	// content := Res{}
	// readFile, _ := io.ReadAll(file)
	// json.Unmarshal(readFile, &content)
	// // fmt.Println(content.AudioContent)

	// audioData, _ := base64.StdEncoding.DecodeString(content.AudioContent)
	// err := ioutil.WriteFile("output1.mp3", audioData, 0644)
	// if err != nil {
	// 	log.Fatal("error write file mp3")
	// }
	// log.Println("success")

	// listvoices.GetListVoices()
	// fmt.Println(voices.ListVoices.Voice[0])
	text.Synthesize()
}
