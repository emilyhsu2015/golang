package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	testGet("f1")
	testPost()
}

func testGet(p string) {
	url := "http://127.0.0.1:1323/amount/" + p

	req, _ := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, _ := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	var result string

	json.Unmarshal(body, &result)

	fmt.Printf("%s\n", string(body[:]))

	defer resp.Body.Close()
}

func testPost() {
	url := "http://127.0.0.1:1323/basic"

	payload := strings.NewReader("name=cute&email=test@a.a")
	req, _ := http.NewRequest("POST", url, payload)

	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result string
	json.Unmarshal(body, &result)
	fmt.Printf("%s\n", string(body[:]))

}
