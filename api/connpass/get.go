package connpass

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetEvents(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://connpass.com/api/v1/event")
	if err != nil {
		log.Fatal("fetching connpass events", err)
	}

	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)

	fmt.Fprint(w, string(byteArray))
}
