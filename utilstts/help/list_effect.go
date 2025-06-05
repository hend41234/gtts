package utilhelp

import (
	"fmt"

	effectprofileid "github.com/hend41234/gtts/effectprofileid"
)

func HelpListEffect() {
	fmt.Println("list effect that can use : ")
	for i, ap := range effectprofileid.EffectAudio.AudioProfile {
		fmt.Println(i, ").", ap)
	}
}
