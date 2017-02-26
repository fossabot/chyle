package chyle

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// ErrCantReadHTTPResponse is triggered when reading http
// response body failed
type ErrCantReadHTTPResponse struct {
	URL *url.URL
}

// Error output error as string
func (e ErrCantReadHTTPResponse) Error() string {
	return fmt.Sprintf("can't read http response from %s", e.URL)
}

// setHeaders setup headers on request from a map header key -> header value
func setHeaders(request *http.Request, headers map[string]string) {
	for k, v := range headers {
		request.Header.Set(k, v)
	}
}

// sendRequest picks a request and send it with given client and handle all error
// boilerplate and return status code and body as byte slice
func sendRequest(client *http.Client, request *http.Request) (int, []byte, error) {
	rep, err := client.Do(request)

	if err != nil {
		return 0, nil, err
	}

	defer func() {
		err = rep.Body.Close()

		if err != nil {
			log.Fatal(err)
		}
	}()

	b, err := ioutil.ReadAll(rep.Body)

	if err != nil {
		return 0, nil, ErrCantReadHTTPResponse{request.URL}
	}

	return rep.StatusCode, b, nil
}