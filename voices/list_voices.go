package voices

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/hend41234/gtts/models"
	"github.com/hend41234/gtts/utilstts"
)

var ListVoices *models.ListVoicesModel
var wg = sync.WaitGroup{}

func init() {
	wg.Add(1)
	defer wg.Done()
	ListVoices = new(models.ListVoicesModel)
	if _, err := os.Stat("data/list_voices.json"); os.IsNotExist(err) {
		GetListVoices()
	}
	file, _ := os.Open("data/list_voices.json")
	defer file.Close()
	byteFile, _ := io.ReadAll(file)
	err := json.Unmarshal(byteFile, &ListVoices)
	if err != nil {
		log.Println("error load list_voices.json")
	}
}

func GetListVoices() {
	endpoint := "/v1/voices"
	url := fmt.Sprintf("%v%v?key=%v", utilstts.Utils.BaseURL, endpoint, utilstts.Utils.API_KEY)
	req, _ := http.NewRequest("GET", url, nil)
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	readResp, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Println(resp.StatusCode)
		fmt.Println(string(readResp))
		fmt.Printf("Failed")
		return
	}
	// fmt.Println(string(readResp))
	err = os.WriteFile("data/list_voices.json", readResp, 0755)
	if err != nil {
		log.Fatal("error write list_voices.json")
	}
	log.Println("get list voices success")
}
