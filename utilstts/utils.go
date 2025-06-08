package utilstts

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EndpointsModel struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type UtilizationModel struct {
	API_KEY string `json:"GC_API_KEY"`
	BaseURL string `json:"BASE_URL"`
	// Endpoint []EndpointsModel `json:"ENDPOINT"`
}

var Utils *UtilizationModel

// require to use, you must have the API KEY of gemini, the key is GC_API_KEY in the env file
//
// if you do not have the API KEY, you can open the url : https://console.cloud.google.com/apis/credentials
func LoadEnv(envFile string) {
	Utils = new(UtilizationModel)
	for {
		envErr := godotenv.Load(envFile)
		if envErr != nil {
			log.Fatal("GC_API_KEY not found, or env file nor found")
			// time.Sleep(2 * time.Second)
			continue
		}
		break
	}
	godotenv.Load(".env")

	Utils.API_KEY = os.Getenv("GC_API_KEY")
	// Utils.BaseURL = os.Getenv("BASE_URL")
	Utils.BaseURL = "https://texttospeech.googleapis.com"
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(".env loaded")
		return
	}
	Utils = new(UtilizationModel)
	gcApiKey := os.Getenv("GC_API_KEY")
	if gcApiKey != "" {
		Utils.API_KEY = gcApiKey
	}
	Utils.BaseURL = "https://texttospeech.googleapis.com"
	// return

}
