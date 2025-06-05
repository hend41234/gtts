# Google Cloud Text-To-Speech

## Using CLI
* create the .env fle
    
    ```bash
    API_KEY="your-api-key-from-google"
    BASE_URL="https://texttospeech.googleapis.com"
    ```
    check on [google console](https://console.cloud.google.com/apis/credentials) to get `API KEY`.
* run the program

    ### example :

    * 1 
        ```shell
        go run main.go -t some_textfile.txt
        ```
    * 2
        ```shell
        go run main.go -t "hello there, this is google cloud text to speech"
        ```
    * 3
        ```shell
        go run main.go -ssml some_xmlfile.xml
        ```
    * 4
        ```shell
        go run main.go -ssml "<speak>hello there, this is google cloud text to speech</speak>"
        ```
    * 5
        ```shell
        go run main.go -t some_textfile.txt -vmlc "en-US" -vmn "es-US-Chirp3-HD-Zubenelgenubi" -lt
        ```
    * see more 

    ```shell
    go run main.go -h
    ``` 

## simple note
if you want latency is ON / true. the voice model `can only use Chirp3`. like `example 5`.

## Using in your code
* ## install
    ```bash
        go get "github.com/hend42134/gtts"
    ```
* ## generate Text-to-Speech according with name voice
    To create a voice easily and according to the name of the voice, we can use the code as below

    ```go
        import (
            "github.com/hend41234/gtts"
            "github.com/hend41234/gtts/generatetts"
        )

        // the Name Voice
        name := "ur-IN-Chirp3-HD-Vindemiatrix"

        // generate default config according Name Voice
	    if newConfErr := generatetts.GenerateDefaultConfig(name); newConfErr != nil {
            log.Println(newConfErr)
	    }

        // input the text or nmae file which contains text 
	    generatetts.Config.Input = models.SynthesizeInputModel{Text: "sample.txt"}

        // running request, the response saved in generatetts.NewAudio variable
	    generatetts.RunGenerateTTS()

        // because response from server is base64 encoding, we must convert the response to buffer string
	    audioBuff, _ := base64.StdEncoding.DecodeString(generatetts.NewAudio.AudioContent)

        // save audio using audioBuff, that we have prepared
	    generatetts.SaveAudio("output/bandol", string(audioBuff))
    ```

* ## see all the voices

    If you don't know what voices you can use, we can see them all using this code.

    ```golang
    import (
        "github.com/hend41234/voices"
        "fmt"
        )

    for _, list := range voices.ListVoices{
        fmt.Println(list.Name)

        // or if you want to see detail, uncomment code as below
        // fmt.Println(list.LanguageCodes[0])   
        // fmt.Println(list.SSMLGender)
        // fmt.Println(list.NaturalSampleRateHertz)
    }

    ```

