package generatetts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/hend41234/gtts/models"
	"github.com/hend41234/gtts/utilstts"
	utilopentxt "github.com/hend41234/gtts/utilstts/opentxt"
	utilopenxml "github.com/hend41234/gtts/utilstts/openxml"
	"github.com/hend41234/gtts/voices"
)

var Config *models.TextBaseModel
var NewAudio *models.ResSynthesize
var ApiKey *string

func RunGenerateTTS(apiKey ...string) {
	var key string
	if len(apiKey) > 0 {
		key = apiKey[0]
	} else {
		if utilstts.Utils == nil {
			log.Fatal("API_KEY not found\nplease set :\n\tutilstts.LoadConf('envFile')")
		}
	}
	key = utilstts.Utils.API_KEY

	if Config.Input.Text != "" {
		if utilopentxt.CheckTXTFile(Config.Input.Text) {
			openText := utilopentxt.ReadTXT(Config.Input.Text)
			Config.Input.Text = openText
			fmt.Println(Config.Input.Text)
		}
	}

	if Config.Input.SSML != "" {
		if Config.LowLatencyJourneySynthesis.LowLatencyJourneySynthesis {
			log.Fatal("please check your config, make sure when used Chirp3, Wavenet not using SSML")
		}
		fmt.Println(Config.LowLatencyJourneySynthesis.LowLatencyJourneySynthesis)
		if utilopenxml.CheckXMLFile(Config.Input.SSML) {
			openXml := utilopenxml.ReadXML(Config.Input.SSML)
			Config.Input.SSML = openXml
		}
	}
	endpoint := "/v1/text:synthesize"
	url := fmt.Sprintf("%v%v?key=%v", utilstts.Utils.BaseURL, endpoint, key)

	body, err := json.Marshal(Config)
	if err != nil {
		log.Fatal(err)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&NewAudio)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func SaveAudio(nameFile string, audioBuffer string, ext ...string) bool {
	if len(ext) > 0 {
		nameFile = fmt.Sprintf("%v.%v", nameFile, ext[0])
	} else {
		nameFile = fmt.Sprintf("%v.%v", nameFile, strings.ToLower(Config.AudioConfig.AudioEncoding))
	}

	err := os.WriteFile(nameFile, []byte(audioBuffer), 0755)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func containsVoices(text string) bool {
	for _, voice := range voices.ListVoices.Voice {
		if text == voice.Name {
			return true
		}
	}
	return false
}
