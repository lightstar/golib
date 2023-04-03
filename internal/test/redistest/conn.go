// Package redistest is used as a helper in tests that need mocked redis connection.
package redistest

import "github.com/stretchr/testify/mock"

// MockConn structure provides mocked redis connection implementing redigo module's Conn interface.
type MockConn struct {
	mock.Mock
}

// NewMockConn function creates mocked redis connection and applies default expectations.
func NewMockConn() *MockConn {
	conn := &MockConn{}

	conn.On("Do", "", mock.Anything).Return(nil, nil).Maybe()
	conn.On("Err").Return(nil).Maybe()
	conn.On("Close").Return(nil).Once()

	return conn
}

// Close method mocks closing the connection.
func (conn *MockConn) Close() error {
	args := conn.Called()
	return args.Error(0)
}

// Err method mocks retrieving last connection error.
func (conn *MockConn) Err() error {
	args := conn.Called()
	return args.Error(0)
}

// Do method mocks sending command to the redis server and waiting for the reply.
func (conn *MockConn) Do(commandName string, params ...interface{}) (interface{}, error) {
	args := conn.Called(commandName, params)
	return args.Get(0), args.Error(1)
}

// Send method mocks sending command to the redis server without waiting for the reply.
func (conn *MockConn) Send(commandName string, params ...interface{}) error {
	args := conn.Called(commandName, params)
	return args.Error(0)
}

// Flush method mocks flushing pending commands to the redis server.
func (conn *MockConn) Flush() error {
	args := conn.Called()
	return args.Error(0)
}

// Receive method mocks receiving reply from the redis server.
func (conn *MockConn) Receive() (interface{}, error) {
	args := conn.Called()
	return args.Get(0), args.Error(1)
}
