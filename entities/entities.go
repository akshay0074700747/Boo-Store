package entities

type User struct {
	Username string `json:"username,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	IsAdmin  bool   `json:"is_admin,omitempty"`
}

type Book struct {
	BookName        string `json:"bookName,omitempty"`
	Author          string `json:"author,omitempty"`
	PublicationYear int   `json:"publicationYear,omitempty"`
}
