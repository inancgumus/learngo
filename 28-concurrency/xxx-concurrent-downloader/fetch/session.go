package fetch

import (
	"fmt"
	"sync"
)

const debug = true

// Session manages the downloading process
type Session struct {
	done chan bool
}

// Transfer sends `updates` and terminates when `done` closes
type Transfer interface {
	Start(updates chan<- Progress, done <-chan bool)
}

// NewSession creates a new downloading session
func NewSession() *Session {
	return &Session{done: make(chan bool)}
}

// Start starts the downloading process
func (s *Session) Start(transfers ...Transfer) <-chan Progress {
	// a buffered chan may unblock transfers in case of a slow ui
	updates := make(chan Progress)

	var wg sync.WaitGroup
	wg.Add(len(transfers))

	for _, t := range transfers {
		go func(t Transfer) {
			t.Start(updates, s.done)
			wg.Done()
		}(t)
	}

	go func() {
		wg.Wait()      // wait until all downloads complete
		close(updates) // let the watchers (ui) know that we're shutting down
	}()

	return updates
}

// Shutdown stops the downloading process and sends a signal to all parties
func (s *Session) Shutdown() {
	// let the transfers know we're shutting down
	// when this is done s.updates will be closed in the Start() routine above
	close(s.done)

	if debug {
		fmt.Printf("[SESSION ] DONE\n")
	}
}
