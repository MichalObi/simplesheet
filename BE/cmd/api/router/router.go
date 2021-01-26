package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/simplesheet/cmd/api/handlers/creategroup"
	"github.com/simplesheet/cmd/api/handlers/createsheet"
	"github.com/simplesheet/cmd/api/handlers/createuser"
	"github.com/simplesheet/cmd/api/handlers/getuser"
	"github.com/simplesheet/pkg/application"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()

	mux.GET("/users/:id", getuser.Do(app))
	mux.POST("/users", createuser.Do(app))
	mux.POST("/sheets", createsheet.Do(app))
	mux.POST("/groups", creategroup.Do(app))

	return mux
}
