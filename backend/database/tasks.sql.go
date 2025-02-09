// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: tasks.sql

package database

import (
	"context"
	"database/sql"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (
  type,
  status,
  args,
  results,
  flow_id,
  message,
  tool_call_id
) VALUES (
  ?, ?, ?, ?, ?, ?, ?
)
RETURNING id, created_at, updated_at, type, status, args, results, message, flow_id, tool_call_id
`

type CreateTaskParams struct {
	Type       sql.NullString
	Status     sql.NullString
	Args       sql.NullString
	Results    sql.NullString
	FlowID     sql.NullInt64
	Message    sql.NullString
	ToolCallID sql.NullString
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.Type,
		arg.Status,
		arg.Args,
		arg.Results,
		arg.FlowID,
		arg.Message,
		arg.ToolCallID,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Type,
		&i.Status,
		&i.Args,
		&i.Results,
		&i.Message,
		&i.FlowID,
		&i.ToolCallID,
	)
	return i, err
}

const readTasksByFlowId = `-- name: ReadTasksByFlowId :many
SELECT id, created_at, updated_at, type, status, args, results, message, flow_id, tool_call_id FROM tasks
WHERE flow_id = ?
ORDER BY created_at ASC
`

func (q *Queries) ReadTasksByFlowId(ctx context.Context, flowID sql.NullInt64) ([]Task, error) {
	rows, err := q.db.QueryContext(ctx, readTasksByFlowId, flowID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Task
	for rows.Next() {
		var i Task
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Type,
			&i.Status,
			&i.Args,
			&i.Results,
			&i.Message,
			&i.FlowID,
			&i.ToolCallID,
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

const updateTaskResults = `-- name: UpdateTaskResults :one
UPDATE tasks
SET results = ?
WHERE id = ?
RETURNING id, created_at, updated_at, type, status, args, results, message, flow_id, tool_call_id
`

type UpdateTaskResultsParams struct {
	Results sql.NullString
	ID      int64
}

func (q *Queries) UpdateTaskResults(ctx context.Context, arg UpdateTaskResultsParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTaskResults, arg.Results, arg.ID)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Type,
		&i.Status,
		&i.Args,
		&i.Results,
		&i.Message,
		&i.FlowID,
		&i.ToolCallID,
	)
	return i, err
}

const updateTaskStatus = `-- name: UpdateTaskStatus :one
UPDATE tasks
SET status = ?
WHERE id = ?
RETURNING id, created_at, updated_at, type, status, args, results, message, flow_id, tool_call_id
`

type UpdateTaskStatusParams struct {
	Status sql.NullString
	ID     int64
}

func (q *Queries) UpdateTaskStatus(ctx context.Context, arg UpdateTaskStatusParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTaskStatus, arg.Status, arg.ID)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Type,
		&i.Status,
		&i.Args,
		&i.Results,
		&i.Message,
		&i.FlowID,
		&i.ToolCallID,
	)
	return i, err
}

const updateTaskToolCallId = `-- name: UpdateTaskToolCallId :one
UPDATE tasks
SET tool_call_id = ?
WHERE id = ?
RETURNING id, created_at, updated_at, type, status, args, results, message, flow_id, tool_call_id
`

type UpdateTaskToolCallIdParams struct {
	ToolCallID sql.NullString
	ID         int64
}

func (q *Queries) UpdateTaskToolCallId(ctx context.Context, arg UpdateTaskToolCallIdParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, updateTaskToolCallId, arg.ToolCallID, arg.ID)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Type,
		&i.Status,
		&i.Args,
		&i.Results,
		&i.Message,
		&i.FlowID,
		&i.ToolCallID,
	)
	return i, err
}
