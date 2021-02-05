package main

import (
	"ettuparser"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost"+":8081", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, ettuparser.ParseStation(r.FormValue("station")))
}
