package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

type (
	Guest struct {
		ID    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Date  time.Time `json:"date"`
	}

	GuestResponse struct {
		ID    uuid.UUID `json:"id"`
		Name  string    `json:"name"`
		Email string    `json:"email"`
		Date  time.Time `json:"date"`
	}
)

func (g *Guest) Response() GuestResponse {
	return GuestResponse{
		ID:    g.ID,
		Name:  g.Name,
		Email: g.Email,
		Date:  g.Date,
	}
}

func (g *Guest) Add(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	query := `INSERT INTO "guest" (name,email,date) VALUES ($1,$2,$3) RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), g.Name, g.Email, g.Date).Scan(&g.ID)
	if err != nil {
		return g.ID, err
	}

	return g.ID, nil
}

func (g *Guest) All(ctx context.Context, db *sql.DB, param ListQuery) ([]Guest, error) {
	all := []Guest{}

	query, args, err := param.Query(strings.Split("id,name,email,date", ","))
	if err != nil {
		return all, err
	}

	rows, err := db.QueryContext(ctx, fmt.Sprintf(`SELECT id,name,email,date FROM "guest" %s`, query), args...)
	if err != nil {
		return all, err
	}

	defer rows.Close()

	for rows.Next() {
		one := Guest{}
		err = rows.Scan(
			&one.ID, &one.Name, &one.Email, &one.Date,
		)
		if err != nil {
			return all, err
		}
		all = append(all, one)
	}

	return all, nil
}

func (g *Guest) One(ctx context.Context, db *sql.DB) (Guest, error) {
	one := Guest{}

	query := `SELECT id,name,email,date FROM "guest" WHERE id = $1 LIMIT 1`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), g.ID).Scan(
		&one.ID, &one.Name, &one.Email, &one.Date,
	)
	if err != nil {
		return one, err
	}

	return one, nil
}

func (g *Guest) Update(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `UPDATE "guest" SET name = $1, email = $2,date = $3 WHERE id = $4 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), g.Name, g.Email, g.Date, g.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (g *Guest) Delete(ctx context.Context, db *sql.DB) (uuid.UUID, error) {
	var id uuid.UUID

	query := `DELETE FROM "guest" WHERE id = $1 RETURNING id`
	err := db.QueryRowContext(ctx, fmt.Sprintf(query), g.ID).Scan(&id)
	if err != nil {
		return id, err
	}

	return id, nil
}
