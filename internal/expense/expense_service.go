package expense

import (
	"context"
	"errors"
	"time"
	models "github.com/BerkatPS/internal"

)

type ExpenseService interface {
	CreateExpense(ctx context.Context, expense models.Expense) error
	UpdateExpense(ctx context.Context, expense models.Expense) error
	DeleteExpense(ctx context.Context, id int64) error
	GetExpenseById(ctx context.Context, id int64) (models.Expense, error)
	GetExpensesByStatus(ctx context.Context, status string) ([]models.Expense, error)
	GetExpensesByApprover(ctx context.Context, approverID int64) ([]models.Expense, error)
	GetExpensesByDateRange(ctx context.Context, startDate, endDate time.Time) ([]models.Expense, error)
	GetTotalExpensesByProjectID(ctx context.Context, projectID int64) (float64, error)
	GetExpensesByProjectID(ctx context.Context, projectID int64) ([]models.Expense, error)
}

type expenseService struct {
	ExpenseRepo ExpenseRepository
}

func NewExpenseService(expenseRepo ExpenseRepository) ExpenseService {
	return &expenseService{ExpenseRepo: expenseRepo}
}

func (s *expenseService) CreateExpense(ctx context.Context, expense models.Expense) error {
	if expense.Amount <= 0 {
		return errors.New("expense amount must be greater than zero")
	}
	if expense.Description == "" {
		return errors.New("expense description is required")
	}
	if expense.ProjectID <= 0 {
		return errors.New("project id is required")
	}
	if expense.Date.IsZero() {
		return errors.New("expense date is required")
	}
	if expense.ApprovedBy <= 0 {
		return errors.New("approver id is required")
	}
	return s.ExpenseRepo.CreateExpense(ctx, expense)
}

func (s *expenseService) UpdateExpense(ctx context.Context, expense models.Expense) error {

	if expense.Amount <= 0 {
		return errors.New("expense amount must be greater than zero")
	}

	if expense.Description == "" {
		return errors.New("expense description is required")
	}

	if expense.ProjectID <= 0 {
		return errors.New("project id is required")
	}

	if expense.Date.IsZero() {
		return errors.New("expense date is required")
	}

	if expense.ApprovedBy <= 0 {
		return errors.New("approver id is required")
	}

	return s.ExpenseRepo.UpdateExpense(ctx, expense)
}

func (s *expenseService) DeleteExpense(ctx context.Context, id int64) error {
	if id <= 0 {
		return errors.New("expense id is required")
	}
	return s.ExpenseRepo.DeleteExpense(ctx, id)
}

func (s *expenseService) GetExpenseById(ctx context.Context, id int64) (models.Expense, error) {
	if id <= 0 {
		return models.Expense{}, errors.New("expense id is required")
	}

	return s.ExpenseRepo.GetExpenseById(ctx, id)
}

func (s *expenseService) GetExpensesByStatus(ctx context.Context, status string) ([]models.Expense, error) {
	if status == "" {
		return nil, errors.New("expense status is required")
	}

	return s.ExpenseRepo.GetExpensesByStatus(ctx, status)
}

func (s *expenseService) GetExpensesByApprover(ctx context.Context, approverID int64) ([]models.Expense, error) {
	if approverID <= 0 {
		return nil, errors.New("approver id is required")
	}
	return s.ExpenseRepo.GetExpensesByApprover(ctx, approverID)
}

func (s *expenseService) GetExpensesByDateRange(ctx context.Context, startDate, endDate time.Time) ([]models.Expense, error) {
	if startDate.IsZero() {
		return nil, errors.New("start date is required")
	}

	if endDate.IsZero() {
		return nil, errors.New("end date is required")
	}
	return s.ExpenseRepo.GetExpensesByDateRange(ctx, startDate, endDate)
}

func (s *expenseService) GetTotalExpensesByProjectID(ctx context.Context, projectID int64) (float64, error) {
	if projectID <= 0 {
		return 0, errors.New("project id is required")
	}

	return s.ExpenseRepo.GetTotalExpensesByProjectID(ctx, projectID)
}

func (s *expenseService) GetExpensesByProjectID(ctx context.Context, projectID int64) ([]models.Expense, error) {
	if projectID <= 0 {
		return nil, errors.New("project id is required")
	}
	return s.ExpenseRepo.GetExpensesByProjectID(ctx, projectID)
}
