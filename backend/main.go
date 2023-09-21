package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func meHandler(w http.ResponseWriter, req *http.Request) {
	url := "https://api.intra.42.fr/oauth/token?grant_type=client_credentials"
	requestUrl := fmt.Sprintf(
		"%s&%s&%s",
		url,
		"client_id=MY_AWESOME_UID",
		"client_secret=MY_AWESOME_SECRET",
	)
	resp, err := http.Get(requestUrl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte (err.Error()))
		return 
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte (err.Error()))
		return 
	}
	w.Write(b)
}

func main() {
	http.HandleFunc("/me", meHandler)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
