package utilstts

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EndpointsModel struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type UtilizationModel struct {
	API_KEY string `json:"API_KEY"`
	BaseURL string `json:"BASE_URL"`
	// Endpoint []EndpointsModel `json:"ENDPOINT"`
}

var Utils *UtilizationModel

// require to use, you must have the API KEY of gemini, the key is API_KEY in the env file
//
// if you do not have the API KEY, you can open the url : https://console.cloud.google.com/apis/credentials
func LoadEnv(envFile string) {
	Utils = new(UtilizationModel)
	for {
		envErr := godotenv.Load(envFile)
		if envErr != nil {
			log.Fatal("API_KEY not found, or env file nor found")
			// time.Sleep(2 * time.Second)
			// continue
		}
		break
	}
	godotenv.Load(".env")

	Utils.API_KEY = os.Getenv("API_KEY")
	// Utils.BaseURL = os.Getenv("BASE_URL")
	Utils.BaseURL = "https://texttospeech.googleapis.com"

}
