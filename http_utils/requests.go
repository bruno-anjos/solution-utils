package http_utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

func BuildRequest(method, host, path string, body interface{}) *http.Request {
	hostUrl := url.URL{
		Scheme: "http",
		Host:   host,
		Path:   path,
	}

	var (
		err        error
		request    *http.Request
		bodyBuffer *bytes.Buffer
	)

	if body != nil {
		var jsonStr []byte
		jsonStr, err = json.Marshal(body)
		if err != nil {
			panic(errors.WithStack(err))
		}
		bodyBuffer = bytes.NewBuffer(jsonStr)
	} else {
		bodyBuffer = new(bytes.Buffer)
	}

	request, err = http.NewRequest(method, hostUrl.String(), bodyBuffer)
	if err != nil {
		panic(errors.WithStack(err))
	}

	request.Header.Set("Content-Type", "application/json")

	return request
}

func DoRequest(httpClient *http.Client, request *http.Request, responseBody interface{}) (int, *http.Response) {
	log.Debugf("Doing request: %s %s", request.Method, request.URL.String())

	if httpClient == nil {
		panic(errorHttpClietNilFormat)
	}

	resp, err := httpClient.Do(request)
	if err != nil {
		panic(errors.WithStack(err))
	}

	if responseBody != nil {
		err = json.NewDecoder(resp.Body).Decode(responseBody)
		if err != nil {
			panic(errors.WithStack(err))
		}
	}

	return resp.StatusCode, resp
}
