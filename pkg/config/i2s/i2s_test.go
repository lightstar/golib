package i2s_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/config/i2s"
	"github.com/lightstar/golib/pkg/errors"
)

// nolint: funlen // cases array is too long to pass that linter
func TestConvert(t *testing.T) {
	convertor := i2s.Instance()

	cases := []struct {
		in       interface{}
		out      interface{}
		expected interface{}
	}{
		{
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

	for i, c := range cases {
		err := convertor.Convert(c.in, c.out)
		require.NoErrorf(t, err, "case %d", i)
		require.Equalf(t, c.out, c.expected, "case %d", i)
	}
}

func TestErrors(t *testing.T) {
	convertor := i2s.Instance()

	cases := []struct {
		in      interface{}
		out     interface{}
		err     error
		wrapped bool
	}{
		{
			in:  map[string]interface{}{},
			out: struct{}{},
			err: i2s.ErrOutputNotPointer,
		},
		{
			in:  map[string]interface{}{"key": "value"},
			out: &struct{ Key int }{},
			err: i2s.ErrMismatchedTypes,
		},
		{
			in:  map[string]interface{}{"key": 5},
			out: &struct{ Key string }{},
			err: i2s.ErrMismatchedTypes,
		},
		{
			in:  map[string]interface{}{"key": 5.},
			out: &struct{ Key string }{},
			err: i2s.ErrMismatchedTypes,
		},
		{
			in:  map[string]interface{}{"key": true},
			out: &struct{ Key string }{},
			err: i2s.ErrMismatchedTypes,
		},
		{
			in:  map[string]interface{}{"key": map[string]interface{}{}},
			out: &struct{ Key string }{},
			err: i2s.ErrMismatchedTypes,
		},
		{
			in:  map[string]interface{}{"key": []interface{}{}},
			out: &struct{ Key string }{},
			err: i2s.ErrMismatchedTypes,
		},
		{
			in:  map[string]interface{}{"key": map[int]interface{}{0: "data"}},
			out: &struct{ Key struct{} }{},
			err: i2s.ErrMapKeyNotString,
		},
		{
			in:      map[string]interface{}{"key": "value"},
			out:     &struct{}{},
			err:     i2s.ErrUnknownField,
			wrapped: true,
		},
		{
			in:      map[string]interface{}{"key": func() {}},
			out:     &struct{ Key func() }{},
			err:     i2s.ErrUnsupportedType,
			wrapped: true,
		},
		{
			in:      map[string]interface{}{"key": []func(){func() {}}},
			out:     &struct{ Key []func() }{},
			err:     i2s.ErrUnsupportedType,
			wrapped: true,
		},
	}

	for i, c := range cases {
		err := convertor.Convert(c.in, c.out)

		if c.wrapped {
			err = errors.Unwrap(err)
		}

		require.Samef(t, c.err, err, "case %d", i)
	}
}

//
// func TestNumberConvert(t *testing.T) {
// 	cfg := config.NewFromRaw(map[string]interface{}{
// 		"key": 5,
// 	})
//
// 	data := &struct {
// 		Key float64
// 	}{}
// 	err := cfg.Get(data)
// 	if err != nil {
// 		t.Errorf("%v", err)
// 		return
// 	}
//
// 	expectedData := &struct {
// 		Key float64
// 	}{Key: 5}
//
// 	if !reflect.DeepEqual(data, expectedData) {
// 		t.Errorf("Config data doesn't match, expected: %v, got: %v", expectedData, data)
// 	}
//
// 	cfg = config.NewFromRaw(map[string]interface{}{
// 		"key": 5.3,
// 	})
//
// 	data2 := &struct {
// 		Key int
// 	}{}
// 	err = cfg.Get(data2)
// 	if err != nil {
// 		t.Errorf("Unexpected error: %v", err)
// 		return
// 	}
//
// 	expectedData2 := &struct {
// 		Key int
// 	}{Key: 5}
//
// 	if !reflect.DeepEqual(data2, expectedData2) {
// 		t.Errorf("Config data doesn't match, expected: %v, got: %v", expectedData2, data2)
// 	}
// }
