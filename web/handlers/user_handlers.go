package handlers

type UserHandler struct {
	secret string
}

func NewUserHandler(secret string) *UserHandler {

	return &UserHandler{
		secret: secret,
	}
}
