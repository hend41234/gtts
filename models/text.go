package models

type SynthesizeInputModel struct {
	// CustomPronunciations string
	Text string `json:"text,omitempty"`
	SSML string `json:"ssml,omitempty"`
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
	SampleRateHertz  int32    `json:"smapleRateHertz,omitempty"`
}

type AdvanceConfModel struct {
	LowLatencyJourneySynthesis bool `json:"lowLatencyJourneySynthesis"`
}

type TextBaseModel struct {
	Input                      SynthesizeInputModel      `json:"input"`
	Voice                      VoiceSelectionParamsModel `json:"voice"`
	AudioConfig                AudioConfigModel          `json:"audioConfig"`
	LowLatencyJourneySynthesis AdvanceConfModel          `json:"advancedVoiceOptions,omitempty"`

	// AdvancedVoiceOptions string
}
