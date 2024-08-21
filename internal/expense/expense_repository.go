package expense

import (
	"context"
	"database/sql"
	"errors"
	"time"
	models "github.com/BerkatPS/internal"

)

type ExpenseRepository interface {
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

type expenseRepository struct {
	db *sql.DB
}

func NewExpenseRepository(db *sql.DB) ExpenseRepository {
	return &expenseRepository{db: db}
}

func (r *expenseRepository) GetExpensesByProjectID(ctx context.Context, projectID int64) ([]models.Expense, error) {
	query := `
		SELECT id, project_id, description, amount, date, approved_by 
		FROM expenses 
		WHERE project_id = $1
	`
	rows, err := r.db.QueryContext(ctx, query, projectID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var expense models.Expense
		if err := rows.Scan(&expense.ID, &expense.ProjectID, &expense.Description, &expense.Amount, &expense.Date, &expense.ApprovedBy); err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}


func (r *expenseRepository) GetTotalExpensesByProjectID(ctx context.Context, projectID int64) (float64, error) {
	query := `
		SELECT COALESCE(SUM(amount), 0) 
		FROM expenses 
		WHERE project_id = $1
	`
	var total float64
	err := r.db.QueryRowContext(ctx, query, projectID).Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}


func (r *expenseRepository) GetExpensesByDateRange(ctx context.Context, startDate, endDate time.Time) ([]models.Expense, error) {
	query := `
		SELECT id, project_id, description, amount, date, approved_by 
		FROM expenses 
		WHERE date BETWEEN $1 AND $2
	`
	rows, err := r.db.QueryContext(ctx, query, startDate, endDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var expense models.Expense
		if err := rows.Scan(&expense.ID, &expense.ProjectID, &expense.Description, &expense.Amount, &expense.Date, &expense.ApprovedBy); err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}


func (r *expenseRepository) GetExpensesByStatus(ctx context.Context, status string) ([]models.Expense, error) {
	query := `
		SELECT id, project_id, description, amount, date, approved_by 
		FROM expenses 
		WHERE status = $1
	`
	rows, err := r.db.QueryContext(ctx, query, status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var expense models.Expense
		if err := rows.Scan(&expense.ID, &expense.ProjectID, &expense.Description, &expense.Amount, &expense.Date, &expense.ApprovedBy); err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func (r *expenseRepository) GetExpensesByApprover(ctx context.Context, approverID int64) ([]models.Expense, error) {
	query := `
		SELECT id, project_id, description, amount, date, approved_by 
		FROM expenses 
		WHERE approved_by = $1
	`
	rows, err := r.db.QueryContext(ctx, query, approverID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []models.Expense
	for rows.Next() {
		var expense models.Expense
		if err := rows.Scan(&expense.ID, &expense.ProjectID, &expense.Description, &expense.Amount, &expense.Date, &expense.ApprovedBy); err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}



func (e *expenseRepository) CreateExpense(ctx context.Context, expense models.Expense) error {
	query := "INSERT INTO expenses (project_id, amount, description, date) VALUES ($1, $2, $3, $4)"

	_, err := e.db.ExecContext(ctx, query, expense.ProjectID, expense.Amount, expense.Description, expense.Date)
	if err != nil {
		return err
	}
	return nil
}

func (e *expenseRepository) UpdateExpense(ctx context.Context, expense models.Expense) error {
	query := "UPDATE expenses SET amount = $1, description = $2, date = $3 WHERE id = $4"

	_, err := e.db.ExecContext(ctx, query, expense.Amount, expense.Description, expense.Date, expense.ID)

	if err != nil {
		return err
	}
	return nil
}

func (e *expenseRepository) DeleteExpense(ctx context.Context, id int64) error {
	query := "DELETE FROM expenses WHERE id = $1"

	_, err := e.db.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}
	return nil
}

func (e *expenseRepository) GetExpenseById(ctx context.Context, id int64) (models.Expense, error) {
	query := "SELECT id, project_id, amount, description, date FROM expenses WHERE id = $1"

	var expense models.Expense

	row := e.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&expense.ID, &expense.ProjectID, &expense.Amount, &expense.Description, &expense.Date)

	if err != nil {
		return models.Expense{}, err
	}

	if expense.ID == 0 {
		return models.Expense{}, errors.New("expense not found")
	}

	return expense, nil
}