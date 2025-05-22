package util_help

import "fmt"

func HelpListEncoding() {
	enc := []string{"MP3", "OGG_OPUS", "MULAW", "ALAW", "PCM"}
	fmt.Println(enc)
}
