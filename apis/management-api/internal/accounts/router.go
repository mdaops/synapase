package accounts

import (
	"context"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

type router struct {
	api huma.API
}

type AccountResponse struct {
	Body struct {
		Message string `json:"message" example:"test" descripton:"another test"`
	} `json:"body"`
}

type AccountsResponse struct {
	Body struct {
		Message string `json:"message" example:"test" descripton:"another test"`
	} `json:"body"`
}

type GetAccountInput struct {
	ID string `path:"id" doc:"ID of account"`
}

func (r *router) Accounts(ctx context.Context, input *struct{}) (*AccountsResponse, error) {
	resp := &AccountsResponse{}
	resp.Body.Message = "Hello, world!"
	return resp, nil
}

func (r *router) Account(ctx context.Context, input GetAccountInput) (*AccountResponse, error) {
	resp := &AccountResponse{}
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

	huma.Register(r.api, huma.Operation{
		OperationID: "get-account",
		Method:      http.MethodGet,
		Path:        "/accounts/{id}",
		Summary:     "Get account",
		Description: "Get account by id.",
		Tags:        []string{"Accounts"},
	}, r.Accounts)
}

func NewRouter(api huma.API) *router {
	return &router{
		api: api,
	}
}
