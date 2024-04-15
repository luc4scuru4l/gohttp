package http

import (
	"bytes"
	"io"
	handler "net/http"
)

type http struct {
	params map[string][]string
	verb   string
	url    string
	port   uint16
	body   []byte
}

type HttpResponse interface {
	getResponseBody() []byte
	getStatusCode() uint16
}

func (h *http) addParam(paramName string, paramValue string) {
	if h.params == nil {
		h.params = make(map[string][]string)
	}
	h.params[paramName] = append(h.params[paramName], paramValue)
}

func (h *http) GetPort() uint16 {
	return h.port
}

func (h *http) GetVerb() string {
	return h.verb
}

func (h *http) GetUrl() string {
	return h.url
}

func (h *http) GetBody() []byte {
	return h.body
}

func (h *http) GetParams() map[string][]string {
	return h.params
}

func (h *http) GiveMeAResponse() (HttpResponse, error) {
	xmlReq := bytes.NewReader(h.GetBody())
	httpRequest, _ := handler.NewRequest(h.GetVerb(), h.GetUrl(), xmlReq)
	httpRequest.Header = h.GetParams()
	response, err := handler.DefaultClient.Do(httpRequest)
	if err != nil {
		return nil, err
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	HttpResponse := NewResponse(responseBody, uint16(response.StatusCode))

	return &HttpResponse, nil
}
