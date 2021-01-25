package request

type AdminLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
