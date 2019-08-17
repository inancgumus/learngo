package main

import (
	"io"
	"sort"
	"strconv"

	c "github.com/wcharczuk/go-chart"
)

func chartWriter(w io.Writer) outputFn {
	return func(res []result) error {
		return chartWrite(w, res)
	}
}

func chartWrite(w io.Writer, res []result) error {
	sort.Slice(res, func(i, j int) bool {
		return res[i].domain > res[j].domain
	})

	donut := c.DonutChart{
		Title: "Total Visits Per Domain",
		TitleStyle: c.Style{
			FontSize:  35,
			Show:      true,
			FontColor: c.ColorAlternateGreen,
		},
		Width:  1920,
		Height: 800,
	}

	for _, r := range res {
		v := c.Value{
			Label: r.domain + r.page + ": " + strconv.Itoa(r.visits),
			Value: float64(r.visits),
			Style: c.Style{
				FontSize: 14,
			},
		}

		donut.Values = append(donut.Values, v)
	}

	return donut.Render(c.SVG, w)
}
