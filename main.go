package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe("localhost:8001", router))
}

func getTime(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string, 0)
	tz := r.URL.Query().Get("tz")
	timezones := strings.Split(tz, ",")

	if len(timezones) == 0 {
		loc, err := time.LoadLocation(tz)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("invalid timezone %s", tz)))
		} else {
			response["current_time"] = time.Now().In(loc).String()
			w.Header().Add("Content-Type", "application/json")
			json.NewEncoder(w).Encode(response)
		}
	} else {
		for _, tzdb := range timezones {
			loc, err := time.LoadLocation(tzdb)
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(fmt.Sprintf("invalid timezone %s in input", tzdb)))
				return
			}
			now := time.Now().In(loc)
			response[tzdb] = now.String()
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
