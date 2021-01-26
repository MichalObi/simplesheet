package models;

import (
	"context"
	"github.com/simplesheet/pkg/application"
)

type Sheet struct {
  ID int `json:"id"`
	HasMetals bool `json:"has_metals"`
	HasCrypto bool `json:"has_crypto"`
}

func (s *Sheet) Create(ctx context.Context, app *application.Application) error {
	stmt_sheet := `
		INSERT INTO sheets (
			has_metals,
      has_crypto
		)
		VALUES ($1, $2)
		RETURNING id
	`
	err := app.DB.Client.QueryRowContext(
		ctx,
		stmt_sheet,
		s.HasMetals,
    s.HasCrypto,
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
    &s.HasMetals,
    &s.HasCrypto,
	)

	if err != nil {
		return err
	}

	return nil
}
