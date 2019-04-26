package fetch

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func (t *HTTPTransfer) fakeStart() {
	jitter()
	t.fakeRequest()
	t.signal()

	for t.Downloaded < 100 {
		jitter()
		t.fakeFetch()

		if !t.signal() {
			// done signal received or there was an error
			break
		}
	}

	t.fakeFinish()
	t.signal()
}

// request requests to the url and adjusts the total length.
func (t *HTTPTransfer) fakeRequest() {
	t.Total = 100

	if debug {
		fmt.Printf("[TRANSFER] started: %s\n", t.URL)
	}
}

// fetch fetches a bit from the resource
func (t *HTTPTransfer) fakeFetch() {
	// TODO: right now hanged goroutine may hang the download completion
	// needs timeout
	// if t.URL == "url1" {
	// 	select {}
	// }

	// NOTE: burada sayacli io.Writer kullan
	if t.URL == "url1" && t.Downloaded > rand.Intn(50) {
		t.Error = errors.New("cekemedim netten")
	}

	n := rand.Intn(20) + 1
	if nn := t.Downloaded + n; nn > 100 {
		n = 100 - t.Downloaded
	}
	t.Current = n
	t.Downloaded += n
}

// finish signals the finish signal to the listeners
func (t *HTTPTransfer) fakeFinish() {
	t.Current = 0
	t.Done = true

	if debug {
		fmt.Printf("[TRANSFER] DONE: %s\n", t.URL)
	}
}

func jitter() {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)+1))
}
