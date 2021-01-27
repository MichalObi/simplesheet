package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/simplesheet/cmd/api/handlers/creategroup"
	"github.com/simplesheet/cmd/api/handlers/createsheet"
	"github.com/simplesheet/cmd/api/handlers/createuser"
	"github.com/simplesheet/cmd/api/handlers/getgroup"
	"github.com/simplesheet/cmd/api/handlers/getsheet"
	"github.com/simplesheet/cmd/api/handlers/getuser"
	"github.com/simplesheet/pkg/application"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()

	mux.POST("/users", createuser.Do(app))
	mux.POST("/sheets", createsheet.Do(app))
	mux.POST("/groups", creategroup.Do(app))

	mux.GET("/users/:id", getuser.Do(app))
	mux.GET("/sheets/:id", getsheet.Do(app))
	mux.GET("/groups/:id", getgroup.Do(app))

	return mux
}
