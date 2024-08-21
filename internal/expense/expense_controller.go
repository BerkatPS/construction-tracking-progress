package expense

import (
	"encoding/json"
	"net/http"
	"github.com/BerkatPS/pkg/utils"
	models "github.com/BerkatPS/internal"
	

)

type ExpenseController struct {
	ExpenseService ExpenseService
}

func NewExpenseController(expenseService ExpenseService) *ExpenseController {
	return &ExpenseController{ExpenseService: expenseService}
}

func (c *ExpenseController) CreateExpense(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var expense models.Expense

	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "invalid request body",
		})
		return
	}

	if err := c.ExpenseService.CreateExpense(ctx, expense); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusCreated, map[string]interface{}{
		"status":  "success",
		"message": "expense created successfully",
	})

}

func (c *ExpenseController) UpdateExpense(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var expense models.Expense

	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "invalid request body",
		})
		return
	}

	if err := c.ExpenseService.UpdateExpense(ctx, expense); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "expense updated successfully",
	})
}

func (c *ExpenseController) DeleteExpense(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var expense models.Expense

	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "invalid request body",
		})
		return
	}

	if err := c.ExpenseService.DeleteExpense(ctx, expense.ID); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "expense deleted successfully",
	})
}

func (c *ExpenseController) GetExpenseById(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	id , err := utils.ParseInt64Param(r)

	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "invalid request body",
		})
		return
	}

	expense, err := c.ExpenseService.GetExpenseById(ctx, id)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "expense retrieved successfully",
		"data":    expense,
	})
}

func (c *ExpenseController) GetExpensesByStatus(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var expense models.Expense

	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "invalid request body",
		})
		return
	}

	expenses, err := c.ExpenseService.GetExpensesByStatus(ctx, expense.Project.Status)
	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "expenses retrieved successfully",
		"data":    expenses,
	})
}

func (c *ExpenseController) GetExpensesByApprover(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var expense models.Expense

	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": "invalid request body",
		})
		return
	}

	expenses, err := c.ExpenseService.GetExpensesByApprover(ctx, expense.ApprovedBy)

	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "expenses retrieved successfully",
		"data":    expenses,
	})
}

// func (c *ExpenseController) GetExpensesByDateRange(w http.ResponseWriter, r *http.Request) {

// 	ctx := r.Context()

// 	var expense models.Expense

// 	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
// 		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
// 			"status":  "error",
// 			"message": "invalid request body",
// 		})
// 		return
// 	}

// 	expenses, err := c.ExpenseService.GetExpensesByDateRange(ctx, expense., expense.EndDate)

// 	if err != nil {
// 		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
// 			"status":  "error",
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
// 		"status":  "success",
// 		"message": "expenses retrieved successfully",
// 		"data":    expenses,
// 	})
// }

func (c *ExpenseController) GetTotalExpensesByProjectID(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	id, err := utils.ParseInt64Param(r)

	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	
	total, err := c.ExpenseService.GetTotalExpensesByProjectID(ctx, id)

	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "total expenses retrieved successfully",
		"data":    total,
	})
}

func (c *ExpenseController) GetExpensesByProjectID(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	id, err := utils.ParseInt64Param(r)

	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}


	expenses, err := c.ExpenseService.GetExpensesByProjectID(ctx, id)

	if err != nil {
		utils.JSONErrorResponse(w, http.StatusBadRequest, map[string]interface{}{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"message": "expenses retrieved successfully",
		"data":    expenses,
	})
}
