package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const base = "https://swapi.co/api/"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please provide a Ship ID")
		return
	}

	url := base + "starships/" + os.Args[1]

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	if code := response.StatusCode; code != http.StatusOK {
		fmt.Println("Error:", http.StatusText(code))
		return
	}

	// var r *http.Response
	// _ = r

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	body := string(bodyBytes)
	fmt.Println(body)
}
