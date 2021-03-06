// Code generated by sqlc. DO NOT EDIT.
// source: transfers.sql

package db

import (
	"context"
)

const createTransfer = `-- name: CreateTransfer :one
INSERT INTO transfers (
  from_account_id,
  to_account_id,
  amount
) VALUES (
  $1, $2,$3
)
RETURNING id, from_account_id, to_account_id, amount, created_at
`

type CreateTransferParams struct {
	FromAccountID int64 `json:"fromAccountID"`
	ToAccountID   int64 `json:"toAccountID"`
	Amount        int64 `json:"amount"`
}

func (q *Queries) CreateTransfer(ctx context.Context, arg CreateTransferParams) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, createTransfer, arg.FromAccountID, arg.ToAccountID, arg.Amount)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const deleteTransfer = `-- name: DeleteTransfer :exec
DELETE FROM transfers
WHERE id = $1
`

func (q *Queries) DeleteTransfer(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteTransfer, id)
	return err
}

const getTransfer = `-- name: GetTransfer :one
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetTransfer(ctx context.Context, id int64) (Transfer, error) {
	row := q.db.QueryRowContext(ctx, getTransfer, id)
	var i Transfer
	err := row.Scan(
		&i.ID,
		&i.FromAccountID,
		&i.ToAccountID,
		&i.Amount,
		&i.CreatedAt,
	)
	return i, err
}

const getTransferByFAI = `-- name: GetTransferByFAI :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE from_account_id = $1
`

func (q *Queries) GetTransferByFAI(ctx context.Context, fromAccountID int64) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, getTransferByFAI, fromAccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
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

const getTransferByTAI = `-- name: GetTransferByTAI :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE to_account_id = $1
`

func (q *Queries) GetTransferByTAI(ctx context.Context, toAccountID int64) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, getTransferByTAI, toAccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
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

const getTransferByTAIandFAI = `-- name: GetTransferByTAIandFAI :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
WHERE from_account_id = $1 and to_account_id =$2
`

type GetTransferByTAIandFAIParams struct {
	FromAccountID int64 `json:"fromAccountID"`
	ToAccountID   int64 `json:"toAccountID"`
}

func (q *Queries) GetTransferByTAIandFAI(ctx context.Context, arg GetTransferByTAIandFAIParams) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, getTransferByTAIandFAI, arg.FromAccountID, arg.ToAccountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
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

const listTransfers = `-- name: ListTransfers :many
SELECT id, from_account_id, to_account_id, amount, created_at FROM transfers
ORDER BY id
LIMIT $1
`

func (q *Queries) ListTransfers(ctx context.Context, limit int32) ([]Transfer, error) {
	rows, err := q.db.QueryContext(ctx, listTransfers, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transfer
	for rows.Next() {
		var i Transfer
		if err := rows.Scan(
			&i.ID,
			&i.FromAccountID,
			&i.ToAccountID,
			&i.Amount,
			&i.CreatedAt,
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
