package client

import (
	"fmt"
	"io"
	"net/http"
)

func GetHttpClient() *http.Client {
	return &http.Client{}
}

func RequestDo(method, url string, body io.Reader) (*http.Response, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(http.CanonicalHeaderKey("Authorization"), GetToken())

	resp, err := GetHttpClient().Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func CloseResponseBody(resp *http.Response) {
	if err := resp.Body.Close(); err != nil {
		fmt.Printf("failed to close response body: %s\n", err)
	}
}
