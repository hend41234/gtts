package main

import (
	"flag"
	"fmt"
)

var (
	ListEffect   bool
	ListEncoding bool
	ListHz       bool
	ListLC       bool
	ListLCName   bool

	// text regular
	//
	Text string
	// text xml / ssml
	//
	SSML string
	// bool
	//
	LowLatency bool

	// see in -list-encoding
	//
	ACAudioEncoding string
	// default is 1.0
	//
	// speaking rate in [0.25 - 2.0]
	//
	ACSpeakingRate float64
	// default is 0.0
	//
	// pitch rate in [-20.0 - 20.0]
	ACPitch float64
	// default is 0.0
	//
	// volume gain db range [-96.0 - 16.0]
	ACVolumeGainDb float64
	// see in '-list-effect'
	//
	ACEffectsProfileId int64
	// see in '-list-hz'
	//
	ACSampleRateHertz int64

	// see in '-list-lc
	//
	VMLanguageCode string
	// see in -list-lc-name'
	//
	VMName string
	// MALE or FEMALE
	//
	VMSsmlGender string
)

const (
	ListEfectUsage  = "-list-effect / -lef \n-- it will show list of effects profile id"
	ListEffectName  = "list-effect"
	ListEffectShort = "lef"

	ListEncodingUsage = "-list-encoding / -len \n-- it will show list of encoding audio result"
	ListEncodingName  = "list-encoding"
	ListEncodingShort = "len"

	ListHzUsage = "-list-hz \n-- it will show what valid of rate hertz "
	ListHzName  = "list-hz"
	ListHzShort = "lhz"

	ListLcUsage = "-list-lc \n-- it will show list of language code"
	ListLCNames = "list-lc"
	ListLCShort = "llc"

	ListLCNameUsage = "-list-lc-name \n-- it will show list name of list name of language code "
	ListLCNameName  = "list-lc-name"
	ListLCNameShort = "llcn"
	// ================================================================================================
	// [*] input text to speech
	TextUsage = "-t 'hi there, my name is john' \n-- use regular text"
	TextName  = "text"
	TextShort = "t"

	SSMLUsage = " -ssml name_file.xml / -ssml '<speak>hi there</speak>' \n-- you can reference it to a .xml file or directly input the ssml code."
	SSMLName  = "ssml"
	SSMLShort = "ssml"

	// [*] low latency
	LTUsage = "-lt false / -lt true \n"
	LTName  = "low-latency"
	LTShort = "lt"

	// [*] audio config
	ACAudioEncodingUsage = "-ace 'MP3' \n-- see in -list-encoding"
	ACAudioEncodingName  = "ac-encoding"
	ACAudioEncodingShort = "ace"

	ACSpeakingRateUsage = "-acsr 0.0 \n-- range speaking rate usage [ 0.25 - 2.0]"
	ACSpeakingRateName  = "ac-speaking-rate"
	ACSpeakingRateShort = "acsr"

	ACPitchUsage = "-acp 0.0 \n-- pitch range [ -20.0 - 20.0 ]"
	ACPitchName  = "ac-pitch"
	ACPitchShort = "acp"

	ACVolumeGainDbUsage = "-acvg 0.0 \n-- volume gain Db range [ -96.0 - 16.0 ]"
	ACVolumeGainDbName  = "ac-volume-gain"
	ACVolumeGainDbShort = "acvg"

	ACEffectsProfileIdUsage = "-aceff 2 \n-- see in '-list-effect'"
	ACEffectsProfileIdName  = "ac-effect-profile"
	ACEffectsProfileIdShort = "aceff"

	ACSampleRateHertzUsage = "-achz 24000 \n-- default 24000, see more in '-list-hz'"
	ACSampleRateHertzName  = "ac-rate-hertz"
	ACSampleRateHertzShort = "achz"

	// [*] voice model
	VMLanguageCodeUsage = "-vmlc 'en-US' \n-- see in '-list-lc'"
	VMLanguageCodeName  = "vm-language-code"
	VMLanguageCodeShort = "vmlc"

	VMNameUsage = "-vmn 'en-US-Neural2-J'"
	VMNameName  = "vm-name"
	VMNameShort = "vmn"

	VMSsmlGenderUsage = "-vmg 'MALE' \n -- use MALE or FEMALE"
	VMSsmlGenderName  = "vm-gender"
	VMSsmlGenderShort = "vmg"
)

