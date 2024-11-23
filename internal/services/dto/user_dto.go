package dto

type User struct {
	ID       string `json:"_id,omitempty"`
	Key      string `json:"_key,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role,omitempty"`
}
