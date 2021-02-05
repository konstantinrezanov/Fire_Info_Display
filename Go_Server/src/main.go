package main

import (
	"encoding/json"
	"log"
	"net/http"
	"Go_API/src"
)

func main() {
	fs:=http.FileServer(http.Dir("./public"))

	http.Handle("/", fs)
	
	http.HandleFunc("/data",dataServer)
	log.Fatal(http.ListenAndServe("192.168.1.240:8081", nil))
}

func dataServer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")

	raw := json.RawMessage(string(GO_API.Ret()))
	var objmap map[string]*json.RawMessage
	err:= json.Unmarshal(raw, &objmap)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(objmap)
}
