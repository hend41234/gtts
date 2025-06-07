package generatetts

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
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
	// chek input apiKey
	if len(apiKey) > 0 {
		// key = apiKey[0]
		// the apiKey is used
		// use apiKey for API_KEY

		// create new Utils
		utilstts.Utils = new(utilstts.UtilizationModel)
		// set API_KEY
		utilstts.Utils.API_KEY = apiKey[0]
		// set BaseURL
		utilstts.Utils.BaseURL = "https://texttospeech.googleapis.com"

	} else {
		// if apiKey is nil / not used, then check the content Utils
		if utilstts.Utils == nil {
			// LoadConf() not used
			log.Fatal("API_KEY not found\nplease set :\n\tutilstts.LoadConf('envFile')")
		}
	}

	if Config.Input.Text != "" {

		// check, for input what is reference file txt or not
		if !utilopentxt.CheckTXTFile(Config.Input.Text) {
			// open file txt
			openText := utilopentxt.ReadTXT(Config.Input.Text)

			// setting content the file txt to COnfig.Input.Text
			Config.Input.Text = openText

			//  print the content
			fmt.Println(Config.Input.Text)
		}
		fmt.Println(Config.Input)
	}

	if Config.Input.SSML != "" {
		if Config.LowLatencyJourneySynthesis.LowLatencyJourneySynthesis {
			log.Fatal("please check your config, make sure when used Chirp3, Wavenet not using SSML")
		}
		// fmt.Println(Config.LowLatencyJourneySynthesis.LowLatencyJourneySynthesis)
		if !utilopenxml.CheckXMLFile(Config.Input.SSML) {
			fmt.Println("ssml ada")
			openXml := utilopenxml.ReadXML(Config.Input.SSML)
			Config.Input.SSML = openXml
			fmt.Println(Config.Input.SSML)
		}
		fmt.Println(Config.Input.SSML)
	}

	endpoint := "/v1/text:synthesize"

	url := fmt.Sprintf("%v%v?key=%v", utilstts.Utils.BaseURL, endpoint, utilstts.Utils.API_KEY)

	// var buf bytes.Buffer
	// {
	// 	enc := json.NewEncoder(&buf)
	// 	enc.SetEscapeHTML(false)
	// 	err := enc.Encode(Config)
	// 	if err != nil {
	// 		log.Fatal("error create buffer body")
	// 	}
	// }

	body, _ := json.Marshal(Config)
	{
		fmt.Println(string(body))
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		var resErr models.ResError
		err := json.NewDecoder(res.Body).Decode(&resErr)
		if err != nil {
			fmt.Println(err)
			log.Fatal("error : decode response error")
		}
		log.Fatal(resErr.Error.Message)

	}
	err = json.NewDecoder(res.Body).Decode(&NewAudio)
	if err != nil {
		log.Fatal(err)
	}
	// return
}

func SaveAudio(nameFile string, audioBuffer string, ext ...string) bool {
	if len(ext) > 0 {
		nameFile = fmt.Sprintf("%v.%v", nameFile, ext[0])
	} else {
		nameFile = fmt.Sprintf("%v.%v", nameFile, strings.ToLower(Config.AudioConfig.AudioEncoding))
	}
	if dir := filepath.Dir(nameFile); utilopentxt.CheckTXTFile(dir) {
		// fmt.Println("ada")
		err := os.Mkdir(dir, 0755)
		if err != nil {
			nameFile = filepath.Base(nameFile)
		}
	}

	err := os.WriteFile(nameFile, []byte(audioBuffer), 0755)
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("saved in " + nameFile)
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
