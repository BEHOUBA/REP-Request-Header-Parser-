package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// RequestHeader struct for the json data to be send as response
type RequestHeader struct {
	Ipaddress string `json:"ipaddress"`
	Language  string `json:"language"`
	Software  string `json:"software"`
	HostName  string `json:"hostname"`
}

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/", home)
	router.HandleFunc("/api/whoami/", reqHeaderParser)

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	server.ListenAndServe()
}

func reqHeaderParser(w http.ResponseWriter, r *http.Request) {
	ip := r.Host
	lang := r.Header["Accept-Language"][0][:5]
	softW := r.UserAgent()
	host, err := os.Hostname()
	if err != nil {
		host = "unknown"
	}
	response := RequestHeader{Ipaddress: ip, Language: lang, Software: softW, HostName: host}

	jsonResp, _ := json.MarshalIndent(response, "", "\t")

	fmt.Fprint(w, string(jsonResp))
}

func home(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path != "/" {
		fmt.Fprint(w, "Cannot get url: "+path+"\n\nTo use this api just add /api/whoami/ to the root url")
		return
	}
	templ := template.Must(template.ParseFiles("index.html"))

	templ.Execute(w, nil)
}
