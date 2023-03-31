package mongo_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/storage/mongo"
)

type sampleDataType struct {
	Name string   `bson:"name"`
	Age  int      `bson:"age"`
	Food []string `bson:"food"`
}

//nolint:gochecknoglobals // this variable is actually read-only, so it's ok to use it.
var sampleData = sampleDataType{
	Name: "George",
	Age:  25,
	Food: []string{"milk", "bread", "meat"},
}

//nolint:gochecknoglobals // this variable is actually read-only, so it's ok to use it.
var sampleData2 = sampleDataType{
	Name: "Michael",
	Age:  25,
	Food: []string{"cookies"},
}

func TestInsertOne(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	dataID, err := helper.session.InsertOne(sampleData)
	require.NoError(t, err)

	helper.requireFindOne(t, dataID, &sampleData)
}

func TestInsertOneError(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	_, err := helper.session.InsertOne(nil)

	require.Error(t, err)
	require.Equal(t, "mongo error (document is nil)", err.Error())
}

func TestUpdateOne(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	dataID, err := helper.session.InsertOne(sampleData)
	require.NoError(t, err)

	updatedData := sampleDataType{
		Name: "Michael",
		Age:  42,
		Food: []string{"cookies"},
	}

	updateResult, err := helper.session.UpdateOne(mongo.Data{{Key: "_id", Value: dataID}}, updatedData)

	require.NoError(t, err)
	require.Equal(t, int64(1), updateResult.MatchedCount)

	helper.requireFindOne(t, dataID, &updatedData)
}

func TestUpdateOneError(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	_, err := helper.session.UpdateOne(nil, nil)

	require.Error(t, err)
	require.Regexp(t, `^mongo error \(cannot transform type mongo.Data to a BSON Document`, err.Error())
}

func TestUpsertOneUpdated(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	dataID, err := helper.session.InsertOne(sampleData)
	require.NoError(t, err)

	upsertedData := sampleDataType{
		Name: "Michael",
		Age:  42,
		Food: []string{"cookies"},
	}

	upsertResult, err := helper.session.UpsertOne(mongo.Data{{Key: "name", Value: "George"}}, upsertedData)

	require.NoError(t, err)
	require.Equal(t, int64(1), upsertResult.MatchedCount)
	require.Equal(t, int64(0), upsertResult.UpsertedCount)

	helper.requireFindOne(t, dataID, &upsertedData)
}

func TestUpsertOneInserted(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	dataID, err := helper.session.InsertOne(sampleData)
	require.NoError(t, err)

	upsertedData := sampleDataType{
		Name: "Michael",
		Age:  42,
		Food: []string{"cookies"},
	}

	upsertResult, err := helper.session.UpsertOne(mongo.Data{{Key: "name", Value: "Fred"}}, upsertedData)

	require.NoError(t, err)
	require.Equal(t, int64(0), upsertResult.MatchedCount)
	require.Equal(t, int64(1), upsertResult.UpsertedCount)

	helper.requireFindOne(t, dataID, &sampleData)
	helper.requireFindOne(t, upsertResult.UpsertedID, &upsertedData)
}

func TestUpsertOneError(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	_, err := helper.session.UpsertOne(nil, nil)

	require.Error(t, err)
	require.Regexp(t, `^mongo error \(cannot transform type mongo.Data to a BSON Document`, err.Error())
}

func TestDeleteOne(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	dataID, err := helper.session.InsertOne(sampleData)
	require.NoError(t, err)

	count, err := helper.session.Count(mongo.Data{{Key: "_id", Value: dataID}})

	require.NoError(t, err)
	require.Equal(t, int64(1), count)

	count, err = helper.session.DeleteOne(mongo.Data{{Key: "_id", Value: dataID}})

	require.NoError(t, err)
	require.Equal(t, int64(1), count)

	count, err = helper.session.Count(mongo.Data{{Key: "_id", Value: dataID}})

	require.NoError(t, err)
	require.Zero(t, count)
}

