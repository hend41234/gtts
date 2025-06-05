package text

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hend41234/gtts/models"
	"github.com/hend41234/gtts/utilstts"
)

// <break time="300ms"/>
// </break time="300ms"/>
var x = `<speak>
  <voice name="id-ID-ArdiNeural">
    <prosody rate="medium" pitch="default">
      Hari ini, <break time="300ms"/> saya akan menceritakan sebuah kisah menarik.
      <break time="500ms"/>
      <emphasis level="moderate">Bayangkan</emphasis>, Anda sedang berada di tengah hutan lebat...
      <break time="700ms"/>
      Suara angin berdesir... <break time="300ms"/> dan dedaunan bergoyang pelan.
      <prosody pitch="+10%">Tiba-tiba,</prosody> <break time="200ms"/> terdengar langkah kaki mendekat...
    </prosody>
  </voice>
</speak>`

// var testInput = `Dulu... manusia berjalan berhari-hari hanya untuk bertukar pesan, kini? Cukup dengan satu klik—dunia pun menjawab! Ironisnya, di tengah teknologi yang merapatkan jarak, hati-hati kita justru saling menjauh... Mengapa? Karena layar yang terang itu, perlahan memudarkan tatapan mata yang tulus. Sejarah mencatat perang demi tanah, kekuasaan, dan kini—data. Di balik setiap sel tubuh, ada miliaran proses biologis yang bekerja tanpa kita sadari, tapi apa gunanya jika kita sendiri lupa bagaimana cara hidup yang benar? Ah, hidup ini... bukan sekadar bertahan, tapi mengerti, merasakan, dan... berani berubah!`
// var testInput1 = `Life... it's unpredictable, isn't it? One moment you're scrolling through endless feeds, consumed by pixels and posts, and the next — you realize you've forgotten the warmth of real conversation, the scent of rain on earth, or the joy of silence. Technology races forward, faster than our emotions can process; we connect more, yet feel less. Did you know that despite all our advancements, the human brain — evolved over millions of years — still can't distinguish between digital stress and real danger? Fascinating, right? History repeats, biology adapts, society transforms... but are we truly evolving, or just upgrading? Think about it!`
// var s = `<speak>In a world shaped by circuits and signals, <break time="300ms"/> technology has become the silent architect of our daily lives. <break time="200ms"/> From the hum of ancient machinery in the Industrial Revolution, <break time="150ms"/> to the near-whisper of quantum processors today, <emphasis level="moderate">we have always pursued progress</emphasis>. But amidst this digital evolution, <break time="250ms"/> have we paused to consider the cost to our social fabric? <break time="300ms"/> Biologically, the human brain still craves touch, voice, presence... <break time="200ms"/> not just data. <prosody pitch="+5%">So,</prosody> as we advance into the unknown, <break time="150ms"/> let us not forget: <prosody rate="slow">true connection</prosody> lies not only in bandwidth — but in empathy.</speak>`
// var s = `<speak>In the grand tapestry of human history, <break time="300ms"/> technology stands as both a marvel and a mirror. <break time="500ms"/> From the crackling spark of the first fire <emphasis level="moderate">to the silent hum of artificial intelligence</emphasis>, each invention reflects our deepest desires — to connect, to create, to control. <break time="300ms"/> But as we build machines that think, <emphasis level="strong">do we risk forgetting how to feel?</emphasis> <break time="400ms"/> Progress whispers promises of convenience... yet often, it steals the presence of the moment. <break time="300ms"/> Still, we push forward — curious, relentless, <prosody rate="slow">hoping that somewhere in the code</prosody>, we rediscover the soul.</speak>`
// var s = `<speak>Life is not a straight line... it's a winding path filled with unexpected turns, quiet pauses, and sudden storms. <break time="300ms"/> We chase dreams through cities of noise, yet often find truth in moments of stillness. <break time="200ms"/> Each heartbeat carries a story — of survival, love, failure, and resilience. <break time="500ms"/> In a world that demands speed, the greatest act of courage... may simply be to slow down. <emphasis level="moderate">To listen. To breathe. To be.</emphasis></speak>`
var s = `<speak>
Life, as we know it, is rarely fair... <break time="400ms"/> It's a journey filled with contrast — joy and sorrow, success and failure, hope and despair. <break time="300ms"/> Yet, it's in this imbalance that we find meaning. <break time="400ms"/> We chase dreams, lose people, discover truths — sometimes too late. <break time="300ms"/> But that's the essence of being human, isn't it? <break time="400ms"/> Not to perfect the path, but to walk it, stumble, learn, and grow... <break time="300ms"/> So tell me — are you living, or just surviving?
</speak>
`
var ss = `<speak>
Technology surrounds us every day... <break time="300ms"/> It shapes the way we communicate, work, and even think. <break time="400ms"/> But amidst this rapid progress, have we stopped to ask — are we mastering technology, or is it mastering us? <break time="300ms"/> With every innovation, new possibilities emerge, yet so do new challenges. <break time="400ms"/> The future is exciting, uncertain, and full of questions — how will we adapt? <break time="300ms"/> Only time will tell.
</speak>
`
var sss = `<speak>
Technology has transformed the way we live, connecting the world in ways once unimaginable. <break time="400ms"/> From the first telegraph to today's lightning-fast internet, each breakthrough has reshaped society, opening doors to new possibilities and challenges. <break time="300ms"/> But with every innovation comes a question: are we mastering technology, or is it mastering us? <break time="400ms"/> The answer lies not just in the machines we build, but in how we choose to use them. <break time="500ms"/>

Social media, for instance, has revolutionized communication, yet it also blurs the line between genuine connection and curated reality. <break time="400ms"/> We live in an era where virtual friendships flourish, but loneliness quietly grows. <break time="300ms"/> In this paradox, we must ask ourselves—how do we balance progress with presence? <break time="400ms"/> Because ultimately, technology should enhance our humanity, not replace it. </speak>
`
var ssss = `<speak>
  Hey there! <break time="300ms"/> 
  Today, we're gonna talk about something <emphasis level="moderate">really</emphasis> interesting. <break time="500ms"/>
  
  Have you ever wondered <prosody rate="slow">how artificial intelligence is changing our lives?</prosody> <break time="300ms"/>
  
  Well, <prosody pitch="+2st">you're not alone.</prosody> <break time="300ms"/>
  
  Every day, AI is helping us in ways we often don't even notice — like recommending music, unlocking phones with our face, or even... <prosody volume="loud">talking to you right now!</prosody> <break time="700ms"/>
  
  <prosody rate="medium">So, let’s dive in and see what’s really going on behind the scenes.</prosody>
</speak>`
var sssss = `<speak>
  <prosody rate="medium">
    <p>
      <s>
        <emphasis level="moderate">Imagine this...</emphasis>
        A world where machines can see, listen... and even understand us. 
      </s>
      <break time="400ms"/>
      <s>
        That world? <prosody pitch="+1st">It's already here.</prosody>
      </s>
    </p>

    <p>
      <s>
        From recommending what song you’ll love next, to recognizing your voice in a crowded room...
      </s>
      <s>
        <emphasis level="strong">AI is quietly shaping everything.</emphasis>
      </s>
      <break time="600ms"/>
      <s>
        But here's the question: <break time="200ms"/> are we really ready for what's coming?
      </s>
    </p>

    <p>
      <s>
        In this video, we’ll dive into how artificial intelligence is not just a tool — 
        <prosody pitch="-1st" rate="slow">but a force that's redefining human potential.</prosody>
      </s>
      <break time="400ms"/>
      <s>
        <emphasis level="moderate">Let's begin.</emphasis>
      </s>
    </p>
  </prosody>
</speak>`

