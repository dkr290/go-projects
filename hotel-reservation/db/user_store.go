package db

import (
	"context"

	"github.com/dkr290/go-projects/hotel-reservation/ctypes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DBNAME   = "reservations"
	userColl = "users"
)

type UserStore interface {
	GetUserById(context.Context, string) (*ctypes.User, error)
}

// this is the implementations, for different databases
type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

// type PostgresuserStore struct{}
func NewMongoUserStore(m *mongo.Client) *MongoUserStore {
	coll := m.Database(DBNAME).Collection(userColl)
	return &MongoUserStore{
		client: m,
		coll:   coll,
	}

}

func (s *MongoUserStore) GetUserById(ctx context.Context, id string) (*ctypes.User, error) {

	//validate the correctness of the ID with helper function

	oid, err := ConvertObjectID(id)
	if err != nil {
		return nil, err
	}

	var user ctypes.User

	if err := s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user); err != nil {
		return nil, err
	}

	return &user, nil

}
