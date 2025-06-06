package models

type VoicesModel struct {
	LanguageCodes          []string `json:"languageCodes"`
	Name                   string   `json:"name"`
	SSMLGender             string   `json:"ssmlGender"`
	NaturalSampleRateHertz int      `json:"naturalSampleRateHertz"`
}


type ListVoicesModel struct {
	Voice []VoicesModel `json:"voices"`
}
