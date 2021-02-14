package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/simplesheet/pkg/application"
)

type Fields struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type FieldsSlice []Fields

type Sheet struct {
	ID     int         `json:"id"`
	Type   string      `json:"type"`
	Fields FieldsSlice `json:"fields"`
}

func (a FieldsSlice) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (s *FieldsSlice) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, s)
	case string:
		return json.Unmarshal([]byte(v), s)
	}
	return errors.New("type assertion failed")
}

func (s *Sheet) Create(ctx context.Context, app *application.Application) error {
	stmt := `
		INSERT INTO sheets (
			type,
      fields
		)
		VALUES ($1, $2)
		RETURNING id
	`
	err := app.DB.Client.QueryRowContext(
		ctx,
		stmt,
		s.Type,
		s.Fields,
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
		&s.Type,
		&s.Fields,
	)

	if err != nil {
		return err
	}

	return nil
}
