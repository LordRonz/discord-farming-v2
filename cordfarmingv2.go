package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"bytes"
)

func main() {
	fmt.Println("kontol world")
	url := "DISCORD_URL"
    fmt.Println("URL:>", url)

	ua := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.131 Safari/537.36"

	token := "YOUR_TOKEN"

    var jsonStr = []byte(`{
		"content":"Buy cheese and bread for breakfast."
	}`)
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
    req.Header.Set("user-agent", ua)
    req.Header.Set("Content-Type", "application/json")
	req.Header.Set("authorization", token)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}