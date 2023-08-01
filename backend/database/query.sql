-- name: GetExpense :one
SELECT * FROM expenses
WHERE id = ? LIMIT 1;

-- name: ListExpenses :many
SELECT * FROM expenses
ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (
  name, bio
) VALUES (
  ?, ?
)
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors
WHERE id = ?;