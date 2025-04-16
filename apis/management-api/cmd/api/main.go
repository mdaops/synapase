package main

import (
	"context"
	"flag"
	"fmt"
	"management-api/internal/accounts"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	"github.com/go-chi/chi/v5"
)

var addr = flag.String("addr", ":8080", "http service address")

type Response struct {
	Body struct {
		Message string `json:"message" example:"test" descripton:"another test"`
	} `json:"body"`
}

func main() {
	flag.Parse()

	router := chi.NewMux()
	api := humachi.New(router, huma.DefaultConfig("Management API", "1.0.0"))
	ar := accounts.NewRouter(api)
	ar.Router()
	huma.Register(api, huma.Operation{
		OperationID: "get-greeting",
		Method:      http.MethodGet,
		Path:        "/greeting/{name}",
		Summary:     "Get a greeting",
		Description: "Get a greeting for a person by name.",
		Tags:        []string{"Greetings"},
	}, func(ctx context.Context, input *struct {
		Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	},
	) (*Response, error) {
		resp := &Response{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	})

	http.ListenAndServe("127.0.0.1:8888", router)
}
