package handlers

import (
	"errors"
	"fmt"

	"github.com/dkr290/go-projects/hotel-reservation/ctypes"
	"github.com/dkr290/go-projects/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
		id = c.Params("id")
	)

	user, err := h.userStore.GetUserById(c.Context(), id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return c.JSON(map[string]string{"error": "not found"})
		}
		return err

	}

	return c.JSON(user)

}
func (h *UserHandler) HandlePostUser(c *fiber.Ctx) error {

	var params ctypes.CreateUserParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}

	if err := params.Validate(); len(err) > 0 {
		return c.JSON(err)
	}

	user, err := ctypes.NewUserFromParams(params)
	if err != nil {
		return err
	}
	insertedUser, err := h.userStore.CreateUser(c.Context(), user)
	if err != nil {
		return err
	}
	return c.JSON(insertedUser)
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {

	users, err := h.userStore.GetUsers(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(users)
}

func (h *UserHandler) HandleDeleteUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	if err := h.userStore.DeleteUser(c.Context(), userId); err != nil {
		return err
	}

	return c.JSON(map[string]string{"deleted": userId})
}

func (h *UserHandler) HandleUpdateUser(c *fiber.Ctx) error {

	var (
		update bson.M
		userID = c.Params("id")
	)

	oid, err := db.ConvertObjectID(userID)
	if err != nil {
		return nil
	}

	if err := c.BodyParser(&update); err != nil {
		return err
	}
	fmt.Println(update)
	filter := bson.M{"_id": oid}
	if err := h.userStore.UpdateUser(c.Context(), filter, update); err != nil {
		return err
	}
	return c.JSON(map[string]string{"updated": userID})
}
