// Package middleware provides commonly-used middleware functions for the httpservice package. Feed them into
// UseMiddleware and UseMiddlewareBefore methods of http service object.
//
// Typical usage:
//
//	service.UseMiddleware(middleware.Recover)
//	service.UseMiddleware(middleware.Log)
package middleware
