package request

type SignUp struct {
	Name           string `json:"name"`
	Lastname       string `json:"last_name"`
	Mail           string `json:"mail"`
	Identification string `json:"identification"`
	Phone          string `json:"phone"`
}

type SignOut struct {
	Token string
}

type Refresh struct {
	Token string
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ValidateContact struct {
	Type      string `json:"type"`
	Contact   string `json:"contact"`
	Recovery  bool   `json:"recovery"`
	SessionId string `json:"session_id"`
	Code      string `json:"code"`
}

type StartContactValidation struct {
	Type      string `json:"type"`
	Contact   string `json:"contact"`
	Recovery  bool   `json:"recovery"`
	SessionId string `json:"session_id"`
	FullName  string `json:"full_name"`
}

type CheckUserExists struct {
	Mail           string `json:"mail"`
	Identification string `json:"identification"`
	Phone          string `json:"phone"`
}

type GetAvailableRoutes struct {
	UserId string `json:"user_id"`
}

type ChangePassword struct {
	UserId      string `json:"user_id"`
	NewPassword string `json:"new_password"`
}

type GetUser struct {
	UserId string `json:"user_id"`
}
