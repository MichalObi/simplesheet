package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/simplesheet/pkg/application"
)

type Position struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type PositionSlice []Position

type Group struct {
	GroupID   int           `json:"group_id"`
	SheetId   int           `json:"sheet_id"`
	Name      string        `json:"name"`
	Positions PositionSlice `json:"positions"`
}

func (a PositionSlice) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (s *PositionSlice) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, s)
	case string:
		return json.Unmarshal([]byte(v), s)
	}
	return errors.New("type assertion failed")
}

func (g *Group) Create(ctx context.Context, app *application.Application) error {

	stmt := `INSERT INTO groups
					(
						sheets_id,
						name,
						positions
					)
					VALUES ($1, $2, $3)
					RETURNING group_id`

	err := app.DB.Client.QueryRowContext(
		ctx,
		stmt,
		g.SheetId,
		g.Name,
		g.Positions,
	).Scan(&g.GroupID)

	if err != nil {
		return err
	}

	return nil
}

func (g *Group) GetByID(ctx context.Context, app *application.Application) error {
	stmt := `
		SELECT *
		FROM groups
		WHERE group_id = $1
	`
	err := app.DB.Client.QueryRowContext(
		ctx,
		stmt,
		g.GroupID,
	).Scan(
		&g.GroupID,
		&g.SheetId,
		&g.Name,
		&g.Positions,
	)

	if err != nil {
		return err
	}

	return nil
}
