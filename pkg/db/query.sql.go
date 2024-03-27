// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"
)

const createHero = `-- name: CreateHero :one
INSERT INTO hero (
  name
) VALUES (
  $1
)
RETURNING id, name
`

func (q *Queries) CreateHero(ctx context.Context, name string) (Hero, error) {
	row := q.db.QueryRow(ctx, createHero, name)
	var i Hero
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteHero = `-- name: DeleteHero :exec
DELETE FROM hero
WHERE id = $1
`

func (q *Queries) DeleteHero(ctx context.Context, id int32) error {
	_, err := q.db.Exec(ctx, deleteHero, id)
	return err
}

const getHero = `-- name: GetHero :one
SELECT id, name FROM hero
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetHero(ctx context.Context, id int32) (Hero, error) {
	row := q.db.QueryRow(ctx, getHero, id)
	var i Hero
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const listHeros = `-- name: ListHeros :many
SELECT id, name FROM hero
ORDER BY name
`

func (q *Queries) ListHeros(ctx context.Context) ([]Hero, error) {
	rows, err := q.db.Query(ctx, listHeros)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Hero
	for rows.Next() {
		var i Hero
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateHero = `-- name: UpdateHero :exec
UPDATE hero
  set name = $2
WHERE id = $1
`

type UpdateHeroParams struct {
	ID   int32
	Name string
}

func (q *Queries) UpdateHero(ctx context.Context, arg UpdateHeroParams) error {
	_, err := q.db.Exec(ctx, updateHero, arg.ID, arg.Name)
	return err
}
