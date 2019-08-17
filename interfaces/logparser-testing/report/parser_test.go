package report_test

import (
	"strings"
	"testing"

	"github.com/inancgumus/learngo/28-methods/logparser-testing/report"
)

func newParser(lines string) *report.Parser {
	p := report.New()
	p.Parse(lines)
	return p
}

func TestParserLineErrs(t *testing.T) {
	p := newParser("a.com 1 2")
	p.Parse("b.com -1 -1")

	want := "#2"
	err := p.Err().Error()

	if !strings.Contains(err, want) {
		t.Errorf("want: %q; got: %q", want, err)
	}
}

func TestParserStopsOnErr(t *testing.T) {
	p := newParser("a.com 10 20")
	p.Parse("b.com -1 -1")
	p.Parse("neverparses.com 30 40")

	s := p.Summarize()
	if want, got := 10, s.Total().Visits; want != got {
		t.Errorf("want: %d; got: %d", want, got)
	}
}

func TestParserIncorrectFields(t *testing.T) {
	tests := []struct {
		in, name string
	}{
		{"a.com", "missing fields"},
		{"a.com -1 2", "incorrect visits"},
		{"a.com 1 -1", "incorrect time spent"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if p := newParser(tt.in); p.Err() == nil {
				t.Errorf("in: %q; got: nil err", tt.in)
			}
		})
	}
}
