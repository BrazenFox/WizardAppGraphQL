package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/magiconair/properties"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/99designs/gqlgen/graphql"
	"github.com/BrazenFox/WizardAppGraphQL/graph/generated"
	"github.com/BrazenFox/WizardAppGraphQL/graph/model"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

var URL = GetUrl()

func GetUrl() string {
	p := properties.MustLoadFile("config.conf", properties.UTF8)
	url := p.GetString("url", "http://192.168.99.102:8080")
	if url == "" {
		return "http://192.168.99.102:8080"
	} else {
		return url
	}

}

func (r *mutationResolver) LoginAuth(ctx context.Context, input model.LoginRequest) (*model.LoginResponse, error) {
	url := URL + "/auth/login"
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
		ErrorResponse(ctx, resp)
		return nil, err
	}
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUser) (string, error) {
	url := URL + "/user/create"
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
		ErrorResponse(ctx, resp)
		return "nil", err
	}
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (string, error) {
	url := URL + "/user/update/" + strconv.Itoa(input.ID)
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
		ErrorResponse(ctx, resp)
		return "nil", err
	}
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input *model.RequestWithID) (string, error) {
	url := URL + "/user/delete/" + strconv.Itoa(input.ID)
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
		ErrorResponse(ctx, resp)
		return "nil", err
	}
}

func (r *mutationResolver) CreateWizard(ctx context.Context, input *model.CreateWizard) (string, error) {
	url := URL + "/wizard/create"
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
		ErrorResponse(ctx, resp)
		return "nil", err
	}
}

func (r *mutationResolver) UpdateWizard(ctx context.Context, input *model.UpdateWizard) (string, error) {
	url := URL + "/wizard/update/" + strconv.Itoa(input.ID)
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
		ErrorResponse(ctx, resp)
		return "nil", err
	}
}

func (r *mutationResolver) DeleteWizard(ctx context.Context, input *model.RequestWithID) (string, error) {
	url := URL + "/wizard/delete/" + strconv.Itoa(input.ID)
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
		ErrorResponse(ctx, resp)
		return "nil", err
	}
}

func (r *mutationResolver) CreateResult(ctx context.Context, input *model.CreateResult) (string, error) {
	url := URL + "/result/create"
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
		ErrorResponse(ctx, resp)
		return "nil", err
	}
}

func (r *queryResolver) LogoutAuth(ctx context.Context) (string, error) {
	url := URL + "/auth/logout"
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
	url := URL + "/user/find"
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
		ErrorResponse(ctx, resp)
		return nil, err
	}
}

func (r *queryResolver) GetUser(ctx context.Context, input model.RequestWithID) (*model.User, error) {
	url := URL + "/user/find/" + strconv.Itoa(input.ID)
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
		ErrorResponse(ctx, resp)
		return nil, err
	}
}

func (r *queryResolver) GetWizards(ctx context.Context, input model.Token) ([]*model.Wizard, error) {
	url := URL + "/wizard/find"
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
		var wizards []*model.Wizard
		json.Unmarshal(bodyResponse, &wizards)
		return wizards, nil
	} else {
		ErrorResponse(ctx, resp)
		return nil, err
	}
}

func (r *queryResolver) GetWizard(ctx context.Context, input model.RequestWithID) (*model.Wizard, error) {
	url := URL + "/wizard/find/" + strconv.Itoa(input.ID)
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
		var wizard model.Wizard
		json.Unmarshal(bodyResponse, &wizard)
		return &wizard, nil
	} else {
		ErrorResponse(ctx, resp)
		return nil, err
	}
}

func (r *queryResolver) GetResultsForUser(ctx context.Context, input model.RequestWithID) ([]*model.Result, error) {
	url := URL + "/result/findforuser/" + strconv.Itoa(input.ID)
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
		var results []*model.Result
		json.Unmarshal(bodyResponse, &results)
		return results, nil
	} else {
		ErrorResponse(ctx, resp)
		return nil, err
	}
}

func (r *queryResolver) GetResultsForModer(ctx context.Context, input model.RequestWithID) ([]*model.Result, error) {
	url := URL + "/result/findformoder/" + strconv.Itoa(input.ID)
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
		var results []*model.Result
		json.Unmarshal(bodyResponse, &results)
		return results, nil
	} else {
		ErrorResponse(ctx, resp)
		return nil, err
	}
}

func (r *queryResolver) GetResults(ctx context.Context, input model.Token) ([]*model.Result, error) {
	url := URL + "/result/findall"
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
		var results []*model.Result
		json.Unmarshal(bodyResponse, &results)
		return results, nil
	} else {
		ErrorResponse(ctx, resp)
		return nil, err
	}
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.

func ErrorResponse(ctx context.Context, resp *http.Response) {
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
}
