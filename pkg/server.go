package pkg

import (
	"app/api/connpass"
	"app/ent"
	"app/ent/hook"
	"net/http"
)

func SetupHandler(mux *http.ServeMux) {
	mux.HandleFunc("/connpass/events", connpass.GetEvents)
	mux.HandleFunc("/connpass/events/thumbnail", connpass.GetEventThumbnail)
}

func SetupHooks(client *ent.Client) {
	client.User.Use(hook.UserHooks(client)...)
	client.Tag.Use(hook.TagHooks(client)...)
}
