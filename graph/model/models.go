package model

type Todo struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
	User *User  `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID          int      `json:"id"`
	Username    string   `json:"username"`
	Roles       []string `json:"roles"`
	TokenType   string   `json:"tokenType"`
	AccessToken string   `json:"accessToken"`
}

type BadResponse struct {
	Path    string `json:"path"`
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error string `json:"error"`
}

/*func MarshalResponse(response http.Response) graphql.Marshaler {
	return graphql.WriterFunc(func(w io.Writer) {
		responseData,err := ioutil.ReadAll(responce)
		w.Write([]byte(responce))

	})
}

func UnmarshalResponse(v interface{}) (http.Response, error) {
	fmt.Println(v)
	return http.Response{}, nil
}*/