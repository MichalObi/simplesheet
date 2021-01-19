package router

import (
	"github.com/simplesheet/BE/cmd/api/handlers/createuser"
	"github.com/simplesheet/BE/cmd/api/handlers/getuser"
	"github.com/simplesheet/BE/cmd/api/handlers/createsheet"
	"github.com/simplesheet/BE/pkg/application"
	"github.com/julienschmidt/httprouter"
)

func Get(app *application.Application) *httprouter.Router {
	mux := httprouter.New()
	mux.GET("/users/:id", getuser.Do(app))
	mux.POST("/users", createuser.Do(app))
	mux.POST("/sheets", createsheet.Do(app))
	return mux
}
