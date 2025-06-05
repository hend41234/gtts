package utilstts

import (
	"os"
	"time"

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

func getEnv() {
	Utils = new(UtilizationModel)
	for {
		envErr := godotenv.Load(".env")
		if envErr != nil {
			time.Sleep(2 * time.Second)
			continue
		}
		break
	}
	godotenv.Load(".env")

	Utils.API_KEY = os.Getenv("API_KEY")
	Utils.BaseURL = os.Getenv("BASE_URL")

}

func init() {
	getEnv()
}
