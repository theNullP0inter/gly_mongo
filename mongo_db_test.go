package mongo_db

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

type MockModel struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id" copier:"must"`
	Name string             `bson:"name" json:"name" copier:"must"`
}

func NewMongoTest(t *testing.T) *mtest.T {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	return mt
}
