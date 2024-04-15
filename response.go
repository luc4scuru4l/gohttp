package http

type Response struct {
	statusCode   uint16
	responseBody []byte
}

func (h *Response) getResponseBody() []byte {
	return h.responseBody
}

func (h *Response) getStatusCode() uint16 {
	return h.statusCode
}

func NewResponse(responseBody []byte, statusCode uint16) Response {
	h := Response{responseBody: responseBody, statusCode: statusCode}
	return h
}
