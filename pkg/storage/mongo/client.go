// Package mongo provides convenient API to mongodb. It uses mongo official driver undercover.
//
// Typical simplified usage:
//
//	client := mongo.MustNewClient(mongo.WithAddress("127.0.0.1:27017"))
//	defer client.Close()
//
//	if err := client.Ping(); err != nil {
//	    ...
//	}
//
//	session := client.Session(ctx, dbName, collectionName)
//	count, err := session.Find(mongo.Data{}, &entities)
//	if err != nil {
//	    ...
//	}
package mongo

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client structure provides mongo client functionality. Don't create manually, use the functions down below
// instead.
type Client struct {
	address        string
	connectTimeout time.Duration
	socketTimeout  time.Duration
	client         *mongo.Client
}

// NewClient function creates new mongo client with provided options.
func NewClient(opts ...Option) (*Client, error) {
	config, err := buildConfig(opts)
	if err != nil {
		return nil, err
	}

	connectOpts := options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s", config.address)).
		SetConnectTimeout(config.connectTimeout).
		SetServerSelectionTimeout(config.connectTimeout).
		SetSocketTimeout(config.socketTimeout)

	client, err := mongo.Connect(context.Background(), connectOpts)
	if err != nil {
		return nil, mongoError(err)
	}

	return &Client{
		address:        config.address,
		connectTimeout: config.connectTimeout,
		socketTimeout:  config.socketTimeout,
		client:         client,
	}, nil
}

// MustNewClient function creates new mongo client with provided options and panics on any error.
func MustNewClient(opts ...Option) *Client {
	client, err := NewClient(opts...)
	if err != nil {
		panic(err)
	}

	return client
}

// Address method retrieves mongo server address.
func (client *Client) Address() string {
	return client.address
}

// ConnectTimeout method retrieves mongo client connect timeout.
func (client *Client) ConnectTimeout() time.Duration {
	return client.connectTimeout
}

// SocketTimeout method retrieves mongo client socket timeout.
func (client *Client) SocketTimeout() time.Duration {
	return client.socketTimeout
}

// Ping method pings mongo server to check that connection is successfully established.
func (client *Client) Ping() error {
	err := client.client.Ping(context.Background(), nil)
	if err != nil {
		return mongoError(err)
	}

	return nil
}

// Session method creates new session to perform useful commands on mongodb database and collection.
func (client *Client) Session(ctx context.Context, databaseName string, collectionName string) *Session {
	database := client.client.Database(databaseName)

	return &Session{
		context:    ctx,
		client:     client.client,
		database:   database,
		collection: database.Collection(collectionName),
	}
}

// Close method releases all resources used by client. Don't use client object after that.
func (client *Client) Close() error {
	err := client.client.Disconnect(context.Background())
	if err != nil {
		return mongoError(err)
	}

	return nil
}
