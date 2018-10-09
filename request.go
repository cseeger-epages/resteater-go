package resteater

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// CreateRequest creates a request object
func (a *Eater) CreateRequest(path string, method string, form url.Values) Request {
	return Request{
		Form:   form,
		Path:   path,
		Method: method,
		Eater:  *a,
	}
}

// Go requests your REST endpoint
func (r *Request) Go() (*http.Response, error) {
	uri := fmt.Sprintf("https://%s:%d%s", r.Eater.Address, r.Eater.Port, r.Path)

	if r.Eater.Debug {
		debug("Request URL: ", uri)
	}

	req, err := http.NewRequest(r.Method, uri, strings.NewReader(r.Form.Encode()))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(r.Eater.Username, r.Eater.Password)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	for _, h := range r.Header {
		req.Header.Add(h.Key, h.Value)
	}

	tr := &http.Transport{}

	if !r.Eater.VerifyTLS {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
	}

	client := &http.Client{Transport: tr}

	return client.Do(req)
}

// AddHeader adds an additional header to your request
func (r *Request) AddHeader(key string, value string) {
	r.Header = append(r.Header, header{key, value})
}
