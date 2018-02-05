package dialogflow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Answer struct {
	Result struct{
		Speech string `json:"speech"`
	} `json:"result"`
}


func (c *Client) Query(query string) string {

	var payload = []byte(fmt.Sprintf(`
	{
		"lang": "fr",
		"query": "%s",
		"sessionId": "12345"
	}
	`, query))

	req, _ := http.NewRequest("POST", c.Url, bytes.NewBuffer(payload))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var a Answer
	err = json.Unmarshal([]byte(body), &a)
	if err != nil {
		log.Println(err)
	}

	return a.Result.Speech
}
