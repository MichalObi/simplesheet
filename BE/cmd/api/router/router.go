package router

import (
	"github.com/simplesheet/cmd/api/handlers/createuser"
	"github.com/simplesheet/cmd/api/handlers/getuser"
	"github.com/simplesheet/cmd/api/handlers/createsheet"
	"github.com/simplesheet/cmd/api/handlers/creategroup"
	"github.com/simplesheet/pkg/application"
	"github.com/julienschmidt/httprouter"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/users/:id", getuser.Do(app))
	mux.POST("/users", createuser.Do(app))
	mux.POST("/sheets", createsheet.Do(app))
	mux.POST("/groups", creategroup.Do(app))
	return mux
}
