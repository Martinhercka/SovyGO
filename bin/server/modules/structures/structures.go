package structures

//LoginRequest req
type LoginRequest struct {
	SessionID string `json:"sessionid,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
}

//LoginResponse res
type LoginResponse struct {
	Message string `json:"message,omitempty"`
}

//RegisterRequest req
type RegisterRequest struct {
	Username string
	Email    string
	Password string
	Class    string
}