func TestDeleteOneError(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	_, err := helper.session.DeleteOne(nil)

	require.Error(t, err)
	require.Regexp(t, `^mongo error \(cannot transform type mongo.Data to a BSON Document`, err.Error())
}

func TestCount(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	count, err := helper.session.Count(mongo.Data{})

	require.NoError(t, err)
	require.Zero(t, count)

	dataID, err := helper.session.InsertOne(sampleData)
	require.NoError(t, err)

	count, err = helper.session.Count(mongo.Data{{Key: "_id", Value: dataID}})

	require.NoError(t, err)
	require.Equal(t, int64(1), count)
}

func TestCountError(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	_, err := helper.session.Count(nil)

	require.Error(t, err)
	require.Regexp(t, `^mongo error \(cannot transform type mongo.Data to a BSON Document`, err.Error())
}

func TestFindOne(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	dataID, err := helper.session.InsertOne(sampleData)
	require.NoError(t, err)

	data := sampleDataType{}
	c, err := helper.session.FindOne(mongo.Data{{Key: "_id", Value: dataID}}, &data)

	require.NoError(t, err)
	require.Equal(t, int64(1), c)
	require.Equal(t, &sampleData, &data)
}

func TestFindOneNoDocuments(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	data := sampleDataType{}
	c, err := helper.session.FindOne(mongo.Data{}, &data)

	require.NoError(t, err)
	require.Zero(t, c)
}

func TestFindOneError(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	data := sampleDataType{}
	_, err := helper.session.FindOne(nil, &data)

	require.Error(t, err)
	require.Regexp(t, `^mongo error \(cannot transform type mongo.Data to a BSON Document`, err.Error())
}

func TestCountAll(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	count, err := helper.session.CountAll()

	require.NoError(t, err)
	require.Zero(t, count)

	_, err = helper.session.InsertOne(sampleData)
	require.NoError(t, err)

	count, err = helper.session.CountAll()

	require.NoError(t, err)
	require.Equal(t, int64(1), count)

	_, err = helper.session.InsertOne(sampleData2)
	require.NoError(t, err)

	count, err = helper.session.CountAll()

	require.NoError(t, err)
	require.Equal(t, int64(2), count)
}

func TestInsertMany(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	dataIDs, err := helper.session.InsertMany([]interface{}{sampleData, sampleData2})

	require.NoError(t, err)
	require.Len(t, dataIDs, 2)

	helper.requireFindOne(t, dataIDs[0], &sampleData)
	helper.requireFindOne(t, dataIDs[1], &sampleData2)
}

func TestInsertManyError(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	_, err := helper.session.InsertMany([]interface{}{})

	require.Error(t, err)
	require.Equal(t, "mongo error (must provide at least one element in input slice)", err.Error())
}

func TestFind(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	var dataSlice []sampleDataType
	count, err := helper.session.Find(mongo.Data{{Key: "age", Value: 25}}, &dataSlice)

	require.NoError(t, err)
	require.Zero(t, count)

	_, err = helper.session.InsertMany([]interface{}{sampleData, sampleData2})
	require.NoError(t, err)

	count, err = helper.session.Find(mongo.Data{{Key: "age", Value: 25}}, &dataSlice)

	require.NoError(t, err)
	require.Equal(t, int64(2), count)
	require.Len(t, dataSlice, 2)
	require.Equal(t, &sampleData, &dataSlice[0])
	require.Equal(t, &sampleData2, &dataSlice[1])
}

func TestFindError(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	var dataSlice []sampleDataType
	_, err := helper.session.Find(nil, &dataSlice)

	require.Error(t, err)
	require.Regexp(t, `^mongo error \(cannot transform type mongo.Data to a BSON Document`, err.Error())
}

