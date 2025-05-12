package text

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gctts/models"
	"gctts/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

type ResSynthesize struct {
	AudioContent string `json:"audioContent"`
}

func Synthesize() {
	sampleText := "Hi there, im a afrizal. im a handsome guy. im have a big dick."
	endpoint := "/v1/text:synthesize"
	url := fmt.Sprintf("%v%v?key=%v", utils.Utils.BaseURL, endpoint, utils.Utils.API_KEY)

	inputBody := models.SynthesizeInputModel{
		Text: sampleText,
	}
	voiceBody := models.VoiceSelectionParamsModel{
		LanguageCode: "en-US",
		Name:         "en-US-Neural2-D",
		SSMLGender:   "MALE",
	}
	audioConfigBody := models.AudioConfigModel{
		AudioEncoding:    "MP3",
		SpeakingRate:     1.0,
		Pitch:            0.0,
		VolumeGainDb:     0.0,
		EffectsProfileId: []string{"telephony-class-application"},
	}
	body := models.TextBaseModel{
		Input:       inputBody,
		Voice:       voiceBody,
		AudioConfig: audioConfigBody,
	}
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

	err = ioutil.WriteFile("output/result.mp3", audioData, 0755)
	if err != nil {
		log.Fatal("error save result.mp3")
	}
	// rs, _ := io.ReadAll(resp.Body)
	// fmt.Println(string(Res.AudioContent))
	fmt.Println("success")
}
