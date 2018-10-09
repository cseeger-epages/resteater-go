package resteater

import (
	"net/http"
	"net/url"
)

// EaterMethods describes Eather methods
type EaterMethods interface {

	// SetBasicAuth sets basic authenication credentials used by REST-API
	SetBasicAuth(string, string)

	// SetVerifyTLS lets you set TLS verification settings
	SetVerifyTLS(bool)

	// SetDebug lets you set debug settings
	SetDebug(bool)

	// Request ... requests
	CreateRequest(string, string, url.Values) (Request, error)
}

// RequestMethods describes Request methods
type RequestMethods interface {

	// Go requests your REST endpoint
	Go() (*http.Response, error)

	// AddHeader adds an additional header to your request
	AddHeader(string, string)
}
