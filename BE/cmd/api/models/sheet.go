package models;

import (
	"context"
	"github.com/simplesheet/pkg/application"
)

type Sheet struct {
  ID       int `json:"id"`
  Field string `json:"field"`
  Value string `json:"value"`
}

func (s *Sheet) Create(ctx context.Context, app *application.Application) error {
	stmt := `
		INSERT INTO sheets (
			field,
      value
		)
		VALUES ($1, $2)
		RETURNING id
	`
	err := app.DB.Client.QueryRowContext(
		ctx,
		stmt,
		s.Field,
    s.Value,
	).Scan(&s.ID)

	if err != nil {
		return err
	}

	return nil
}

func (s *Sheet) GetByID(ctx context.Context, app *application.Application) error {
	stmt := `
		SELECT *
		FROM sheets
		WHERE id = $1
	`
	err := app.DB.Client.QueryRowContext(
		ctx,
		stmt,
		s.ID,
	).Scan(
		&s.ID,
    s.Field,
    s.Value,
	)

	if err != nil {
		return err
	}

	return nil
}
