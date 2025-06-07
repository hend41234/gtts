package generatetts

import (
	"fmt"

	"github.com/hend41234/gtts/models"
	"github.com/hend41234/gtts/text"
	"github.com/hend41234/gtts/voices"
)

func GenerateDefaultConfig(nameVoices ...string) error {
	Config = new(models.TextBaseModel)
	// setting voice body
	voiceBody := text.DefaultVoiceBody()
	var naturalHertz int
	{
		if len(nameVoices) > 0 {
			if !containsVoices(nameVoices[0]) {
				return fmt.Errorf("name voice does'nt exist")
			}
			for _, list := range voices.ListVoices.Voice {
				if list.Name == nameVoices[0] {

					voiceBody.LanguageCode = list.LanguageCodes[0]
					voiceBody.Name = list.Name
					voiceBody.SSMLGender = list.SSMLGender
					naturalHertz = list.NaturalSampleRateHertz
					break
				}
			}

		}

	}

	// setting audio config
	audioConfig := text.DefaultAudioConf()
	audioConfig.SampleRateHertz = int64(naturalHertz)
	audioConfig.EffectsProfileId = "headphone-class-device"
	lowLatency := text.DefaultLowLatency

	// build config Base Models
	baseModel := models.TextBaseModel{
		Voice:                      voiceBody,
		AudioConfig:                audioConfig,
		LowLatencyJourneySynthesis: models.AdvanceConfModel{lowLatency},
	}
	// NewModels = new(models.TextBaseModel)
	Config = &baseModel
	return nil
}
