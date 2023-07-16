package db

import (
	"context"
	"fmt"

	"github.com/dkr290/go-projects/hotel-reservation/ctypes"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	userColl = "users"
)

type Dropper interface {
	Drop(context.Context) error
}

type UserStore interface {
	GetUserById(context.Context, string) (*ctypes.User, error)
	GetUsers(context.Context) ([]*ctypes.User, error)
	CreateUser(context.Context, *ctypes.User) (*ctypes.User, error)
	DeleteUser(context.Context, string) error
	UpdateUser(ctx context.Context, filter bson.M, params ctypes.UpdateUserParams) error
	Dropper
}

// this is the implementations, for different databases
type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

// type PostgresuserStore struct{}
func NewMongoUserStore(m *mongo.Client, dbname string) *MongoUserStore {
	coll := m.Database(dbname).Collection(userColl)
	return &MongoUserStore{
		client: m,
		coll:   coll,
	}

}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*ctypes.User, error) {
	var users []*ctypes.User

	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}

	return users, nil

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

func (s *MongoUserStore) CreateUser(ctx context.Context, user *ctypes.User) (*ctypes.User, error) {
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)

	return user, nil
}

func (s *MongoUserStore) DeleteUser(ctx context.Context, id string) error {

	oid, err := ConvertObjectID(id)
	if err != nil {
		return err
	}
	//TODO to handle if we did not delete any user , something to log maybe somehow
	_, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoUserStore) UpdateUser(ctx context.Context, filter bson.M, params ctypes.UpdateUserParams) error {

	update := bson.D{
		{
			"$set", params.ToBSN(),
		},
	}

	_, err := s.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

func (s *MongoUserStore) Drop(ctx context.Context) error {
	fmt.Println("--- dropping user collection")
	return s.coll.Drop(ctx)
}
