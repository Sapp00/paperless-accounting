-- name: CreateExpense :one
INSERT INTO expenses (
  id, price, expenseDate
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: GetExpense :one
SELECT * FROM expenses
WHERE id = ? LIMIT 1;

-- name: ListExpenses :many
SELECT * FROM expenses
ORDER BY id;

-- name: DeleteExpense :exec
DELETE FROM expenses
WHERE id = ?;



-- name: CreatePayment :one
INSERT INTO payments (
  expenseID, price, paidDate
) VALUES (
  ?, ?, ?
)
RETURNING *;

-- name: GetPayment :one
SELECT * FROM payments
WHERE id = ? LIMIT 1;

-- name: ListPayments :many
SELECT * FROM payments
ORDER BY id;

-- name: DeletePayment :exec
DELETE FROM payments
WHERE id = ?;



-- name: CreateIncome :one
INSERT INTO incomes (
  price, incomeDate
) VALUES (
  ?, ?
)
RETURNING *;

-- name: GetIncome :one
SELECT * FROM incomes
WHERE id = ? LIMIT 1;

-- name: Listincomes :many
SELECT * FROM incomes
ORDER BY id;

-- name: DeleteIncome :exec
DELETE FROM incomes
WHERE id = ?;