// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package database

import (
	"context"
	"database/sql"
	"time"
)

const createExpense = `-- name: CreateExpense :one
INSERT INTO expenses (
  id, price, expenseDate
) VALUES (
  ?, ?, ?
)
RETURNING id, price, expensedate
`

type CreateExpenseParams struct {
	ID          int64
	Price       sql.NullFloat64
	Expensedate time.Time
}

func (q *Queries) CreateExpense(ctx context.Context, arg CreateExpenseParams) (Expense, error) {
	row := q.db.QueryRowContext(ctx, createExpense, arg.ID, arg.Price, arg.Expensedate)
	var i Expense
	err := row.Scan(&i.ID, &i.Price, &i.Expensedate)
	return i, err
}

const createIncome = `-- name: CreateIncome :one
INSERT INTO incomes (
  id, price, incomeDate
) VALUES (
  ?, ?, ?
)
RETURNING id, price, incomedate
`

type CreateIncomeParams struct {
	ID         int64
	Price      sql.NullFloat64
	Incomedate time.Time
}

func (q *Queries) CreateIncome(ctx context.Context, arg CreateIncomeParams) (Income, error) {
	row := q.db.QueryRowContext(ctx, createIncome, arg.ID, arg.Price, arg.Incomedate)
	var i Income
	err := row.Scan(&i.ID, &i.Price, &i.Incomedate)
	return i, err
}

const createPayment = `-- name: CreatePayment :one
INSERT INTO payments (
  id, expenseID, price, paidDate
) VALUES (
  ?, ?, ?, ?
)
RETURNING id, expenseid, price, paiddate
`

type CreatePaymentParams struct {
	ID        int64
	Expenseid int64
	Price     float64
	Paiddate  time.Time
}

func (q *Queries) CreatePayment(ctx context.Context, arg CreatePaymentParams) (Payment, error) {
	row := q.db.QueryRowContext(ctx, createPayment,
		arg.ID,
		arg.Expenseid,
		arg.Price,
		arg.Paiddate,
	)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.Expenseid,
		&i.Price,
		&i.Paiddate,
	)
	return i, err
}

const deleteExpense = `-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id = ?
`

func (q *Queries) DeleteExpense(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteExpense, id)
	return err
}

const deleteIncome = `-- name: DeleteIncome :exec
DELETE FROM incomes
WHERE id = ?
`

func (q *Queries) DeleteIncome(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteIncome, id)
	return err
}

const deletePayment = `-- name: DeletePayment :exec
DELETE FROM payments
WHERE id = ?
`

func (q *Queries) DeletePayment(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePayment, id)
	return err
}

const getExpense = `-- name: GetExpense :one
SELECT id, price, expensedate FROM expenses
WHERE id = ? LIMIT 1
`

func (q *Queries) GetExpense(ctx context.Context, id int64) (Expense, error) {
	row := q.db.QueryRowContext(ctx, getExpense, id)
	var i Expense
	err := row.Scan(&i.ID, &i.Price, &i.Expensedate)
	return i, err
}

const getIncome = `-- name: GetIncome :one
SELECT id, price, incomedate FROM incomes
WHERE id = ? LIMIT 1
`

func (q *Queries) GetIncome(ctx context.Context, id int64) (Income, error) {
	row := q.db.QueryRowContext(ctx, getIncome, id)
	var i Income
	err := row.Scan(&i.ID, &i.Price, &i.Incomedate)
	return i, err
}

const getPayment = `-- name: GetPayment :one
SELECT id, expenseid, price, paiddate FROM payments
WHERE id = ? LIMIT 1
`

func (q *Queries) GetPayment(ctx context.Context, id int64) (Payment, error) {
	row := q.db.QueryRowContext(ctx, getPayment, id)
	var i Payment
	err := row.Scan(
		&i.ID,
		&i.Expenseid,
		&i.Price,
		&i.Paiddate,
	)
	return i, err
}

const getPaymentsBetween = `-- name: GetPaymentsBetween :many
SELECT id, expenseid, price, paiddate FROM payments
WHERE paidDate >= ? AND paidDate <= ?
`

type GetPaymentsBetweenParams struct {
	Paiddate   time.Time
	Paiddate_2 time.Time
}

func (q *Queries) GetPaymentsBetween(ctx context.Context, arg GetPaymentsBetweenParams) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, getPaymentsBetween, arg.Paiddate, arg.Paiddate_2)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Payment
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.Expenseid,
			&i.Price,
			&i.Paiddate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getPaymentsByExpense = `-- name: GetPaymentsByExpense :many
SELECT id, expenseid, price, paiddate FROM payments
WHERE expenseID = ?
`

func (q *Queries) GetPaymentsByExpense(ctx context.Context, expenseid int64) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, getPaymentsByExpense, expenseid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Payment
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.Expenseid,
			&i.Price,
			&i.Paiddate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listExpenses = `-- name: ListExpenses :many
SELECT id, price, expensedate FROM expenses
ORDER BY id
`

func (q *Queries) ListExpenses(ctx context.Context) ([]Expense, error) {
	rows, err := q.db.QueryContext(ctx, listExpenses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Expense
	for rows.Next() {
		var i Expense
		if err := rows.Scan(&i.ID, &i.Price, &i.Expensedate); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPayments = `-- name: ListPayments :many
SELECT id, expenseid, price, paiddate FROM payments
ORDER BY id
`

func (q *Queries) ListPayments(ctx context.Context) ([]Payment, error) {
	rows, err := q.db.QueryContext(ctx, listPayments)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Payment
	for rows.Next() {
		var i Payment
		if err := rows.Scan(
			&i.ID,
			&i.Expenseid,
			&i.Price,
			&i.Paiddate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listincomes = `-- name: Listincomes :many
SELECT id, price, incomedate FROM incomes
ORDER BY id
`

func (q *Queries) Listincomes(ctx context.Context) ([]Income, error) {
	rows, err := q.db.QueryContext(ctx, listincomes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Income
	for rows.Next() {
		var i Income
		if err := rows.Scan(&i.ID, &i.Price, &i.Incomedate); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateExpense = `-- name: UpdateExpense :exec
UPDATE expenses
SET price = ?, expenseDate = ?
WHERE id = ?
`

type UpdateExpenseParams struct {
	Price       sql.NullFloat64
	Expensedate time.Time
	ID          int64
}

func (q *Queries) UpdateExpense(ctx context.Context, arg UpdateExpenseParams) error {
	_, err := q.db.ExecContext(ctx, updateExpense, arg.Price, arg.Expensedate, arg.ID)
	return err
}

const updateIncome = `-- name: UpdateIncome :exec
UPDATE incomes
SET price = ?, incomeDate = ?
WHERE id = ?
`

type UpdateIncomeParams struct {
	Price      sql.NullFloat64
	Incomedate time.Time
	ID         int64
}

func (q *Queries) UpdateIncome(ctx context.Context, arg UpdateIncomeParams) error {
	_, err := q.db.ExecContext(ctx, updateIncome, arg.Price, arg.Incomedate, arg.ID)
	return err
}

const updatePayment = `-- name: UpdatePayment :exec
UPDATE payments
SET price = ?, paidDate = ?
WHERE id = ?
`

type UpdatePaymentParams struct {
	Price    float64
	Paiddate time.Time
	ID       int64
}

func (q *Queries) UpdatePayment(ctx context.Context, arg UpdatePaymentParams) error {
	_, err := q.db.ExecContext(ctx, updatePayment, arg.Price, arg.Paiddate, arg.ID)
	return err
}
