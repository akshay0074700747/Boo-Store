package entities

type User struct {
	Username string `json:"username,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	IsAdmin  bool   `json:"is_admin,omitempty"`
}
