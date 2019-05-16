package fetch

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

// HTTPTransfer uses a fetcher and a storager to fetch and store an url
type HTTPTransfer struct {
	Progress
	r io.ReadCloser
	w io.WriteCloser

	timeout time.Duration

	done    <-chan bool
	updates chan<- Progress
}

// NewHTTPTransfer creates and returns a new transfer
func NewHTTPTransfer(url string, timeout time.Duration) *HTTPTransfer {
	return &HTTPTransfer{
		Progress: Progress{URL: url},
		timeout:  timeout,
	}
}

// Start starts the transfer progress
// The transfer will send its updates to the update chan
// and it will stop when receives a signal from the done chan
//
// All the t.signal calls are here. I believe this increases the visibility
// of the Start function (what it does). If the signal calls were in the
// other funcs of HTTPTransfer, it could easily lead to bugs.
func (t *HTTPTransfer) Start(updates chan<- Progress, done <-chan bool) {
	defer t.cleanup()

	t.done, t.updates = done, updates

	t.w, t.Error = os.Create(path.Base(t.URL))
	if t.Error != nil {
		t.signal()
		return
	}

	t.request()
	if !t.signal() {
		return
	}

	// sniff the on-going transfer until the signal returns false
	sniff := sniffer(func(p []byte) bool {
		l := len(p)
		t.Current = l
		t.Downloaded += l

		return t.signal()
	})

	t.transfer(sniff)
	t.signal()
}

func (t *HTTPTransfer) cleanup() {
	if t.r != nil {
		t.r.Close()
	}
	if t.w != nil {
		t.w.Close()
	}
}

func (t *HTTPTransfer) request() {
	var resp *http.Response

	resp, t.Error = HTTPGet(t.URL, t.timeout)
	if t.Error != nil {
		return
	}

	t.Total = int(resp.ContentLength) // TODO: int(int64)
	if t.Total <= 0 {
		t.Error = fmt.Errorf("unknown content length: %d", t.Total)
	}

	t.r = resp.Body
}

func (t *HTTPTransfer) transfer(sniff sniffer) {
	// initiate the transfer and monitor it
	_, t.Error = io.Copy(io.MultiWriter(t.w, sniff), t.r)

	// if the err is from sniffer ignore it
	if _, ok := t.Error.(sniffer); ok {
		t.Error = nil
	}

	// the next signal will say: "no new bytes received" and its done
	t.Current = 0
	t.Done = true
}

// signal signals the listeners about the last transfer progress.
// it returns false when the done signal is received or there was an error
// in the transfer.
func (t *HTTPTransfer) signal() bool {
	select {
	case t.updates <- t.Progress:
	case <-t.done:
		// shutting down signal received
		return false
	}

	// check the error only after sending the last progress
	// if this check was above, the last update won't be sent
	if t.Error != nil {
		return false
	}

	// go on your duties
	return true
}
