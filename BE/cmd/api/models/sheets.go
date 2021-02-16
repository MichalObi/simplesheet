package models

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/simplesheet/pkg/application"
)

type SheetSlice []*Sheet

type AllSheets struct {
	Sheets SheetSlice `json:"sheets"`
}

func (a SheetSlice) Value() (driver.Value, error) {
	return json.Marshal(a)
}

func (s *SheetSlice) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		return json.Unmarshal(v, s)
	case string:
		return json.Unmarshal([]byte(v), s)
	}
	return errors.New("type assertion failed")
}

func (s *AllSheets) GetAll(ctx context.Context, app *application.Application) error {
	stmt := `
		SELECT *
		FROM sheets
	`
	rows, err := app.DB.Client.QueryContext(
		ctx,
		stmt,
	)

	if err != nil {
		return err
	}

	defer rows.Close()

	sheets := SheetSlice{}

	for rows.Next() {
		r := new(Sheet)
		err := rows.Scan(&r.ID, &r.Type, &r.Fields)

		if err != nil {
			return err
		}

		sheets = append(sheets, r)
	}

	s.Sheets = sheets

	return nil
}
