package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close() // close the response body to prevent resource leaks

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("response body: ", string(body))
}
