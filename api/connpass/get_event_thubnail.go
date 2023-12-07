package connpass

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"regexp"
)

type EventThumbnail struct {
	Url string `json:"url"`
}

func GetEventThumbnail(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Fatal("Method not allowed")
	}

	resp, err := http.Get("https://connpass.com/event/" + r.URL.Query().Get("id"))
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

	buf := bytes.NewBuffer(byteArray)
	html := buf.String()

	imgMetaR := regexp.MustCompile("<meta.*property=\"og:image\".*content=\".*\"")
	imgMeta := imgMetaR.FindString(html)

	urlR := regexp.MustCompile("https://.*")
	url := urlR.FindString(imgMeta)

	urlLen := len(url)
	if urlLen < 1 {
		w.WriteHeader(http.StatusNotFound)
		log.Fatal("Failed to find thumbnail url")
	}

	url = url[:urlLen-1] // 末尾を削除

	eventThumbnail := EventThumbnail{
		Url: url,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(eventThumbnail); err != nil {
		log.Println(err)
	}
}
