package main

import (
	"os"
	"strconv"

	c "github.com/wcharczuk/go-chart"
)

// You need to run:
// go get -u github.com/wcharczuk/go-chart

type chartReport struct {
	title         string
	width, height int
}

func (s *chartReport) report(results iterator) error {
	w := os.Stdout

	donut := c.DonutChart{
		Title: s.title,
		TitleStyle: c.Style{
			FontSize:  35,
			Show:      true,
			FontColor: c.ColorAlternateGreen,
		},
		Width:  s.width,
		Height: s.height,
	}

	results.each(func(r result) {
		v := c.Value{
			Label: r.domain + r.page + ": " + strconv.Itoa(r.visits),
			Value: float64(r.visits),
			Style: c.Style{
				FontSize: 14,
			},
		}

		donut.Values = append(donut.Values, v)
	})

	return donut.Render(c.SVG, w)
}
