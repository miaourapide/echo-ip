package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type IpResponse struct {
	Ip string `json:"ip"`
}

func IpHandler(w http.ResponseWriter, r *http.Request) {
	ipRes := IpResponse{r.RemoteAddr}
	res, err := json.Marshal(ipRes)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(res)
}

func main() {
	http.HandleFunc("/ip", IpHandler)
	log.Println("Listening on port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}
