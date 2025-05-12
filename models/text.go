package models

type SynthesizeInputModel struct {
	// CustomPronunciations string
	Text string `json:"text"`
	// SSML                 string
	// NuktiSpeakerMarkup   string
}
type VoiceSelectionParamsModel struct {
	LanguageCode string `json:"languageCode"`
	Name         string `json:"name"`
	SSMLGender   string `json:"ssmlGender"`
	// CustomVoice
	// VoiceClone
}

type AudioConfigModel struct {
	AudioEncoding    string   `json:"audioEncoding"`
	SpeakingRate     float32  `json:"speakingRate"`
	Pitch            float32  `json:"pitch"`
	VolumeGainDb     float32  `json:"volumeGainDb"`
	EffectsProfileId []string `json:"effectsProfileId"`
}

type TextBaseModel struct {
	Input       SynthesizeInputModel      `json:"input"`
	Voice       VoiceSelectionParamsModel `json:"voice"`
	AudioConfig AudioConfigModel          `json:"audioConfig"`
	// AdvancedVoiceOptions string
}
