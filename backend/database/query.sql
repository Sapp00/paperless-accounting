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

-- name: UpdateExpense :exec
UPDATE expenses
SET price = ?, expenseDate = ?
WHERE id = ?;


-- name: CreatePayment :one
INSERT INTO payments (
  id, expenseID, price, paidDate
) VALUES (
  ?, ?, ?, ?
)
RETURNING *;

-- name: GetPayment :one
SELECT * FROM payments
WHERE id = ? LIMIT 1;

-- name: GetPaymentsByExpense :many
SELECT * FROM payments
WHERE expenseID = ?;

-- name: GetPaymentsBetween :many
SELECT * FROM payments
WHERE paidDate >= ? AND paidDate <= ?;

-- name: ListPayments :many
SELECT * FROM payments
ORDER BY id;

-- name: DeletePayment :exec
DELETE FROM payments
WHERE id = ?;

-- name: UpdatePayment :exec
UPDATE payments
SET price = ?, paidDate = ?
WHERE id = ?;


-- name: CreateIncome :one
INSERT INTO incomes (
  id, price, incomeDate
) VALUES (
  ?, ?, ?
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

-- name: UpdateIncome :exec
UPDATE incomes
SET price = ?, incomeDate = ?
WHERE id = ?;