package util_help

import (
	"fmt"
	effectprofileid "gctts/effect_profile_id"
)

func HelpListEffect() {
	fmt.Println("list effect that can use : ")
	for i, ap := range effectprofileid.EffectAudio.AudioProfile {
		fmt.Println(i, ").", ap)
	}
}
