package main

import (
	"apiv1/exampleRoutes"
	"apiv1/tba_source/team"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
)

func main() {

	cli := humacli.New(func(hook humacli.Hooks, options *exampleRoutes.Options) {
		router := chi.NewMux()
		api := humachi.New(router, huma.DefaultConfig("My First API", "1.0.0"))

		hook.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", 8888)
			http.ListenAndServe(fmt.Sprintf(":%d", 8888), router)
		})
		team.AddTeamRegisters(api)

		exampleRoutes.AddRoutes(api)
	})

	cli.Run()
}
