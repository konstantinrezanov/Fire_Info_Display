package ratePlot

import (
	"bytes"
	"fmt"
	"time"
	"github.com/wcharczuk/go-chart"
)

func Plot(data RateData) string {
	graph := chart.Chart{
		Series: []chart.Series{
			chart.TimeSeries{
				Style: chart.Style{
					StrokeColor: chart.GetDefaultColor(0).WithAlpha(64),
					FillColor:   chart.GetDefaultColor(0).WithAlpha(64),
				},
				XValues: []time.Time{
					time.Now().AddDate(0, 0, -6),
					time.Now().AddDate(0, 0, -5),
					time.Now().AddDate(0, 0, -4),
					time.Now().AddDate(0, 0, -3),
					time.Now().AddDate(0, 0, -2),
					time.Now().AddDate(0, 0, -1),
					time.Now(),
				},
				YValues: reverse(data.Values),
			},
		},
	}
	buffer := bytes.NewBuffer([]byte{})
	_ = graph.Render(chart.SVG, buffer)
	return fmt.Sprintln(buffer)
}

func reverse(numbers []float64) []float64 {
	newNumbers := make([]float64, len(numbers))
	for i, j := 0, len(numbers)-1; i <= j; i, j = i+1, j-1 {
		newNumbers[i], newNumbers[j] = numbers[j], numbers[i]
	}
	return newNumbers
}