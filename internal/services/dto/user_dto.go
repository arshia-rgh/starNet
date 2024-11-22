package dto

type User struct {
	Key      string `json:"_key,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"`
}
