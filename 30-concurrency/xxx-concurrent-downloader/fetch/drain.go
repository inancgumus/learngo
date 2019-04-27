package fetch

// Drain drains the progress updates and returns the latest progresses
func Drain(updates <-chan Progress) map[string]Progress {
	latest := make(map[string]Progress)

	// save the latest progress
	for p := range updates {
		latest[p.URL] = p
	}
	return latest
}
