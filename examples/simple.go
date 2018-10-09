package main

import (
	"fmt"
	"log"
	"net/url"

	eater "github.com/cseeger-epages/resteater-go"
)

func main() {

	// initialize your new eater
	e := eater.NewEater("127.0.0.1", 9443)

	// create a request
	req := e.CreateRequest("/help", "POST", url.Values{})

	// fire your request and get a *http.Response
	resp, err := req.Go()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.StatusCode)
}
