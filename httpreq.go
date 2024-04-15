package http

import handler "net/http"

/* type httpRequestHandler interface {
	AddHeader(string, string)
	LoadBody([]byte)
	LoadBodyFromString(string)
	SetCharset(string)
	SetVerb(string)
} */

/*
 */
type httpRequest struct {
	params  map[string][]string
	verb    string
	body    *[]byte
	url     string
	handler handler.Request
}

/*Añade un item al header de la petición */
func (this *httpRequest) AddHeader(param string, value string) {
	if this.params == nil {
		this.params = make(map[string][]string)
	}
	this.params[param] = append(this.params[param], value)
}

func (this *httpRequest) LoadBody(body *[]byte) {
	this.body = body
}

func (this *httpRequest) SetVerb(verb string) {
	this.verb = verb
}

func (this *httpRequest) GetVerb() string {
	verb := this.verb
	if &verb == nil || len(verb) == 0 {
		verb = "POST"
	}
	return verb
}

func (this *httpRequest) SetUrl(url string) {
	this.url = url
}

func (this *httpRequest) GetUrl() string {
	return this.url
}
