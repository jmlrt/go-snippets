// Source: https://stackoverflow.com/questions/13263492/set-useragent-in-http-request

package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const url = "http://httpbin.org/user-agent"
const useragent = "Go-spider-bot/3.0"

func main() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", useragent)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))
}
