package presence

import (
	"encoding/json"
	models "github.com/BerkatPS/internal"
	"github.com/BerkatPS/pkg/utils"
	"net/http"
)

type PresenceController struct {
	presenceService PresenceService
}

func NewPresenceController(presenceService PresenceService) *PresenceController {
	return &PresenceController{presenceService}
}

func (p *PresenceController) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	presences, err := p.presenceService.FindAll(ctx)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve presences: " + err.Error(),
		})
		return
	}
	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Presences found successfully",
		"data":    presences,
	})

}

func (p *PresenceController) FindPresenceByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	id, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid presence ID: " + err.Error(),
		})
		return
	}

	presence, err := p.presenceService.FindPresenceByID(ctx, id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve presence: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Presence found successfully",
		"data":    presence,
	})
}

func (p *PresenceController) FindPresenceByUserID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	userID, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid user ID: " + err.Error(),
		})
		return
	}

	presence, err := p.presenceService.FindPresenceByUserID(ctx, userID)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve presence: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Presence found successfully",
		"data":    presence,
	})
}

func (p *PresenceController) CreatePresence(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var presence models.Presence
	if err := json.NewDecoder(r.Body).Decode(&presence); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid presence data: " + err.Error(),
		})
		return
	}

	if err := p.presenceService.CreatePresence(ctx, &presence); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to create presence: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Presence created successfully",
		"data":    presence,
	})
}

func (p *PresenceController) UpdatePresence(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var presence models.Presence
	if err := json.NewDecoder(r.Body).Decode(&presence); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid presence data: " + err.Error(),
		})
		return
	}

	if err := p.presenceService.UpdatePresence(ctx, &presence); err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to update presence: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Presence updated successfully",
		"data":    presence,
	})
}
