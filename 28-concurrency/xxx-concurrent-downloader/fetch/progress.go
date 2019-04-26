package fetch

// Progress contains data about the downloading progress
type Progress struct {
	URL string

	Total, Downloaded, Current int
	Done                       bool
	Error                      error
}
