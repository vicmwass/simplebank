-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id,
  to_account_id,
  amount
) VALUES (
  $1, $2,$3
)
RETURNING *;

-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: GetTransferByFAI :many
SELECT * FROM transfers
WHERE from_account_id = $1 ;

-- name: GetTransferByTAI :many
SELECT * FROM transfers
WHERE to_account_id = $1 ;

-- name: GetTransferByTAIandFAI :many
SELECT * FROM transfers
WHERE from_account_id = $1 and to_account_id =$2 ;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1;


-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1;