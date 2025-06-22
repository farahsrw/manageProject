package routes

import (
	"github.com/gorilla/mux"
	"github.com/justinas/alice"

	"github.com/farahsrw/manageProject/config"
	"github.com/farahsrw/manageProject/handlers"
	"github.com/farahsrw/manageProject/middleware"
)

func SetupRoutes(router *mux.Router, app *config.App, userSchema, projectSchema string) {
	userChain := alice.New(
		middleware.LoggingMiddleware,
		middleware.ValidateMiddleware(userSchema),
	)

	router.Handle("/register", userChain.ThenFunc(handlers.Register(app))).Methods("POST")
	router.Handle("/login", userChain.ThenFunc(handlers.Login(app))).Methods("POST")

	projectChain := alice.New(
		middleware.LoggingMiddleware,
		middleware.JWTMiddleware(app),
	)

	router.Handle("/projects", projectChain.ThenFunc(handlers.GetProjects(app))).Methods("GET")
	router.Handle("/projects/{xata_id}", projectChain.ThenFunc(handlers.GetProject(app))).Methods("GET")

	projectWithValidation := projectChain.Append(middleware.ValidateMiddleware(projectSchema))
	router.Handle("/projects", projectWithValidation.ThenFunc(handlers.CreateProject(app))).Methods("POST")
	router.Handle("/projects/{xata_id}", projectWithValidation.ThenFunc(handlers.UpdateProject(app))).Methods("PUT")
	router.Handle("/projects/{xata_id}", projectChain.ThenFunc(handlers.DeleteProject(app))).Methods("DELETE")
}
