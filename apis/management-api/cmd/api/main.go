package main

import (
	"context"
	"fmt"
	"log/slog"
	"management-api/internal/accounts"
	"net/http"
	"os"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
)

type Options struct {
	Port int `help:"Port to listen on" short:"p" default:"8888"`
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, err.Error())
	}
}

func run() error {
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		mux := chi.NewMux()
		mux.Route("/api", func(r chi.Router) {
			config := huma.DefaultConfig("Managment api", "1.0.0")
			config.Servers = []*huma.Server{
				{URL: "https://example.com/api"},
			}
			api := humachi.New(r, config)
			ar := accounts.NewRouter(api)
			ar.Router()
		})

		server := http.Server{
			Addr:    fmt.Sprintf(":%d", options.Port),
			Handler: mux,
		}

		hooks.OnStart(func() {
			slog.Info("starting server", slog.Int("port", options.Port))
			server.ListenAndServe()
		})

		hooks.OnStop(func() {
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			server.Shutdown(ctx)
		})
	})

	cli.Run()
	return nil
}
