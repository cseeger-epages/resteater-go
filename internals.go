package resteater

// NewEater creates a new Eater for your REST endpoint
func NewEater(addr string, port int) Eater {
	return Eater{
		Address:   addr,
		Port:      port,
		BasicAuth: false,
		VerifyTLS: true,
		Debug:     false,
		Username:  "",
		Password:  "",
	}
}

// SetBasicAuth activates basic auth and sets username and password
func (a *Eater) SetBasicAuth(user string, pass string) {
	a.BasicAuth = true
	a.Username = user
	a.Password = pass
}

// SetVerifyTLS sets InsecureSkipVerify to tls.Config
func (a *Eater) SetVerifyTLS(test bool) {
	a.VerifyTLS = test
}

// SetDebug sets debugging options
func (a *Eater) SetDebug(test bool) {
	a.Debug = test
}
