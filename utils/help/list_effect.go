package util_help

import (
	"fmt"
	effectprofileid "gctts/effect_profile_id"
)

func HelpListEffect() {
	for i, ap := range effectprofileid.EffectAudio.AudioProfile {
		fmt.Println(i, ").", ap)
	}
}
