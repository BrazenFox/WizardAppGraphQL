package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/BrazenFox/WizardAppGraphQL/graph/generated"
	"github.com/BrazenFox/WizardAppGraphQL/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) LoginAuth(ctx context.Context, input model.LoginRequest) (*model.LoginResponse, error) {
	url := "http://192.168.99.102:8080/auth/login"
	requestBody, _ := json.Marshal(input)
	body := bytes.NewBuffer(requestBody)
	log.Println("*************************************************************************************************************************************************")
	log.Println("Request body: ", body)
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	log.Println("Status response:", resp.Status)
	if resp.StatusCode == http.StatusOK {
		bodyResponse, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var response model.LoginResponse
		json.Unmarshal(bodyResponse, &response)
		log.Println("Response body: ", response)
		return &response, nil

	} else {
		bodyResponse, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var badResponse model.BadResponse
		json.Unmarshal(bodyResponse, &badResponse)
		log.Println("Bad response body: ", badResponse)

		graphql.AddError(ctx, &gqlerror.Error{
			Path:       graphql.GetPath(ctx),
			Message:    badResponse.Message,
			Extensions: map[string]interface{}{badResponse.Error: badResponse.Status},
		})
		return nil, err
	}
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) LogoutAuth(ctx context.Context) (string, error) {
	return "Logout", nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) LoginTest1(ctx context.Context, input model.LoginRequest) (string, error) {
	panic(fmt.Errorf("not implemented"))
}
