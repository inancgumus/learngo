// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

package pipe

/*
// You need to run:
// go get -u github.com/wcharczuk/go-chart

// Chart renders a chart.
type Chart struct {
	Title         string
	Width, Height int

	w io.Writer
}

// NewChartReport returns a Chart report generator.
func NewChartReport(w io.Writer) *Chart {
	return &Chart{w: w}
}

// Consume generates a chart report.
func (c *Chart) Consume(records Iterator) error {
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

	err := records.Each(func(r Record) error {
		v := chart.Value{
			Label: r.Domain + r.Page + ": " + strconv.Itoa(r.Visits),
			Value: float64(r.Visits),
			Style: chart.Style{
				FontSize: 14,
			},
		}

		donut.Values = append(donut.Values, v)

		return nil
	})
	if err != nil {
		return err
	}

	return donut.Render(chart.SVG, w)
}
*/
