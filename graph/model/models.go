package model

type User struct {
	ID       int     `json:"id"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Roles    []*Role `json:"roles"`
}

type Role struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Token struct {
	Token string `json:"token"`
}
type RequestWithID struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

type CreateUser struct {
	Username string   `json:"username"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
	Token    string   `json:"token"`
}
type UpdateUser struct {
	ID       int      `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Roles    []string `json:"roles"`
	Token    string   `json:"token"`
}

type InputUser struct {
	ID int `json:"id"`
}
type CreateWizard struct {
	Name    string       `json:"name"`
	Creator *InputUser   `json:"creator"`
	Pages   []*InputPage `json:"pages"`
	Token   string       `json:"token"`
}
type UpdateWizard struct {
	ID      int          `json:"id"`
	Name    string       `json:"name"`
	Creator *InputUser   `json:"creator"`
	Pages   []*InputPage `json:"pages"`
	Token   string       `json:"token"`
}
type InputPage struct {
	Name    string         `json:"name"`
	Type    string         `json:"type"`
	Content string         `json:"content"`
	Buttons []*InputButton `json:"buttons"`
}
type InputToPage struct {
	Name string `json:"name"`
}
type InputButton struct {
	Name   string       `json:"name"`
	ToPage *InputToPage `json:"toPage"`
}

type CreateResult struct {
	Token    string       `json:"token"`
	UserID   int          `json:"user_id"`
	WizardID int          `json:"wizard_id"`
	Notes    []*InputNote `json:"notes"`
}
type InputNote struct {
	Page   *InputPageResult   `json:"page"`
	Button *InputButtonResult `json:"button"`
}
type InputPageResult struct {
	ID int `json:"id"`
}
type InputButtonResult struct {
	ID int `json:"id"`
}

type Wizard struct {
	ID      int     `json:"id"`
	Name    string  `json:"name"`
	Creator *User   `json:"creator"`
	Pages   []*Page `json:"pages"`
}

type Page struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Type    string    `json:"type"`
	Content string    `json:"content"`
	Buttons []*Button `json:"buttons"`
}

type Button struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	ToPage *Page  `json:"toPage"`
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
	Error   string `json:"error"`
}

type Result struct {
	ID     int     `json:"id"`
	Wizard *Wizard `json:"wizard"`
	User   *User   `json:"user"`
	Date   string  `json:"date"`
	Note   []*Note `json:"note"`
}

type Note struct {
	Page   *Page   `json:"page"`
	Button *Button `json:"button"`
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
