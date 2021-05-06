package shared

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
)

func DoHttpRequest(method, url string, header, query map[string]string, data []byte) ([]byte, error) {
	if method != "GET" && method != "POST" {
		return nil, errors.New("get or post is required")
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	setQuery(req, query)
	setHeader(req, header)

	c := &http.Client{}
	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func setQuery(req *http.Request, query map[string]string) {
	// TODO: Query Validation
	q := req.URL.Query()
	for k, v := range query {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
}
func setHeader(req *http.Request, header map[string]string) {
	// TODO: Header Validation
	for k, v := range header {
		req.Header.Add(k, v)
	}
}
