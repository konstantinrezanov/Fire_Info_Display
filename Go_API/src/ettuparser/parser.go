package ettuparser

import (
	"github.com/antchfx/htmlquery"
)

const ettuURL = "https://mobile.ettu.ru/station/"

type Status struct {
	Routes    []string
	Distances []string
	Arrivals  []string
}

//ParseStation -parses ETTU station by id
func ParseStation(id string) Status {
	doc, _ := htmlquery.LoadURL(ettuURL + id)
	var data Status
	routes := htmlquery.Find(doc,
		"//div[@style=\"width: 3em;display:inline-block;text-align:center;\"]/b")
	distances := htmlquery.Find(doc,
		"//div[@style=\"width: 4em;display:inline-block;text-align:right;\"]/text()")
	arrivesAt := htmlquery.Find(doc,
		"//div[@style=\"width: 5em;display:inline-block;text-align:right;\"]/text()")
	for i := 0; i < len(routes); i++ {
		data.Routes = append(data.Routes, htmlquery.InnerText(routes[i]))
		data.Distances = append(data.Distances, htmlquery.InnerText(distances[i]))
		data.Arrivals = append(data.Arrivals, htmlquery.InnerText(arrivesAt[i]))
	}

	return data
}
