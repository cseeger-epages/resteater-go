# RestEater-go eats your REST API

[![GoDoc](https://img.shields.io/badge/godoc-reference-green.svg)](https://godoc.org/github.com/cseeger-epages/resteater-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/cseeger-epages/resteater-go)](https://goreportcard.com/report/github.com/cseeger-epages/resteater-go)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/cseeger-epages/resteater-go/blob/master/LICENSE)
[![Build Status](https://travis-ci.org/cseeger-epages/resteater-go.svg?branch=master)](https://travis-ci.org/cseeger-epages/resteater-go)

This is just a simple addition to my [restfool-go](https://github.com/cseeger-epages/restfool-go) API implementation to consume REST Endpoints.

## Features
- BasicAuth support
- TLS settings
- supports additional requests Headers


## example

```
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
```