// var ss = `<speak>Life... <break time="500ms"/> it moves quietly, like a whisper carried by the wind. <break time="400ms"/> One day you're dreaming beneath the stars, <break time="300ms"/> the next, you're caught in the rhythm of endless responsibilities. <break time="500ms"/> We chase success, love, meaning... yet often forget to breathe. <break time="300ms"/> Isn't it strange? <break time="300ms"/> In a world full of noise, silence feels louder than ever. <break time="400ms"/> But within that silence — <emphasis level="moderate">truth</emphasis> waits. <break time="300ms"/> The truth that real life isn’t always beautiful... <break time="200ms"/> but it is real, raw, and absolutely worth living. <break time="500ms"/> Always.</speak>`

// var s = `<speak><google:style name="empathetic">Hello I'm so happy today!</google:style></speak>`

type ResSynthesize struct {
	AudioContent string `json:"audioContent"`
}

func DefaultVoiceBody() (defaultConfig models.VoiceSelectionParamsModel) {
	defaultConfig.LanguageCode = "en-US"
	defaultConfig.Name = "en-US-Neural2-J"
	defaultConfig.SSMLGender = "MALE"
	return
}
func DefaultAudioConf() models.AudioConfigModel {
	return models.AudioConfigModel{
		AudioEncoding:   "MP3",
		SpeakingRate:    1.0,
		Pitch:           0.0,
		VolumeGainDb:    0.0,
		SampleRateHertz: 24000,
	}
}

var DefaultLowLatency = false

func Synthesize(inputBody models.SynthesizeInputModel, voiceBody models.VoiceSelectionParamsModel, audioConf models.AudioConfigModel, lowLtency bool) {
	endpoint := "/v1/text:synthesize"
	url := fmt.Sprintf("%v%v?key=%v", utilstts.Utils.BaseURL, endpoint, utilstts.Utils.API_KEY)
		
	// voiceBody := models.VoiceSelectionParamsModel{
	// LanguageCode: "en-US",
	// Name:         "en-US-Neural2-J",
	// SSMLGender:   "MALE",
	// }

	// audioConfigBody := models.AudioConfigModel{
	// AudioEncoding: "MP3",
	// SpeakingRate:  1.0,
	// Pitch:         0.0,
	// VolumeGainDb:  0.0,
	// EffectsProfileId: []string{"headphone-class-device"},
	// }
	body := models.TextBaseModel{
		Input:                      inputBody,
		Voice:                      voiceBody,
		AudioConfig:                audioConf,
		LowLatencyJourneySynthesis: models.AdvanceConfModel{LowLatencyJourneySynthesis: lowLtency},
	}

	{
		byteBody, _ := json.Marshal(body)

		req, _ := http.NewRequest("POST", url, bytes.NewBuffer(byteBody))
		client := http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			log.Fatal("error request synthesize")
		}
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			log.Println(resp.StatusCode)
			readResp, _ := io.ReadAll(resp.Body)
			fmt.Println(string(readResp))
			return
		}
		Res := ResSynthesize{}
		json.NewDecoder(resp.Body).Decode(&Res)
		audioData, _ := base64.StdEncoding.DecodeString(Res.AudioContent)
		resultName := "result." + strings.ToLower(audioConf.AudioEncoding)
		err = ioutil.WriteFile("output/"+resultName, audioData, 0755)
		if err != nil {
			log.Fatal("error save result.mp3")
		}
		// rs, _ := io.ReadAll(resp.Body)

		// fmt.Println(string(Res.AudioContent))
		fmt.Println("success")
		pwd, _ := os.Getwd()
		scs := fmt.Sprintf("saved in %v/output/%v", pwd, resultName)
		fmt.Println(scs)
	}

}
