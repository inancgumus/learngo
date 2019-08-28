package main

// You need to run:
// go get -u github.com/wcharczuk/go-chart

// type chartReport struct {
// 	title         string
// 	width, height int
// }

// func (s *chartReport) digest(records iterator) error {
// 	w := os.Stdout

// 	donut := chart.DonutChart{
// 		Title: s.title,
// 		TitleStyle: chart.Style{
// 			FontSize:  35,
// 			Show:      true,
// 			FontColor: chart.ColorAlternateGreen,
// 		},
// 		Width:  s.width,
// 		Height: s.height,
// 	}

// 	records.each(func(r record) {
// 		v := chart.Value{
// 			Label: r.domain + r.page + ": " + strconv.Itoa(r.visits),
// 			Value: float64(r.visits),
// 			Style: chart.Style{
// 				FontSize: 14,
// 			},
// 		}

// 		donut.Values = append(donut.Values, v)
// 	})

// 	return donut.Render(chart.SVG, w)
// }
