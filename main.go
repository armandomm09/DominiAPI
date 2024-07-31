package main

import (
	"apiv1/TBA_src/team"
	"fmt"
	"net/http"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
)

func requestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Obtener la dirección IP del cliente
		clientIP := r.RemoteAddr
		if ip := r.Header.Get("X-Real-IP"); ip != "" {
			clientIP = ip
		} else if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
			clientIP = ip
		}

		// Imprimir detalles de la solicitud
		fmt.Println("\n- - - - - - - INCOMING REQUEST - - - - - - - -\n")
		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)
		fmt.Printf("Client IP: %s\n", clientIP)
		fmt.Printf("User Agent: %s\n", r.UserAgent())
		fmt.Printf("Headers:\n")
		for name, values := range r.Header {
			for _, value := range values {
				fmt.Printf("  %s: %s\n", name, value)
			}
		}

		next.ServeHTTP(w, r)

		fmt.Printf("Completed request: %s %s in %v\n", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {

	cli := humacli.New(func(hook humacli.Hooks, options *team.Options) {
		router := chi.NewMux()
		router.Use(requestLogger)

		api := humachi.New(router, huma.DefaultConfig("Domini API", "1.0.0"))

		hook.OnStart(func() {
			fmt.Printf("Starting server on port %d...\n", 8888)
			http.ListenAndServe(fmt.Sprintf(":%d", 8888), router)
		})
		team.AddTeamRegisters(api)

	})

	cli.Run()
}
