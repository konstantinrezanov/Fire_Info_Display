package ettuparser

import (
	"fmt"

	"github.com/antchfx/htmlquery"
)

const ettuURL = "https://mobile.ettu.ru/station/"

//ParseStation -parses ETTU station by id
func ParseStation(id string) string {
	doc, _ := htmlquery.LoadURL(ettuURL + id)
	var ret string
	routes := htmlquery.Find(doc, 
		"//div[@style=\"width: 3em;display:inline-block;text-align:center;\"]/b")
	distances := htmlquery.Find(doc, 
		"//div[@style=\"width: 4em;display:inline-block;text-align:right;\"]/text()")
	arrivesAt := htmlquery.Find(doc, 
		"//div[@style=\"width: 5em;display:inline-block;text-align:right;\"]/text()")
	for i := 0; i < len(routes); i++ {
		ret+=fmt.Sprintf("%s\t%s\t%s\n", htmlquery.InnerText(routes[i]), 
		htmlquery.InnerText(distances[i]), htmlquery.InnerText(arrivesAt[i]))
	}
	return ret
}
