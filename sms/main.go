package main

import (
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"
)

func main() {

	url := "https://api.sendchamp.com//api/v1/sms/send"

	payload := strings.NewReader("{
    "to": [
        "2348188469520"
    ],
    "message": "Lorem ipsum d, no lele",
    "sender_name": "Dash",
    "route": "non_dnd"
}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer ACCESS_KEY")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
