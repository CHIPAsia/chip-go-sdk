package http

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func request(method string, host string, headers Header, data *string) ([]byte, int, error) {
	client := &http.Client{}

	var requestData io.Reader = nil
	if data != nil {
		requestData = strings.NewReader(*data)
	}

	request, _ := http.NewRequest(method, host, requestData)

	for key, value := range headers {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("Request:Do:Method:%s \nURL:%s \nError:%+v\n", method, host, err)
		return nil, response.StatusCode, err
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Request:ReadAll:Method:%s \nURL:%s \nError:%+v\n", method, host, err)
		return nil, response.StatusCode, err
	}

	if response.StatusCode > 399 {
		fmt.Printf("Request:Response:StatusCode:%d \nMethod:%s \nURL:%s \nError:%+v\n", response.StatusCode, method, host, err)
		err = fmt.Errorf("error with status code %d", response.StatusCode)
		return body, response.StatusCode, err
	}
	return body, response.StatusCode, nil
}

func addJsonHeaders(headers Header) Header {
	if headers == nil {
		headers = Header{
			"Content-Type": "application/json",
			"Accept":       "application/json",
		}
		return headers
	}

	headers["Content-Type"] = "application/json"
	headers["Accept"] = "application/json"
	return headers
}

func (h DefaultHttpRequest) post(host string, headers Header, data *string) ([]byte, int, error) {
	return request("POST", host, headers, data)
}

func (h DefaultHttpRequest) PostJson(host string, headers Header, data *string) ([]byte, int, error) {
	headers = addJsonHeaders(headers)
	return h.post(host, headers, data)
}

func (h DefaultHttpRequest) get(host string, headers Header) ([]byte, int, error) {
	return request("GET", host, headers, nil)
}

func (h DefaultHttpRequest) GetJson(host string, headers Header) ([]byte, int, error) {
	headers = addJsonHeaders(headers)
	return h.get(host, headers)
}

func (h DefaultHttpRequest) patch(host string, headers Header, data *string) ([]byte, int, error) {
	return request("PATCH", host, headers, data)
}

func (h DefaultHttpRequest) PatchJson(host string, headers Header, data *string) ([]byte, int, error) {
	headers = addJsonHeaders(headers)
	return h.patch(host, headers, data)
}

func (h DefaultHttpRequest) put(host string, headers Header, data *string) ([]byte, int, error) {
	return request("PUT", host, headers, data)
}

func (h DefaultHttpRequest) PutJson(host string, headers Header, data *string) ([]byte, int, error) {
	headers = addJsonHeaders(headers)
	return h.put(host, headers, data)
}

func (h DefaultHttpRequest) delete(host string, headers Header) ([]byte, int, error) {
	return request("DELETE", host, headers, nil)
}

func (h DefaultHttpRequest) DeleteJson(host string, headers Header) ([]byte, int, error) {
	headers = addJsonHeaders(headers)
	return h.delete(host, headers)
}
