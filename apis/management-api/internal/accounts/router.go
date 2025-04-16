package accounts

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type router struct {
	api huma.API
}

type AccountsResponse struct {
	Body struct {
		Message string `json:"message" example:"test" descripton:"another test"`
	} `json:"body"`
}

func (r *router) Accounts(ctx context.Context, input *struct{}) (*AccountsResponse, error) {
	resp := &AccountsResponse{}
	resp.Body.Message = "Hello, world!"
	return resp, nil
}

func (r *router) Router() {
	huma.Register(r.api, huma.Operation{
		OperationID: "get-accounts",
		Method:      http.MethodGet,
		Path:        "/accounts",
		Summary:     "Get accounts",
		Description: "Get accounts.",
		Tags:        []string{"Accounts"},
	}, r.Accounts)
}

func NewRouter(api huma.API) *router {
	return &router{
		api: api,
	}
}
