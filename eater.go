package resteater

import "net/url"

// Eater struct for implementing the EaterMethods interface
type Eater struct {
	Address, Username, Password string
	Port                        int
	BasicAuth                   bool
	VerifyTLS                   bool
	Debug                       bool
}

// Request contains request information
type Request struct {
	Form   url.Values
	Path   string
	Method string
	Eater  Eater
	Header []header
}

type header struct {
	Key, Value string
}
