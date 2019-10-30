// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// MaxResponseSize limits the response bytes from the API
const MaxResponseSize = 2 << 16

// creating a robust http getter (lecture? :))
func request(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// Get the error from the context.
		// It may contain more useful data.
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
		}
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Bad Status: %d", resp.StatusCode)
	}

	// Prevents the api to shoot us unlimited amount of data
	r := io.LimitReader(resp.Body, MaxResponseSize)

	return ioutil.ReadAll(r)
}
