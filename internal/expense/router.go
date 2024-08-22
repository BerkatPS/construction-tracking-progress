package expense

import "net/http"


func RegisterRoutes(router *http.ServeMux, handler *ExpenseController) {
	router.HandleFunc("POST /expenses", handler.CreateExpense)
	router.HandleFunc("PUT /expenses", handler.UpdateExpense)
	router.HandleFunc("DELETE /expenses", handler.DeleteExpense)
	router.HandleFunc("GET /expenses", handler.GetExpenseById)
	router.HandleFunc("GET /expenses/approver", handler.GetExpensesByApprover)
	router.HandleFunc("GET /expenses/total/{projectID}", handler.GetTotalExpensesByProjectID)
	router.HandleFunc("GET /expenses/status", handler.GetExpensesByStatus)
	router.HandleFunc("GET /expenses/project", handler.GetExpensesByProjectID)
}