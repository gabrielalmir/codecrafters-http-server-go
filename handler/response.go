package handler

import "strconv"

type Response struct {
	StatusCode int
	Headers    map[string]string
	Body       string
}

func (r *Response) AddHeader(key string, value string) {
	r.Headers[key] = value
}

func (r *Response) SetBody(body string) {
	r.Body = body
}

func (r *Response) SetStatusCode(statusCode int) {
	r.StatusCode = statusCode
}

func (r *Response) String() string {
	statusCode := r.StatusCode
	headers := r.Headers
	body := r.Body
	contentLength := len(body)

	response := "HTTP/1.1 " + strconv.Itoa(statusCode) + "\r\n"

	for key, value := range headers {
		response += key + ": " + value + "\r\n"
	}

	response += "Content-Length: " + strconv.Itoa(contentLength) + "\r\n"
	response += "\r\n"
	response += body

	return response
}

func NewResponse(
	statusCode int,
	headers map[string]string,
	body string,
) *Response {
	return &Response{
		StatusCode: statusCode,
		Headers:    headers,
		Body:       body,
	}
}

func NotFound(r []byte) string {
	return SendResponse(r, 404, map[string]string{}, "Not Found")
}

func SendResponse(r []byte, statusCode int, headers map[string]string, body string) string {
	response := NewResponse(statusCode, headers, body)
	return response.String()
}