func TestFindEntitiesError(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	_, err := helper.session.Find(mongo.Data{}, nil)

	require.Error(t, err)
	require.Equal(t, "mongo error (results argument must be a pointer to a slice, but was a invalid)",
		err.Error())
}

func TestUpdateMany(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	dataIDs, err := helper.session.InsertMany([]interface{}{sampleData, sampleData2})

	require.NoError(t, err)
	require.Len(t, dataIDs, 2)

	updatedData := sampleDataType{
		Name: "George",
		Age:  31,
		Food: []string{"milk", "bread"},
	}

	updateResult, err := helper.session.UpdateMany(mongo.Data{{Key: "age", Value: 25}}, updatedData)

	require.NoError(t, err)
	require.Equal(t, int64(2), updateResult.MatchedCount)

	helper.requireFindOne(t, dataIDs[0], &updatedData)
	helper.requireFindOne(t, dataIDs[1], &updatedData)
}

func TestUpdateManyError(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	_, err := helper.session.UpdateMany(nil, nil)

	require.Error(t, err)
	require.Regexp(t, `^mongo error \(cannot transform type mongo.Data to a BSON Document`, err.Error())
}

func TestUpsertManyUpdated(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	dataIDs, err := helper.session.InsertMany([]interface{}{sampleData, sampleData2})

	require.NoError(t, err)
	require.Len(t, dataIDs, 2)

	upsertedData := sampleDataType{
		Name: "George",
		Age:  31,
		Food: []string{"milk", "bread"},
	}

	updateResult, err := helper.session.UpsertMany(mongo.Data{{Key: "age", Value: 25}}, upsertedData)

	require.NoError(t, err)
	require.Equal(t, int64(2), updateResult.MatchedCount)
	require.Equal(t, int64(0), updateResult.UpsertedCount)

	helper.requireFindOne(t, dataIDs[0], &upsertedData)
	helper.requireFindOne(t, dataIDs[1], &upsertedData)
}

func TestUpsertManyInserted(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	dataIDs, err := helper.session.InsertMany([]interface{}{sampleData, sampleData2})

	require.NoError(t, err)
	require.Len(t, dataIDs, 2)

	upsertedData := sampleDataType{
		Name: "George",
		Age:  31,
		Food: []string{"milk", "bread"},
	}

	updateResult, err := helper.session.UpsertMany(mongo.Data{{Key: "age", Value: 40}}, upsertedData)

	require.NoError(t, err)
	require.Equal(t, int64(0), updateResult.MatchedCount)
	require.Equal(t, int64(1), updateResult.UpsertedCount)

	helper.requireFindOne(t, dataIDs[0], &sampleData)
	helper.requireFindOne(t, dataIDs[1], &sampleData2)
	helper.requireFindOne(t, updateResult.UpsertedID, &upsertedData)
}

func TestUpsertManyError(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	_, err := helper.session.UpsertMany(nil, nil)

	require.Error(t, err)
	require.Regexp(t, `^mongo error \(cannot transform type mongo.Data to a BSON Document`, err.Error())
}

func TestDeleteMany(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	dataIDs, err := helper.session.InsertMany([]interface{}{sampleData, sampleData2})

	require.NoError(t, err)
	require.Len(t, dataIDs, 2)

	count, err := helper.session.DeleteMany(mongo.Data{{Key: "age", Value: 25}})

	require.NoError(t, err)
	require.Equal(t, int64(2), count)

	count, err = helper.session.Count(mongo.Data{{Key: "_id", Value: dataIDs[0]}})

	require.NoError(t, err)
	require.Zero(t, count)

	count, err = helper.session.Count(mongo.Data{{Key: "_id", Value: dataIDs[1]}})

	require.NoError(t, err)
	require.Zero(t, count)
}

func TestDeleteManyError(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	_, err := helper.session.DeleteMany(nil)

	require.Error(t, err)
	require.Regexp(t, `^mongo error \(cannot transform type mongo.Data to a BSON Document`, err.Error())
}
