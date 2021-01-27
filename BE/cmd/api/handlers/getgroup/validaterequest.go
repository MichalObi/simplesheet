package getgroup

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/simplesheet/cmd/api/models"
)

func validateRequest(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		gid := p.ByName("id")

		id, err := strconv.Atoi(gid)
		if err != nil {
			w.WriteHeader(http.StatusPreconditionFailed)
			fmt.Fprintf(w, "malformed id")
			return
		}

		ctx := context.WithValue(r.Context(), models.CtxKey("groupid"), id)
		r = r.WithContext(ctx)
		next(w, r, p)
	}
}
