package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/BrazenFox/WizardAppGraphQL/graph/generated"
	"github.com/BrazenFox/WizardAppGraphQL/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) LoginAuth(ctx context.Context, input model.LoginRequest) (*model.LoginResponse, error) {
	url := "http://192.168.99.102:8080/auth/login"
	requestBody, _ := json.Marshal(input)
	body := bytes.NewBuffer(requestBody)
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
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
		graphql.AddError(ctx, &gqlerror.Error{
			Path:       graphql.GetPath(ctx),
			Message:    badResponse.Message,
			Extensions: map[string]interface{}{badResponse.Error: badResponse.Status},
		})
		return nil, err
	}
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (string, error) {
	url := "http://192.168.99.102:8080/user/create"
	requestBody, _ := json.Marshal(input)
	body := bytes.NewBuffer(requestBody)
	req, err := http.NewRequest("POST", url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", input.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode == http.StatusOK {
		return "Ok", nil

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
		return "nil", err
	}
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (string, error) {
	url := "http://192.168.99.102:8080/user/update/" + strconv.Itoa(input.ID)
	requestBody, _ := json.Marshal(input)
	body := bytes.NewBuffer(requestBody)
	req, err := http.NewRequest("PUT", url, body)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", input.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode == http.StatusOK {
		return "Ok", nil

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
		return "nil", err
	}
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input *model.UserDeleteByID) (string, error) {
	url := "http://192.168.99.102:8080/user/delete/" + strconv.Itoa(input.ID)
	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", input.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode == http.StatusOK {
		return "Ok", nil

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
		return "nil", err
	}
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.todos, nil
}

func (r *queryResolver) LogoutAuth(ctx context.Context) (string, error) {
	url := "http://192.168.99.102:8080/auth/logout"
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}
	log.Println(resp)
	return "Logout", nil
}

func (r *queryResolver) GetUsers(ctx context.Context, input model.Token) ([]*model.User, error) {
	url := "http://192.168.99.102:8080/user/find"
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", input.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyResponse, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var users []*model.User
		json.Unmarshal(bodyResponse, &users)
		return users, nil
	} else {
		bodyResponse, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var badResponse model.BadResponse
		json.Unmarshal(bodyResponse, &badResponse)
		graphql.AddError(ctx, &gqlerror.Error{
			Path:       graphql.GetPath(ctx),
			Message:    badResponse.Message,
			Extensions: map[string]interface{}{badResponse.Error: badResponse.Status},
		})
		return nil, err
	}
}

func (r *queryResolver) GetUser(ctx context.Context, input model.UserFindByID) (*model.User, error) {
	url := "http://192.168.99.102:8080/user/find/" + strconv.Itoa(input.ID)
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", input.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyResponse, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var user model.User
		json.Unmarshal(bodyResponse, &user)
		return &user, nil
	} else {
		bodyResponse, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var badResponse model.BadResponse
		json.Unmarshal(bodyResponse, &badResponse)
		graphql.AddError(ctx, &gqlerror.Error{
			Path:       graphql.GetPath(ctx),
			Message:    badResponse.Message,
			Extensions: map[string]interface{}{badResponse.Error: badResponse.Status},
		})
		return nil, err
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
