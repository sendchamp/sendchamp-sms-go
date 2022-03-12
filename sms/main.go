package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	url := "https://api.sendchamp.com/api/v1/sms/send"

	type SendSMSPayload struct {
		To         []string `json:"to"`
		Message    string   `json:"message"`
		SenderName string   `json:"sender_name"`
		Route      string   `json:"route"`
	}

	sp := SendSMSPayload{
		To:         []string{"2348188469520"},
		Message:    "Lorem ipsum d, no lele",
		SenderName: "Dash",
		Route:      "non_dnd",
	}

	j, _ := json.Marshal(sp)
	payload := bytes.NewReader(j)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer ACCESS_KEY")

	res, e := http.DefaultClient.Do(req)

	if e != nil {
		log.Fatal(e)
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
