package util_help

import "fmt"

func HelpListEncoding() {
	fmt.Println("list permitted encoding audio")
	enc := []string{"MP3", "OGG_OPUS", "MULAW", "ALAW", "PCM"}
	fmt.Println(enc)
}
