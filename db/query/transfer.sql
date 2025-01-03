-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: GetTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: CreateTransfer :one
INSERT INTO transfers (
  owner,
  balance,
  currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateTransfer :one
UPDATE transfers
  SET balance = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTransfer :one
DELETE FROM transfers
WHERE id = $1
RETURNING *;