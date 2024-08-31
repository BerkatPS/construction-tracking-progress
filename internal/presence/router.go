package presence

import "net/http"

func RegisterRoutes(router *http.ServeMux, handler *PresenceController) {
	router.HandleFunc("GET /presences", handler.FindAll)
	router.HandleFunc("GET /presences/:id", handler.FindPresenceByID)
	router.HandleFunc("GET /presences/user/:id", handler.FindPresenceByUserID)
	router.HandleFunc("POST /presences", handler.CreatePresence)
	router.HandleFunc("PUT /presences/:id", handler.UpdatePresence)
}
