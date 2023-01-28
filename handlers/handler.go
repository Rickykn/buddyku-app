package handlers

import "github.com/Rickykn/buddyku-app.git/services"

type Handler struct {
	userService services.UserService
}

type HandlerConfig struct {
	UserService services.UserService
}

func New(c *HandlerConfig) *Handler {
	return &Handler{
		userService: c.UserService,
	}
}
