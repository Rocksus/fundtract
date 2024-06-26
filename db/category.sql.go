// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: category.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const insertCategory = `-- name: InsertCategory :one
INSERT INTO category (
    category_name,
    created_at,
    updated_at
) VALUES ($1, $2, $3)
RETURNING category_id, category_name, parent_category_id, created_at, updated_at
`

type InsertCategoryParams struct {
	CategoryName string
	CreatedAt    pgtype.Timestamptz
	UpdatedAt    pgtype.Timestamptz
}

func (q *Queries) InsertCategory(ctx context.Context, arg InsertCategoryParams) (Category, error) {
	row := q.db.QueryRow(ctx, insertCategory, arg.CategoryName, arg.CreatedAt, arg.UpdatedAt)
	var i Category
	err := row.Scan(
		&i.CategoryID,
		&i.CategoryName,
		&i.ParentCategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listCategories = `-- name: ListCategories :many
SELECT 
    category_id,
    category_name,
    parent_category_id,
    created_at,
    updated_at
FROM category
`

func (q *Queries) ListCategories(ctx context.Context) ([]Category, error) {
	rows, err := q.db.Query(ctx, listCategories)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Category
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.CategoryID,
			&i.CategoryName,
			&i.ParentCategoryID,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCategoryByID = `-- name: UpdateCategoryByID :exec
UPDATE category set category_name=$1, updated_at=NOW()
WHERE category_id=$2
`

type UpdateCategoryByIDParams struct {
	CategoryName string
	CategoryID   int64
}

func (q *Queries) UpdateCategoryByID(ctx context.Context, arg UpdateCategoryByIDParams) error {
	_, err := q.db.Exec(ctx, updateCategoryByID, arg.CategoryName, arg.CategoryID)
	return err
}
