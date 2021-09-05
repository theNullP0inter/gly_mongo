package mongo_db

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/theNullP0inter/googly/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestBaseMongoResourceManagerCreate(t *testing.T) {
	mongo_test := NewMongoTest(t)
	mongo_test.Run("BaseMongoResourceManager: Create", func(mt *mtest.T) {
		var m MockModel
		rm := NewBaseMongoResourceManager(
			mt.DB, "mock",
			new(logger.MockGooglyLogger),
			m,
			new(MockMongoListQueryBuilder),
		)
		create_request := &MockModel{Name: "mock"}
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		item, err := rm.Create(create_request)
		assert.Nil(t, err)
		assert.NotNil(t, item)

	})

}

func TestBaseMongoResourceManagerGet(t *testing.T) {
	mongo_test := NewMongoTest(t)
	mongo_test.Run("BaseMongoResourceManager: Get", func(mt *mtest.T) {
		var m MockModel
		rm := NewBaseMongoResourceManager(
			mt.DB, "mock",
			new(logger.MockGooglyLogger),
			m,
			new(MockMongoListQueryBuilder),
		)
		id := primitive.NewObjectID().Hex()
		mt.AddMockResponses(
			mtest.CreateCursorResponse(
				1, "db.mock", mtest.FirstBatch, primitive.D{{"foo", "bar"}}))
		item, err := rm.Get(id)
		assert.Nil(t, err)
		assert.NotNil(t, item)

	})

}

func TestBaseMongoResourceManagerList(t *testing.T) {
	mongo_test := NewMongoTest(t)
	mongo_test.Run("BaseMongoResourceManager: Get", func(mt *mtest.T) {
		var m MockModel
		qb := new(MockMongoListQueryBuilder)
		rm := NewBaseMongoResourceManager(
			mt.DB, "mock",
			new(logger.MockGooglyLogger),
			m,
			qb,
		)
		params := primitive.NewObjectID().Hex()
		qb.On("ListQuery", mock.Anything).Return(mock.Anything, mock.Anything)
		mt.AddMockResponses(
			mtest.CreateCursorResponse(1, "db.mock", mtest.FirstBatch, primitive.D{{"foo", "bar"}}),
			mtest.CreateCursorResponse(1, "db.mock", mtest.NextBatch, primitive.D{{"foo", "bar"}}),
			mtest.CreateCursorResponse(1, "db.mock", mtest.NextBatch),
		)
		items, err := rm.List(params)
		assert.Nil(t, err)
		assert.NotNil(t, items)

	})

}

func TestBaseMongoResourceManagerUpdate(t *testing.T) {
	mongo_test := NewMongoTest(t)
	mongo_test.Run("BaseMongoResourceManager: Update", func(mt *mtest.T) {
		var m MockModel
		rm := NewBaseMongoResourceManager(
			mt.DB, "mock",
			new(logger.MockGooglyLogger),
			m,
			new(MockMongoListQueryBuilder),
		)
		id := primitive.NewObjectID().Hex()
		update_request := &MockModel{Name: "mock"}
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		err := rm.Update(id, update_request)
		assert.Nil(t, err)
	})

}

func TestBaseMongoResourceManagerDelete(t *testing.T) {
	mongo_test := NewMongoTest(t)
	mongo_test.Run("BaseMongoResourceManager: Update", func(mt *mtest.T) {
		var m MockModel
		rm := NewBaseMongoResourceManager(
			mt.DB, "mock",
			new(logger.MockGooglyLogger),
			m,
			new(MockMongoListQueryBuilder),
		)
		id := primitive.NewObjectID().Hex()
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		err := rm.Delete(id)
		assert.Nil(t, err)

	})

}
