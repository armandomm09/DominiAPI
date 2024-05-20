package exampleRoutes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/danielgtaylor/huma/v2"
)

func AddRoutes(api huma.API) {
	huma.Register(api, huma.Operation{
		OperationID: "get-greeting",
		Method:      http.MethodGet,
		Path:        "/greeting/{name}",
		Summary:     "Get a greeting",
		Description: "Get a greeting for a person by name.",
		Tags:        []string{"Greetings"},
	}, func(ctx context.Context, input *struct {
		Name string `path:"name" maxLength:"10" example:"world" doc:"Name to greet"`
	}) (*GreetingOutput, error) {
		resp := &GreetingOutput{}
		resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
		return resp, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "review-post",
		Method:      http.MethodPost,
		Path:        "/reviews",
		Summary:     "Review a post",
		Tags:        []string{"Review"},

		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, input *ReviewInput) (*StringOutput, error) {
		resp := StringOutput{}
		resp.Body.Word = "DOS"
		return &resp, nil
	})

	huma.Register(api, huma.Operation{
		OperationID: "sum-post",
		Method:      http.MethodPost,
		Path:        "/sum",
		Summary:     "Sum 2 numbers",
		Tags:        []string{"Math"},

		DefaultStatus: http.StatusCreated,
	}, func(ctx context.Context, input *SumNumbersStruct) (*SumNumbersOutput, error) {
		output := SumNumbersOutput{
			input.Body.Num1 + input.Body.Num2,
		}

		return &output, nil
	})

}
