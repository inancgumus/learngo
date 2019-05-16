package report_test

import (
	"testing"

	"github.com/inancgumus/learngo/28-methods/logparser-testing/report"
)

func TestSummaryTotal(t *testing.T) {
	p := newParser("a.com 1 2")
	p.Parse("b.com 3 4")

	s := p.Summarize()

	want := report.Result{Domain: "", Visits: 4, TimeSpent: 6}
	if got := s.Total(); want != got {
		t.Errorf("want: %+v; got: %+v", want, got)
	}
}

func TestSummaryIterator(t *testing.T) {
	p := newParser("a.com 1 2")
	p.Parse("a.com 3 4")
	p.Parse("b.com 5 6")

	s := p.Summarize()
	next, cur := s.Iterator()

	wants := []report.Result{
		{Domain: "a.com", Visits: 4, TimeSpent: 6},
		{Domain: "b.com", Visits: 5, TimeSpent: 6},
	}

	for _, want := range wants {
		t.Run(want.Domain, func(t *testing.T) {
			if got := next(); !got {
				t.Errorf("next(): want: %t; got: %t", true, got)
			}
			if got := cur(); want != got {
				t.Errorf("cur(): want: %+v; got: %+v", want, got)
			}
		})
	}
}
