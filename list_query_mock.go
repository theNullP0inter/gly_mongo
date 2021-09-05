package mongo_db

import (
	"github.com/stretchr/testify/mock"
	"github.com/theNullP0inter/googly/resource"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MockMongoListQueryBuilder struct {
	mock.Mock
}

func (b *MockMongoListQueryBuilder) ListQuery(resource.ListQuery) (bson.M, *options.FindOptions) {
	b.Called()
	return bson.M{}, options.Find()
}
