package report

/*
// You need to run:
// go get -u github.com/wcharczuk/go-chart

// Chart renders a chart.
type Chart struct {
	Title         string
	Width, Height int

	w io.Writer
}

// AsChart returns a Chart report generator.
func AsChart(w io.Writer) *Chart {
	return &Chart{w: w}
}

// Digest generates a chart report.
func (c *Chart) Digest(records pipe.Iterator) error {
	w := os.Stdout

	donut := chart.DonutChart{
		Title: c.Title,
		TitleStyle: chart.Style{
			FontSize:  35,
			Show:      true,
			FontColor: chart.ColorAlternateGreen,
		},
		Width:  c.Width,
		Height: c.Height,
	}

	records.Each(func(r pipe.Record) {
		v := chart.Value{
			Label: r.Str("domain") + r.Str("page") + ": " + strconv.Itoa(r.Int("visits")),
			Value: float64(r.Int("visits")),
			Style: chart.Style{
				FontSize: 14,
			},
		}

		donut.Values = append(donut.Values, v)
	})

	return donut.Render(chart.SVG, w)
}
*/
