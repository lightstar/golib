// Package i2s provides convertor from raw data representation as map[string]interface{} into external structure
// using reflection. It can also convert []map[string]interface{} or []interface{} into slice of structures.
//
// Designed to be used inside config package with data unmarshalled by json, yaml and toml encoders.
//
// It provides singleton Convertor instance that must be obtained with Instance function.
package i2s

import (
	"reflect"
	"sync"

	"github.com/lightstar/golib/pkg/errors"
)

// These variables are used to support singleton pattern.
//
//nolint:gochecknoglobals // used for singleton pattern
var (
	instance *Convertor
	once     sync.Once
)

// Convertor structure that provides converting functionality. Don't create it manually, use Instance function instead.
type Convertor struct {
	processFuncMap map[reflect.Kind]func(reflect.Value, reflect.Value) error
}

// Instance function retrieves Convertor singleton instance.
func Instance() *Convertor {
	once.Do(func() {
		instance = &Convertor{}
		instance.processFuncMap = map[reflect.Kind]func(reflect.Value, reflect.Value) error{
			reflect.String:  instance.processString,
			reflect.Int:     instance.processInt,
			reflect.Int8:    instance.processInt,
			reflect.Int16:   instance.processInt,
			reflect.Int32:   instance.processInt,
			reflect.Int64:   instance.processInt,
			reflect.Float32: instance.processFloat,
			reflect.Float64: instance.processFloat,
			reflect.Bool:    instance.processBool,
			reflect.Map:     instance.processMap,
			reflect.Slice:   instance.processSlice,
		}
	})

	return instance
}

// Convert method converts raw data in 'data' parameter into structure (or slice of structures) that 'out' parameter
// points to.
// It will return an error if the structure doesn't have some field or it is not of an appropriate type.
func (c *Convertor) Convert(data interface{}, out interface{}) error {
	outValue := reflect.ValueOf(out)
	if outValue.Kind() != reflect.Ptr {
		return ErrOutputNotPointer
	}

	return c.process(reflect.ValueOf(data), outValue.Elem())
}

func (c *Convertor) process(dataValue reflect.Value, outValue reflect.Value) error {
	if dataValue.Kind() == reflect.Interface {
		dataValue = dataValue.Elem()
	}

	if processFunc, ok := c.processFuncMap[dataValue.Kind()]; ok {
		if err := processFunc(dataValue, outValue); err != nil {
			return err
		}
	} else {
		return errors.NewFmt("unsupported type '%v'", dataValue.Kind()).WithCause(ErrUnsupportedType)
	}

	return nil
}

func (c *Convertor) processString(dataValue reflect.Value, outValue reflect.Value) error {
	if outValue.Kind() != reflect.String {
		return ErrMismatchedTypes
	}

	outValue.SetString(dataValue.String())

	return nil
}

func (c *Convertor) processInt(dataValue reflect.Value, outValue reflect.Value) error {
	switch outValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		outValue.SetInt(dataValue.Int())
	case reflect.Float32, reflect.Float64:
		outValue.SetFloat(float64(dataValue.Int()))
	default:
		return ErrMismatchedTypes
	}

	return nil
}

func (c *Convertor) processFloat(dataValue reflect.Value, outValue reflect.Value) error {
	switch outValue.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		outValue.SetInt(int64(dataValue.Float()))
	case reflect.Float32, reflect.Float64:
		outValue.SetFloat(dataValue.Float())
	default:
		return ErrMismatchedTypes
	}

	return nil
}

func (c *Convertor) processBool(dataValue reflect.Value, outValue reflect.Value) error {
	if outValue.Kind() != reflect.Bool {
		return ErrMismatchedTypes
	}

	outValue.SetBool(dataValue.Bool())

	return nil
}

func (c *Convertor) processMap(dataValue reflect.Value, outValue reflect.Value) error {
	if outValue.Kind() == reflect.Map && outValue.Type().Elem().Kind() == dataValue.Type().Elem().Kind() {
		outValue.Set(dataValue)
		return nil
	}

	if outValue.Kind() != reflect.Struct {
		return ErrMismatchedTypes
	}

	mapIter := dataValue.MapRange()
	for mapIter.Next() {
		mapKey, mapValue := mapIter.Key(), mapIter.Value()

		if mapKey.Kind() != reflect.String {
			return ErrMapKeyNotString
		}

		if mapValue.Kind() == reflect.Interface {
			mapValue = mapValue.Elem()
		}

		mapKeyBytes := []byte(mapKey.String())
		if len(mapKeyBytes) > 0 {
			mapKeyBytes[0] -= 32
		}

		fieldValue := outValue.FieldByName(string(mapKeyBytes))
		if !fieldValue.IsValid() {
			return errors.NewFmt("unknown field '%s'", mapKey.String()).WithCause(ErrUnknownField)
		}

		if err := c.process(mapValue, fieldValue); err != nil {
			return err
		}
	}

	return nil
}

func (c *Convertor) processSlice(dataValue reflect.Value, outValue reflect.Value) error {
	if outValue.Kind() != reflect.Slice {
		return ErrMismatchedTypes
	}

	sliceValue := reflect.MakeSlice(outValue.Type(), 0, dataValue.Len())

	for i := 0; i < dataValue.Len(); i++ {
		elemValue := reflect.New(outValue.Type().Elem())

		if err := c.process(dataValue.Index(i), elemValue.Elem()); err != nil {
			return err
		}

		sliceValue = reflect.Append(sliceValue, elemValue.Elem())
	}

	outValue.Set(sliceValue)

	return nil
}
