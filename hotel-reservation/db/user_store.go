package db

import (
	"github.com/dkr290/go-projects/hotel-reservation/ctypes"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserStore interface {
	GetUserById(string) (*ctypes.User, error)
}

// this is the implementations, for different databases
type MongoUserStore struct {
	client *mongo.Client
}

// type PostgresuserStore struct{}
func NewMongoUserStore(m *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		client: m,
	}

}

func (s *MongoUserStore) GetUserById(id string) (*ctypes.User, error) {

	return nil, mongo.ErrNilCursor
}