func TTSConfig() {
	// [*] show helping
	flag.BoolVar(&ListEncoding, ListEncodingName, false, ListEncodingUsage)
	flag.BoolVar(&ListEncoding, ListEncodingShort, false, ListEncodingUsage)

	flag.BoolVar(&ListLC, ListLCNames, false, ListLcUsage)
	flag.BoolVar(&ListLC, ListLCShort, false, ListLcUsage)

	flag.BoolVar(&ListLCName, ListLCNameName, false, ListLCNameUsage)
	flag.BoolVar(&ListLCName, ListLCNameShort, false, ListLCNameUsage)

	flag.BoolVar(&ListHz, ListHzName, false, ListHzUsage)
	flag.BoolVar(&ListHz, ListHzShort, false, ListHzUsage)

	flag.BoolVar(&ListEffect, ListEffectName, false, ListEfectUsage)
	flag.BoolVar(&ListEffect, ListEffectShort, false, ListEfectUsage)
	// [*] require input
	flag.StringVar(&Text, TextName, "", TextUsage)
	flag.StringVar(&Text, TextShort, "", TextUsage)

	flag.StringVar(&SSML, SSMLName, "", SSMLUsage)

	// [*] audio config
	flag.StringVar(&ACAudioEncoding, ACAudioEncodingName, "MP3", ACAudioEncodingUsage)
	flag.StringVar(&ACAudioEncoding, ACAudioEncodingShort, "MP3", ACAudioEncodingUsage)

	flag.Float64Var(&ACSpeakingRate, ACSpeakingRateName, 1.0, ACSpeakingRateUsage)
	flag.Float64Var(&ACSpeakingRate, ACSpeakingRateShort, 1.0, ACSpeakingRateUsage)

	flag.Float64Var(&ACPitch, ACPitchName, 0.0, ACPitchUsage)
	flag.Float64Var(&ACPitch, ACPitchShort, 0.0, ACPitchUsage)

	flag.Float64Var(&ACVolumeGainDb, ACVolumeGainDbName, 0.0, ACVolumeGainDbUsage)
	flag.Float64Var(&ACVolumeGainDb, ACVolumeGainDbShort, 0.0, ACVolumeGainDbUsage)

	flag.Int64Var(&ACEffectsProfileId, ACEffectsProfileIdName, 0, ACEffectsProfileIdUsage)
	flag.Int64Var(&ACEffectsProfileId, ACEffectsProfileIdShort, 0, ACEffectsProfileIdUsage)

	flag.Int64Var(&ACSampleRateHertz, ACSampleRateHertzName, 24000, ACSampleRateHertzUsage)
	flag.Int64Var(&ACSampleRateHertz, ACSampleRateHertzShort, 24000, ACSampleRateHertzUsage)

	// [*voice model]
	flag.StringVar(&VMLanguageCode, VMLanguageCodeName, "en-US", VMLanguageCodeUsage)
	flag.StringVar(&VMLanguageCode, VMLanguageCodeShort, "en-US", VMLanguageCodeUsage)

	flag.StringVar(&VMName, VMNameName, "en-US-Neural2-J", VMNameUsage)
	flag.StringVar(&VMName, VMNameShort, "en-US-Neural2-J", VMNameUsage)

	flag.StringVar(&VMSsmlGender, VMSsmlGenderName, "MALE", VMSsmlGenderUsage)
	flag.StringVar(&VMSsmlGender, VMSsmlGenderShort, "MALE", VMSsmlGenderUsage)

	flag.Parse()

	if flag.NFlag() == 0 {
		fmt.Print(`example usage:
[*] gtts -t "hello there, this is google text to speech" -ace "MP3"
[*] gtts -ssml "<speak>hello there, this is google text to speech</speak>"
[*] gtts -ssml file_ssml.xml
`)

	}
}

func main() {
	//  default
	// text.Synthesize(
	// models.SynthesizeInputModel{Text: "hi, there. my name is afrizal"},
	// text.DefaultVoiceBody(),
	// text.DefaultAudioConf(),
	// false,
	// )
	// util_help.HelpListLanguageCode()
	TTSConfig()
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Println(f.Name, "", f.Value)
	})

}
