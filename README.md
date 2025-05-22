# Google Cloud Text-To-Speech

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