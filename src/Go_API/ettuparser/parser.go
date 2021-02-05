package ettuparser

import (
	"github.com/antchfx/htmlquery"
	"encoding/json"
)

const ettuURL = "https://mobile.ettu.ru/station/"

type Statuses struct {
	Title string
	Routes    []string
	Distances []string
	Arrivals  []string
}


type station struct {
	ID     string
	Status Statuses
}
type stations struct {
	List []station
}

//ParseStation -parses ETTU station by id
func ParseStation(id string) Statuses {
	doc, _ := htmlquery.LoadURL(ettuURL + id)
	var data Statuses
	data.Title=htmlquery.InnerText(htmlquery.Find(doc,"//p[1]")[0])
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

func StationSend() string {
	var data stations
	pstations := []string{"3442", "962313"}
	for _, n := range pstations {
		data.List= append(data.List, station{
			ID:     n,
			Status: ParseStation(n),
		})
	}
	jsonData, _ := json.MarshalIndent(data, "", "  ")
	return string(jsonData)
}