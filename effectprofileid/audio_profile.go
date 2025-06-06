package effectprofileid

import (
	_ "embed"
	"encoding/json"
	"log"

	"github.com/hend41234/gtts/models"
)

var EffectAudio *models.EffectProfileID

//go:embed list_audio_profile.json
var audioProfile []byte

func init() {
	// fileJson, _ := os.Open("data/list_audio_profile.json")
	// json.NewDecoder(fileJson).Decode(&EffectAudio)

	EffectAudio = new(models.EffectProfileID)
	decodeErr := json.Unmarshal(audioProfile, &EffectAudio)
	if decodeErr != nil {
		log.Fatal("decode error list_audio_profile.json")
	}
}
