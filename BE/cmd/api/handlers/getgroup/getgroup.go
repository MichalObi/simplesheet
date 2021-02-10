package getgroup

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

func getGroup(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		id := r.Context().Value(models.CtxKey("sheetid"))
		group := &models.Group{SheetID: id.(int)}

		if err := group.GetByID(r.Context(), app); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				w.WriteHeader(http.StatusPreconditionFailed)
				fmt.Fprintf(w, "group does not exist")
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Oops")
			return
		}

		w.Header().Set("Content-Type", "application/json")
		response, _ := json.Marshal(group)
		w.Write(response)
	}
}

func Do(app *application.Application) httprouter.Handle {
	mdw := []middleware.Middleware{
		middleware.LogRequest,
		validateRequest,
	}

	return middleware.Chain(getGroup(app), mdw...)
}
