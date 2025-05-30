// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const deleteUsersByName = `-- name: DeleteUsersByName :execrows
DELETE FROM users WHERE first_name = $1 AND last_name = $2
`

type DeleteUsersByNameParams struct {
	FirstName string
	LastName  sql.NullString
}

func (q *Queries) DeleteUsersByName(ctx context.Context, arg DeleteUsersByNameParams) (int64, error) {
	result, err := q.exec(ctx, q.deleteUsersByNameStmt, deleteUsersByName, arg.FirstName, arg.LastName)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getUserByID = `-- name: GetUserByID :one
SELECT first_name, id, last_name FROM users WHERE id = $1
`

type GetUserByIDRow struct {
	FirstName string
	ID        int32
	LastName  sql.NullString
}

func (q *Queries) GetUserByID(ctx context.Context, targetID int32) (GetUserByIDRow, error) {
	row := q.queryRow(ctx, q.getUserByIDStmt, getUserByID, targetID)
	var i GetUserByIDRow
	err := row.Scan(&i.FirstName, &i.ID, &i.LastName)
	return i, err
}

const insertNewUser = `-- name: InsertNewUser :exec
INSERT INTO users (first_name, last_name) VALUES ($1, $2)
`

type InsertNewUserParams struct {
	FirstName string
	LastName  sql.NullString
}

func (q *Queries) InsertNewUser(ctx context.Context, arg InsertNewUserParams) error {
	_, err := q.exec(ctx, q.insertNewUserStmt, insertNewUser, arg.FirstName, arg.LastName)
	return err
}

const insertNewUserWithResult = `-- name: InsertNewUserWithResult :execresult
INSERT INTO users (first_name, last_name) VALUES ($1, $2)
`

type InsertNewUserWithResultParams struct {
	FirstName string
	LastName  sql.NullString
}

func (q *Queries) InsertNewUserWithResult(ctx context.Context, arg InsertNewUserWithResultParams) (sql.Result, error) {
	return q.exec(ctx, q.insertNewUserWithResultStmt, insertNewUserWithResult, arg.FirstName, arg.LastName)
}

const listUsers = `-- name: ListUsers :many
SELECT first_name, last_name FROM users
`

type ListUsersRow struct {
	FirstName string
	LastName  sql.NullString
}

func (q *Queries) ListUsers(ctx context.Context) ([]ListUsersRow, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListUsersRow
	for rows.Next() {
		var i ListUsersRow
		if err := rows.Scan(&i.FirstName, &i.LastName); err != nil {
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
