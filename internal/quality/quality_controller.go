package quality

import (
	"encoding/json"
	models "github.com/BerkatPS/internal"
	"github.com/BerkatPS/pkg/utils"
	"net/http"
)

type QualityController struct {
	QualityService QualityService
}

func NewQualityController(qualityService QualityService) *QualityController {
	return &QualityController{qualityService}
}

func (q *QualityController) FindQualityByTaskID(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	taskID, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid task ID: " + err.Error(),
		})
		return
	}

	quality, err := q.QualityService.FindQualityByTaskID(ctx, taskID)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve quality: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Quality found successfully",
		"data":    quality,
	})
}

// func (q *QualityController) FindQualityByDateRange(w http.ResponseWriter, r *http.Request) {

// 	ctx := r.Context()

// 	var dateRange models.QualityCheck
// 	if err := json.NewDecoder(r.Body).Decode(&dateRange); err != nil {
// 		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
// 			"status":  "error",
// 			"message": "Invalid request body: " + err.Error(),
// 		})
// 		return
// 	}
// 	qualities, err := q.QualityService.FindQualityByDateRange(ctx, dateRange)
// 	if err != nil {
// 		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
// 			"status":  "error",
// 			"message": "Failed to retrieve quality: " + err.Error(),
// 		})
// 		return
// 	}

// 	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
// 		"status":  "success",
// 		"message": "Quality found successfully",
// 		"data":    qualities,
// 	})
// }

func (q *QualityController) FindQualityIssues(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	qualities, err := q.QualityService.FindQualityIssues(ctx)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve quality: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Quality found successfully",
		"data":    qualities,
	})
}

func (q *QualityController) FindNonCompliantQualityChecks(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	qualities, err := q.QualityService.FindNonCompliantQualityChecks(ctx)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve quality: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Quality found successfully",
		"data":    qualities,
	})
}

func (q *QualityController) FindQualityChecksByInspector(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	inspectorID, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid inspector ID: " + err.Error(),
		})
		return
	}

	qualities, err := q.QualityService.FindQualityChecksByInspector(ctx, inspectorID)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve quality: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Quality found successfully",
		"data":    qualities,
	})
}

// func (q *QualityController) UpdateQualityStatus( w http.ResponseWriter, r *http.Request) {

// 	ctx := r.Context()

// 	var quality models.QualityCheck
// 	if err := json.NewDecoder(r.Body).Decode(&quality); err != nil {
// 		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
// 			"status":  "error",
// 			"message": "Invalid request body: " + err.Error(),
// 		})
// 		return
// 	}

// 	err := q.QualityService.UpdateQualityStatus(ctx, &quality)
// 	if err != nil {
// 		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
// 			"status":  "error",
// 			"message": "Failed to update quality: " + err.Error(),
// 		})
// 		return
// 	}

// 	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
// 		"status":  "success",
// 		"message": "Quality updated successfully",
// 	})
// }
// func (q *QualityController) FindQualityByProjectID(w http.ResponseWriter, r *http.Request) {

// 	ctx := r.Context()

// 	projectID, err := utils.ParseInt64Param(r)
// 	if err != nil {
// 		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
// 			"status":  "error",
// 			"message": "Invalid project ID: " + err.Error(),
// 		})
// 		return
// 	}

// 	qualities, err := q.QualityService.FindQualityByProjectID(ctx, projectID)
// 	if err != nil {
// 		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
// 			"status":  "error",
// 			"message": "Failed to retrieve quality: " + err.Error(),
// 		})
// 		return
// 	}

// 	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
// 		"status":  "success",
// 		"message": "Quality found successfully",
// 		"data":    qualities,
// 	})

// }


func (q *QualityController) FindQualityByID(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	qualityID, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid quality ID: " + err.Error(),
		})
		return
	}

	quality, err := q.QualityService.FindQualityByID(ctx, qualityID)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve quality: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Quality found successfully",
		"data":    quality,
	})
}

func (q *QualityController) ShowQualityPerProject(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	projectID, err := utils.ParseInt64Param(r)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid project ID: " + err.Error(),
		})
		return
	}

	qualities, err := q.QualityService.ShowQualityPerProject(ctx, projectID)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to retrieve quality: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Quality found successfully",
		"data":    qualities,
	})
}

func (q *QualityController) CreateQuality(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var quality models.QualityCheck
	if err := json.NewDecoder(r.Body).Decode(&quality); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request body: " + err.Error(),
		})
		return
	}

	err := q.QualityService.CreateQuality(ctx, &quality)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to create quality: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Quality created successfully",
		"data":    quality,
	})
}

func (q *QualityController) UpdateQuality(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var quality models.QualityCheck
	if err := json.NewDecoder(r.Body).Decode(&quality); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "Invalid request body: " + err.Error(),
		})
		return
	}

	err := q.QualityService.UpdateQuality(ctx, &quality)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusInternalServerError, map[string]interface{}{
			"status":  "error",
			"message": "Failed to update quality: " + err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "Quality updated successfully",
		"data":    quality,
	})
}
