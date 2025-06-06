package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	effectprofileid "github.com/hend41234/gtts/effectprofileid"
	"github.com/hend41234/gtts/models"
	"github.com/hend41234/gtts/text"
	utilhelp "github.com/hend41234/gtts/utilstts/help"
	utilopentxt "github.com/hend41234/gtts/utilstts/opentxt"
	utilopenxml "github.com/hend41234/gtts/utilstts/openxml"
)

var (
	ListEffect   bool
	ListEncoding bool
	ListHz       bool
	ListLC       bool
	ListLCName   bool
	Help         bool

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
	HelpUsage = "-h"
	HelpName  = "help"
	HelpShort = "h"

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
	LTUsage = "-lt  \nif u use the flag, this mean lowlatency is true"
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
	flag.BoolVar(&Help, HelpName, false, HelpUsage)
	flag.BoolVar(&Help, HelpShort, false, HelpUsage)

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

	flag.BoolVar(&LowLatency, LTName, false, LTUsage)
	flag.BoolVar(&LowLatency, LTShort, false, LTUsage)

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

func CLI() {
	var helper = map[string]func(){
		"help":          flag.PrintDefaults,
		"list-lc":       utilhelp.HelpListLanguageCode,
		"list-lc-name":  utilhelp.HelpListLanguageCodeName,
		"list-effect":   utilhelp.HelpListEffect,
		"list-encoding": utilhelp.HelpListEncoding,
		"list-hz":       utilhelp.HelpListHz,
	}

	TTSConfig()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		return
	} else {
		flag.VisitAll(func(f *flag.Flag) {
			if fn, ok := helper[f.Name]; ok {
				if f.Value.String() == "true" {
					fn()
					os.Exit(0)
				}
			}
		})
	}

	//  default
	// text.Synthesize(
	// models.SynthesizeInputModel{Text: "hi, there. my name is afrizal"},
	// text.DefaultVoiceBody(),
	// text.DefaultAudioConf(),
	// false,
	// )
	// util_help.HelpListLanguageCode()

	// [*] set input text/ssml

	input := models.SynthesizeInputModel{}
	if flag.Lookup("text").Value.String() == "" {
		//  setting when input is ssml
		ssmlInput := flag.Lookup("ssml")
		{
			// check file xml
			_, err := os.Stat(ssmlInput.Value.String())
			if notExist := os.IsNotExist(err); notExist {
				input.SSML = ssmlInput.Value.String()
				// return
			} else {
				xmlString := utilopenxml.ReadXML(ssmlInput.Value.String())
				input.SSML = xmlString
				// return
			}
		}
	} else {
		// setting when input is text
		textInput := flag.Lookup("t")
		{
			_, err := os.Stat(textInput.Value.String())
			if os.IsNotExist(err) {
				input.Text = textInput.Value.String()
				// return
			} else {
				readTxt := utilopentxt.ReadTXT(textInput.Value.String())
				input.Text = readTxt
				// return

			}
		}
	}

	// return
	// [*] set voice model
	voiceBody := text.DefaultVoiceBody()
	voiceBody.LanguageCode = flag.Lookup("vmlc").Value.String()
	voiceBody.Name = flag.Lookup("vmn").Value.String()
	voiceBody.SSMLGender = flag.Lookup("vmg").Value.String()

	//  [*] set audio config
	audioConf := text.DefaultAudioConf()
	audioConf.AudioEncoding = flag.Lookup("ace").Value.String()

	// eff, _ := strconv.Atoi(flag.Lookup("aceff").Value.String())
	audioConf.EffectsProfileId = effectprofileid.EffectAudio.AudioProfile[ACEffectsProfileId]

	pitch, _ := strconv.ParseFloat(flag.Lookup("acp").Value.String(), 64)
	audioConf.Pitch = pitch

	hertz, _ := strconv.ParseInt(flag.Lookup("achz").Value.String(), 10, 64)
	audioConf.SampleRateHertz = hertz

	speakingRate, _ := strconv.ParseFloat(flag.Lookup("acsr").Value.String(), 64)
	audioConf.SpeakingRate = speakingRate

	volGain, _ := strconv.ParseFloat(flag.Lookup("acvg").Value.String(), 64)
	audioConf.VolumeGainDb = volGain

	// [*] low latency

	text.Synthesize(input, voiceBody, audioConf, LowLatency)

}

func main() {
	// using CLI
	CLI()

	// using library
	// utilstts.LoadEnv(".env")
	// name := "ur-IN-Chirp3-HD-Vindemiatrix"
	// if newConfErr := generatetts.GenerateDefaultConfig(name); newConfErr != nil {
	// log.Println(newConfErr)
	// }
	// generatetts.Config.Input = models.SynthesizeInputModel{Text: "sample.txt"}
	// generatetts.RunGenerateTTS()
	// audioBuff, _ := base64.StdEncoding.DecodeString(generatetts.NewAudio.AudioContent)
	// generatetts.SaveAudio("output/mytest", string(audioBuff))
}
