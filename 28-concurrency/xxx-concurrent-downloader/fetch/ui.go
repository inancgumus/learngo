package fetch

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

const refreshPeriod = time.Second / 10

// uiProgress is the default UI for the downloader
type uiProgress struct {
	urls      []string
	transfers map[string]Progress
}

// UI listens for the progress updates from the updates chan
// and it refreshes the ui
func UI(updates <-chan Progress) {
	ui := &uiProgress{transfers: make(map[string]Progress)}

	// NOTE: we didn't use time.After here directly
	// because doing so can create a lot of Timer chans unnecessarily
	// instead we're just waiting on the same timer value
	tick := time.After(refreshPeriod)
	for {
		select {
		case p, ok := <-updates:
			// if no more updates close the ui
			if !ok {
				ui.shutdown()
				return
			}
			ui.update(p)

		case <-tick:
			// `case <-tick:` allows updating the ui independently
			// from the progress update signals. or the ui would hang
			// the updaters (transfers).
			ui.refresh()
			tick = time.After(refreshPeriod)
		}
	}

}

// shutdown refreshes the ui for the last time and closes it
func (ui *uiProgress) shutdown() {
	ui.refresh()

	if debug {
		fmt.Printf("[ UI ] DONE\n")
	}
}

// update updates the progress data from the received message
func (ui *uiProgress) update(p Progress) {
	if _, ok := ui.transfers[p.URL]; !ok {
		ui.urls = append(ui.urls, p.URL)
	}

	// update the latest progress for the url
	ui.transfers[p.URL] = p
}

// refresh refreshes the UI with the latest progress
func (ui *uiProgress) refresh() {
	if !debug {
		screen.Clear()
		screen.MoveTopLeft()
	}

	var total, downloaded int

	for _, u := range ui.urls {
		p := ui.transfers[u]

		msg := "Downloading"
		if p.Done && p.Error == nil {
			msg = "👍 Completed"
		}
		if p.Error != nil {
			msg = fmt.Sprintf("❌ %s", p.Error)
		}

		fmt.Println(p.URL)
		fmt.Printf("\t%d/%d\n", p.Downloaded, p.Total)
		fmt.Printf("\t%s\n", msg)

		total += p.Total
		downloaded += p.Downloaded
	}

	fmt.Printf("\n%s %d/%d\n", "TOTAL DOWNLOADED BYTES:", downloaded, total)
}
