package handlers

import (
	"context"

	"github.com/dkr290/go-projects/hotel-reservation/ctypes"
	"github.com/dkr290/go-projects/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(uSt db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: uSt,
	}
}

func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)

	user, err := h.userStore.GetUserById(ctx, id)
	if err != nil {
		return err
	}

	return c.JSON(user)

}
func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {

	u := ctypes.User{
		FirstName: "James",
		LastName:  "Ath the watercooler",
	}
	return c.JSON(u)
}
