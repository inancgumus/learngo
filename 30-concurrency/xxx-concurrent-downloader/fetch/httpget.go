package fetch

import (
	"fmt"
	"net/http"
	"time"
)

// HTTPGet requests to a url with the specified timeout.
// And it returns the response with an error.
// The caller should drain and close the body or it will leak.
func HTTPGet(url string, timeout time.Duration) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Close = true
	if err != nil {
		return nil, err
	}

	c := &http.Client{Timeout: timeout}
	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	// checkout for the bad urls
	if s := resp.StatusCode; s < 200 || s > 299 {
		return resp, fmt.Errorf("bad status: %d", s)
	}

	return resp, nil
}
