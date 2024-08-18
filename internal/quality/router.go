package quality

import "net/http"

func RegisterRoutes(router *http.ServeMux, handler *QualityController) {
	router.HandleFunc("GET /quality/{id}", handler.FindQualityByID)
	router.HandleFunc("GET /quality", handler.ShowQualityPerProject)
	router.HandleFunc("POST /quality/add", handler.CreateQuality)
	router.HandleFunc("PUT /quality/{id}", handler.UpdateQuality)
}
