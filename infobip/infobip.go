package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func main() {

	url := "https://8gek51.api.infobip.com/email/3/send"
	method := "POST"

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	_ = writer.WriteField("from", "tiar@wacaku.com")
	_ = writer.WriteField("subject", "Free trial")
	_ = writer.WriteField("to", "{\"to\":\"tiara.agisti@enigmacamp.com\",\"placeholders\":{\"firstName\":\"Tiar\"}}")
	_ = writer.WriteField("text", "Hi {{firstName}}, this is a test message from Infobip. Have a nice day!")
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "App ")
	req.Header.Add("Content-Type", "multipart/form-data")
	req.Header.Add("Accept", "application/json")

	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
