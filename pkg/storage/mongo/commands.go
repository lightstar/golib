package mongo

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Data type is used to provide filters in various mongodb commands.
type Data bson.D

// InsertOne method inserts entity into collection. On success it returns id of inserted entity.
func (session *Session) InsertOne(entity interface{}) (interface{}, error) {
	result, err := session.collection.InsertOne(session.context, entity)
	if err != nil {
		return nil, mongoError(err)
	}

	return result.InsertedID, nil
}

// InsertMany method inserts list of entities into collection. On success it returns ids of inserted entities.
func (session *Session) InsertMany(entities []interface{}) ([]interface{}, error) {
	result, err := session.collection.InsertMany(session.context, entities)
	if err != nil {
		return nil, mongoError(err)
	}

	return result.InsertedIDs, nil
}

// UpdateOne method updates entity in collection using filter. On success it returns structure with update result.
func (session *Session) UpdateOne(filter Data, entity interface{}) (*mongo.UpdateResult, error) {
	result, err := session.collection.UpdateOne(session.context, filter, bson.D{{Key: "$set", Value: entity}})
	if err != nil {
		return nil, mongoError(err)
	}

	return result, nil
}

// UpdateMany method updates list of entities in collection using filter. On success it returns structure with update
// result.
func (session *Session) UpdateMany(filter Data, entity interface{}) (*mongo.UpdateResult, error) {
	result, err := session.collection.UpdateMany(session.context, filter, bson.D{{Key: "$set", Value: entity}})
	if err != nil {
		return nil, mongoError(err)
	}

	return result, nil
}

// UpsertOne method updates entity in collection using filter. If nothing to update, new entity will be inserted.
// On success it returns structure with update result.
func (session *Session) UpsertOne(filter Data, entity interface{}) (*mongo.UpdateResult, error) {
	result, err := session.collection.UpdateOne(session.context, filter, bson.D{{Key: "$set", Value: entity}},
		options.Update().SetUpsert(true))
	if err != nil {
		return nil, mongoError(err)
	}

	return result, nil
}

// UpsertMany method updates list of entities in collection using filter. If nothing to update, new entity will be
// inserted. On success it returns structure with update result.
func (session *Session) UpsertMany(filter Data, entity interface{}) (*mongo.UpdateResult, error) {
	result, err := session.collection.UpdateMany(session.context, filter, bson.D{{Key: "$set", Value: entity}},
		options.Update().SetUpsert(true))
	if err != nil {
		return nil, mongoError(err)
	}

	return result, nil
}

// DeleteOne method deletes entity from collection using filter. On success it returns number of deleted entities
// (0 or 1).
func (session *Session) DeleteOne(filter Data) (int64, error) {
	result, err := session.collection.DeleteOne(session.context, filter)
	if err != nil {
		return 0, mongoError(err)
	}

	return result.DeletedCount, nil
}

// DeleteMany method deletes all entities from collection using filter. On success it returns number of deleted
// entities.
func (session *Session) DeleteMany(filter Data) (int64, error) {
	result, err := session.collection.DeleteMany(session.context, filter)
	if err != nil {
		return 0, mongoError(err)
	}

	return result.DeletedCount, nil
}

// Count method retrieves number of entities in collection satisfying filter.
func (session *Session) Count(filter Data) (int64, error) {
	count, err := session.collection.CountDocuments(session.context, filter)
	if err != nil {
		return 0, mongoError(err)
	}

	return count, nil
}

// CountAll method retrieves total number of entities in collection.
func (session *Session) CountAll() (int64, error) {
	count, err := session.collection.EstimatedDocumentCount(session.context)
	if err != nil {
		return 0, mongoError(err)
	}

	return count, nil
}

// FindOne method retrieves one entity from collection using filter into entity parameter that must be a pointer to
// appropriate structure. On success it returns number of retrieved entities.
func (session *Session) FindOne(filter Data, entity interface{}, opts ...*options.FindOneOptions) (int64, error) {
	err := session.collection.FindOne(session.context, filter, opts...).Decode(entity)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return 0, nil
		}

		return 0, mongoError(err)
	}

	return 1, nil
}

// Find method retrieves list of entities from collection using filter into entities parameter that must be a pointer to
// a slice of appropriate structures. On success it returns number of retrieved entities.
func (session *Session) Find(filter Data, entities interface{}, opts ...*options.FindOptions) (int64, error) {
	cursor, err := session.collection.Find(session.context, filter, opts...)
	if err != nil {
		return 0, mongoError(err)
	}

	err = cursor.All(session.context, entities)
	if err != nil {
		return 0, mongoError(err)
	}

	return int64(reflect.ValueOf(entities).Elem().Len()), nil
}

// Drop method drops collection completely.
func (session *Session) Drop() error {
	err := session.collection.Drop(session.context)
	if err != nil {
		return mongoError(err)
	}

	return nil
}
