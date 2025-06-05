package effectprofileid

import (
	"encoding/json"
	"github.com/hend41234/gtts/models"
	"os"
)

var EffectAudio *models.EffectProfileID

func init() {
	fileJson, _ := os.Open("data/list_audio_profile.json")
	json.NewDecoder(fileJson).Decode(&EffectAudio)
}
