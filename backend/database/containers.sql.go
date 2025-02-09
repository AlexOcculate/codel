// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: containers.sql

package database

import (
	"context"
	"database/sql"
)

const createContainer = `-- name: CreateContainer :one
INSERT INTO containers (
  name, image, status
)
VALUES (
  ?, ?, ?
)
RETURNING id, name, local_id, image, status
`

type CreateContainerParams struct {
	Name   sql.NullString
	Image  sql.NullString
	Status sql.NullString
}

func (q *Queries) CreateContainer(ctx context.Context, arg CreateContainerParams) (Container, error) {
	row := q.db.QueryRowContext(ctx, createContainer, arg.Name, arg.Image, arg.Status)
	var i Container
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LocalID,
		&i.Image,
		&i.Status,
	)
	return i, err
}

const getAllRunningContainers = `-- name: GetAllRunningContainers :many
SELECT id, name, local_id, image, status FROM containers WHERE status = 'running'
`

func (q *Queries) GetAllRunningContainers(ctx context.Context) ([]Container, error) {
	rows, err := q.db.QueryContext(ctx, getAllRunningContainers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Container
	for rows.Next() {
		var i Container
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.LocalID,
			&i.Image,
			&i.Status,
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

const updateContainerLocalId = `-- name: UpdateContainerLocalId :one
UPDATE containers
SET local_id = ?
WHERE id = ?
RETURNING id, name, local_id, image, status
`

type UpdateContainerLocalIdParams struct {
	LocalID sql.NullString
	ID      int64
}

func (q *Queries) UpdateContainerLocalId(ctx context.Context, arg UpdateContainerLocalIdParams) (Container, error) {
	row := q.db.QueryRowContext(ctx, updateContainerLocalId, arg.LocalID, arg.ID)
	var i Container
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LocalID,
		&i.Image,
		&i.Status,
	)
	return i, err
}

const updateContainerStatus = `-- name: UpdateContainerStatus :one
UPDATE containers
SET status = ?
WHERE id = ?
RETURNING id, name, local_id, image, status
`

type UpdateContainerStatusParams struct {
	Status sql.NullString
	ID     int64
}

func (q *Queries) UpdateContainerStatus(ctx context.Context, arg UpdateContainerStatusParams) (Container, error) {
	row := q.db.QueryRowContext(ctx, updateContainerStatus, arg.Status, arg.ID)
	var i Container
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.LocalID,
		&i.Image,
		&i.Status,
	)
	return i, err
}
