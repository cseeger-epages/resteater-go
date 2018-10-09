package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	eater "github.com/cseeger-epages/resteater-go"
)

func main() {

	// initialize your new eater
	e := eater.NewEater("127.0.0.1", 9443)

	// set your basic auth credentials
	e.SetBasicAuth("username", "password")

	// disable TLS verification e.g. self-signed certificates
	e.SetVerifyTLS(false)

	// enable debug output
	e.SetDebug(true)

	// add post request parameter
	form := url.Values{}
	form.Add("key", "value")

	// create a request
	req := e.CreateRequest("/sth/set", "POST", form)

	// fire your request and get a *http.Response
	resp, err := req.Go()
	if err != nil {
		log.Fatal(err)
	}

	// check the StatusCode and parse your response
	if resp.StatusCode == http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	} else {
		log.Fatal(resp.StatusCode)
	}
}
