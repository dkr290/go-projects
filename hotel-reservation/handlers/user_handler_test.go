package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/dkr290/go-projects/hotel-reservation/ctypes"
	"github.com/dkr290/go-projects/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dburi  = "mongodb://localhost:27017"
	dbname = "test_reservations"
)

type testdb struct {
	db.UserStore
}

func (db *testdb) teardown(t *testing.T) {
	if err := db.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dburi))
	if err != nil {
		log.Fatal(err)
	}

	return &testdb{
		UserStore: db.NewMongoUserStore(client, dbname),
	}

}

func TestPostUser(t *testing.T) {
	db := setup(t)

	defer db.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler(db.UserStore)
	app.Post("/", userHandler.HandlePostUser)

	params := ctypes.CreateUserParams{
		Email:     "some@foo.com",
		FirstName: "James",
		LastName:  "Foo",
		Password:  "Pass12",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")

	r, _ := app.Test(req)

	var user ctypes.User
	json.NewDecoder(r.Body).Decode(&user)

	if len(user.ID) == 0 {
		t.Error("expecting user id to be set")
	}

	if len(user.EncryptedPassword) > 0 {
		t.Errorf("expected the encrypted password not to be included into the json response")
	}

	if user.FirstName != params.FirstName {
		t.Errorf("expected firstname %s but got %s", params.FirstName, user.FirstName)
	}

	if user.LastName != params.LastName {
		t.Errorf("expected lastname %s but got %s", params.LastName, user.LastName)
	}

	if user.Email != params.Email {
		t.Errorf("expected email %s but got %s", params.Email, user.Email)
	}

}
