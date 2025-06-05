package utilhelp

import "fmt"

func HelpListHz() {
	fmt.Println("list valid Rate Hertz : ")
	hz := []int{8000, 16000, 22050, 24000, 44100, 48000}
	fmt.Println(hz)

}
