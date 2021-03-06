package createsheet

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/simplesheet/cmd/api/models"
	"github.com/simplesheet/pkg/application"
	"github.com/simplesheet/pkg/middleware"
)

func createSheet(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		sheet := &models.Sheet{}
		json.NewDecoder(r.Body).Decode(sheet)

		if err := sheet.Create(r.Context(), app); err != nil {
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
	return middleware.Chain(createSheet(app), middleware.LogRequest)
}
