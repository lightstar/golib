package i2s_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/config/i2s"
	"github.com/lightstar/golib/pkg/errors"
)

type Test struct {
	name     string
	in       interface{}
	out      interface{}
	expected interface{}
}

type TestError struct {
	name    string
	in      interface{}
	out     interface{}
	err     error
	wrapped bool
}

//nolint:funlen // tests slice is too long to pass that linter
func TestConvert(t *testing.T) {
	convertor := i2s.Instance()

	tests := []Test{
		{
			name: "String",
			in: map[string]interface{}{
				"key": "value",
			},
			out: &struct {
				Key string
			}{},
			expected: &struct {
				Key string
			}{
				Key: "value",
			},
		},
		{
			name: "Int",
			in: map[string]interface{}{
				"key": 5,
			},
			out: &struct {
				Key int
			}{},
			expected: &struct {
				Key int
			}{
				Key: 5,
			},
		},
		{
			name: "Float",
			in: map[string]interface{}{
				"key": 2.3,
			},
			out: &struct {
				Key float64
			}{},
			expected: &struct {
				Key float64
			}{
				Key: 2.3,
			},
		},
		{
			name: "Float2Int",
			in: map[string]interface{}{
				"key": 3.6,
			},
			out: &struct {
				Key int
			}{},
			expected: &struct {
				Key int
			}{
				Key: 3,
			},
		},
		{
			name: "Int2Float",
			in: map[string]interface{}{
				"key": 4,
			},
			out: &struct {
				Key float64
			}{},
			expected: &struct {
				Key float64
			}{
				Key: 4.,
			},
		},
		{
			name: "Bool",
			in: map[string]interface{}{
				"key": true,
			},
			out: &struct {
				Key bool
			}{},
			expected: &struct {
				Key bool
			}{
				Key: true,
			},
		},
		{
			name: "Map",
			in: map[string]interface{}{
				"key": map[string]interface{}{
					"foo": "bar",
				},
			},
			out: &struct {
				Key struct {
					Foo string
				}
			}{},
			expected: &struct {
				Key struct {
					Foo string
				}
			}{
				Key: struct {
					Foo string
				}{
					Foo: "bar",
				},
			},
		},
		{
			name: "Map2Map",
			in: map[string]interface{}{
				"key": map[string]interface{}{
					"foo": "bar",
				},
			},
			out: &struct {
				Key map[string]interface{}
			}{},
			expected: &struct {
				Key map[string]interface{}
			}{
				Key: map[string]interface{}{"foo": "bar"},
			},
		},
		{
			name: "SliceOfStrings",
			in: map[string]interface{}{
				"key": []interface{}{"value1", "value2"},
			},
			out: &struct {
				Key []string
			}{},
			expected: &struct {
				Key []string
			}{
				Key: []string{"value1", "value2"},
			},
		},
		{
			name: "SliceOfMaps",
			in: map[string]interface{}{
				"key": []map[string]interface{}{
					{"foo": "bar1"},
					{"foo": "bar2"},
				},
			},
			out: &struct {
				Key []struct {
					Foo string
				}
			}{},
			expected: &struct {
				Key []struct {
					Foo string
				}
			}{
				Key: []struct {
					Foo string
				}{
					{Foo: "bar1"},
					{Foo: "bar2"},
				},
			},
		},
		{
			name: "Compound",
			in: map[string]interface{}{
				"name":    "Peter",
				"age":     30,
				"weight":  112.5,
				"married": true,
				"profile": map[string]interface{}{
					"email":    "peter@mail.com",
					"score":    12,
					"children": []interface{}{"Mary", "Ann", "Michael"},
				},
			},
			out: &struct {
				Name    string
				Age     int
				Weight  float32
				Married bool
				Profile struct {
					Email    string
					Score    int32
					Children []string
				}
			}{},
			expected: &struct {
				Name    string
				Age     int
				Weight  float32
				Married bool
				Profile struct {
					Email    string
					Score    int32
					Children []string
				}
			}{
				Name:    "Peter",
				Age:     30,
				Weight:  112.5,
				Married: true,
				Profile: struct {
					Email    string
					Score    int32
					Children []string
				}{
					Email:    "peter@mail.com",
					Score:    12,
					Children: []string{"Mary", "Ann", "Michael"},
				},
			},
		},
	}

	for _, test := range tests {
		func(test Test) {
			t.Run(test.name, func(t *testing.T) {
				err := convertor.Convert(test.in, test.out)
				require.NoError(t, err)
				require.Equal(t, test.out, test.expected)
			})
		}(test)
	}
}

func TestErrors(t *testing.T) {
	convertor := i2s.Instance()

	tests := []TestError{
		{
			name: "OutputNotPointer",
			in:   map[string]interface{}{},
			out:  struct{}{},
			err:  i2s.ErrOutputNotPointer,
		},
		{
			name: "String2Int",
			in:   map[string]interface{}{"key": "value"},
			out:  &struct{ Key int }{},
			err:  i2s.ErrMismatchedTypes,
		},
		{
			name: "Int2String",
			in:   map[string]interface{}{"key": 5},
			out:  &struct{ Key string }{},
			err:  i2s.ErrMismatchedTypes,
		},
		{
			name: "Float2String",
			in:   map[string]interface{}{"key": 5.},
			out:  &struct{ Key string }{},
			err:  i2s.ErrMismatchedTypes,
		},
		{
			name: "Bool2String",
			in:   map[string]interface{}{"key": true},
			out:  &struct{ Key string }{},
			err:  i2s.ErrMismatchedTypes,
		},
		{
			name: "Map2String",
			in:   map[string]interface{}{"key": map[string]interface{}{}},
			out:  &struct{ Key string }{},
			err:  i2s.ErrMismatchedTypes,
		},
		{
			name: "Map2Map",
			in:   map[string]interface{}{"key": "value"},
			out:  &map[string]string{},
			err:  i2s.ErrMismatchedTypes,
		},
		{
			name: "Slice2String",
			in:   map[string]interface{}{"key": []interface{}{}},
			out:  &struct{ Key string }{},
			err:  i2s.ErrMismatchedTypes,
		},
		{
			name: "IntMapKey",
			in:   map[string]interface{}{"key": map[int]interface{}{0: "data"}},
			out:  &struct{ Key struct{} }{},
			err:  i2s.ErrMapKeyNotString,
		},
		{
			name:    "MissedField",
			in:      map[string]interface{}{"key": "value"},
			out:     &struct{}{},
			err:     i2s.ErrUnknownField,
			wrapped: true,
		},
		{
			name:    "Func",
			in:      map[string]interface{}{"key": func() {}},
			out:     &struct{ Key func() }{},
			err:     i2s.ErrUnsupportedType,
			wrapped: true,
		},
		{
			name:    "SliceOfFuncs",
			in:      map[string]interface{}{"key": []func(){func() {}}},
			out:     &struct{ Key []func() }{},
			err:     i2s.ErrUnsupportedType,
			wrapped: true,
		},
	}

	for _, test := range tests {
		func(test TestError) {
			t.Run(test.name, func(t *testing.T) {
				err := convertor.Convert(test.in, test.out)

				if test.wrapped {
					err = errors.Unwrap(err)
				}

				require.Same(t, test.err, err)
			})
		}(test)
	}
}
