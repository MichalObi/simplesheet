package getsheet

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/simplesheet/cmd/api/models"
	"github.com/simplesheet/pkg/application"
	"github.com/simplesheet/pkg/middleware"
)

func getSheet(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		id := r.Context().Value(models.CtxKey("sheetid"))
		sheet := &models.Sheet{ID: id.(int)}

		if err := sheet.GetByID(r.Context(), app); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusPreconditionFailed)
				fmt.Fprintf(w, "sheet does not exist")
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Oops")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(sheet)
		w.Write(response)
	}
}

func Do(app *application.Application) httprouter.Handle {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		validateRequest,
	}

	return middleware.Chain(getSheet(app), mdw...)
}
