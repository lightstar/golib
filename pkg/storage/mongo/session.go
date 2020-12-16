package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Session structure provides access to operations with mongodb database and collection.
type Session struct {
	context    context.Context
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

// Context method retrieves session context.
func (session *Session) Context() context.Context {
	return session.context
}

// Database method retrieves session database name.
func (session *Session) Database() string {
	return session.database.Name()
}

// Collection method retrieves session collection name.
func (session *Session) Collection() string {
	return session.collection.Name()
}

// WithContext method changes session context.
func (session *Session) WithContext(ctx context.Context) *Session {
	session.context = ctx
	return session
}

// WithDatabase method changes session database and collection.
func (session *Session) WithDatabase(databaseName string, collectionName string) *Session {
	session.database = session.client.Database(databaseName)
	session.collection = session.database.Collection(collectionName)

	return session
}

// WithCollection method changes session collection.
func (session *Session) WithCollection(name string) *Session {
	session.collection = session.database.Collection(name)
	return session
}
