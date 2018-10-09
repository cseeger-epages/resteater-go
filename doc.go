/*
Package resteater is a simple REST API consumtion implementation

import resteater-go

	import(
		eater "github.com/cseeger-epages/resteater-go"
	)

initialize your eater

	e := eater.NewEater("127.0.0.1", 9443)

create a request

	req := e.CreateRequest("/help", "POST", url.Values{})

request your resource and get a *http.Responsea

	resp, err := req.Go()
	if err != nil {
		log.Fatal(err)
	}
*/
package resteater
