package connpass

import (
	"io"
	"log"
	"net/http"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Fatal("Method not allowed")
	}

	resp, err := http.Get("https://connpass.com/api/v1/event?" + r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusServiceUnavailable)
		log.Fatal("Failed fetching connpass events: ", err)
		defer resp.Body.Close()
	}
	defer resp.Body.Close()

	byteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatal("Failed Reading response body: ", err)
	}

	w.Write(byteArray)
}
