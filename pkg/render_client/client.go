package render_client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type RenderClient struct {
	Host string
	Port int
}

func New(host string, port int) *RenderClient {
	return &RenderClient{Host: host, Port: port}
}

func (rc *RenderClient) Render(html string) ([]byte, error) {
	json, err := json.Marshal(Request{Query: html})
	if err != nil {
		return nil, fmt.Errorf("cant marshall docx-gen req to json: %w", err)
	}

	request, err := http.NewRequest("POST", "http://"+rc.Host+":"+strconv.Itoa(rc.Port)+"/render",
		bytes.NewBuffer(json))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, fmt.Errorf("cant get html-to-image js service response: %w", err)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)

	return body, nil
}
