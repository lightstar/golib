package configtest

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/config"
)

// SampleConfigType structure used for sample configuration.
type SampleConfigType struct {
	Name    string
	Profile UserProfile
}

// UserProfile structure used as inner type in sample configuration.
type UserProfile struct {
	Sex      string
	Age      int
	Married  bool
	Children []ChildProfile
}

// ChildProfile structure used as inner aggregate type sample configuration.
type ChildProfile struct {
	Name   string
	Weight float32
	Age    int8
}

// SampleConfigDataJSON variable contains sample configuration as raw JSON data.
//
//nolint:gochecknoglobals // it's ok for sample test data to be global.
var SampleConfigDataJSON = []byte(`{
  "name": "Peter",
  "profile": {
    "sex": "m",
    "age": 32,
    "married": true,
    "children": [
      { "name": "George", "weight": 5.4, "age": 5},
      { "name": "Olivia", "weight": 12.2, "age": 12}
    ]
  }
}`)

// SampleConfigDataYAML variable contains sample configuration as raw YAML data.
//
//nolint:gochecknoglobals // it's ok for sample test data to be global.
var SampleConfigDataYAML = []byte(`
name: Peter
profile:
  sex: m
  age: 32
  married: true
  children:
    - name: George
      weight: 5.4
      age: 5
    - name: Olivia
      weight: 12.2
      age: 12
`)

// SampleConfigDataTOML variable contains sample configuration as raw TOML data.
//
//nolint:gochecknoglobals // it's ok for sample test data to be global.
var SampleConfigDataTOML = []byte(`
name = "Peter"

[profile]
sex = "m"
age = 32
married = true

[[profile.children]]
name = "George"
weight = 5.4
age = 5

[[profile.children]]
name = "Olivia"
weight = 12.2
age = 12
`)

// SampleConfigDataWrongJSON variable contains sample wrong JSON data.
//
//nolint:gochecknoglobals // it's ok for sample test data to be global.
var SampleConfigDataWrongJSON = []byte(`{
  "name": "Peter",
`)

// SampleConfigDataWrongYAML variable contains sample wrong YAML data.
//
//nolint:gochecknoglobals // it's ok for sample test data to be global.
var SampleConfigDataWrongYAML = []byte(`
name: - Peter
`)

// SampleConfigDataWrongTOML variable contains sample wrong TOML data.
//
//nolint:gochecknoglobals // it's ok for sample test data to be global.
var SampleConfigDataWrongTOML = []byte(`
name = Peter
`)

// ExpectedSampleConfig variable contains expected result of configuration parsing.
//
//nolint:gochecknoglobals,gomnd // it's ok for sample test data to be global and use raw numbers.
var ExpectedSampleConfig = SampleConfigType{
	Name: "Peter",
	Profile: UserProfile{
		Sex:     "m",
		Age:     32,
		Married: true,
		Children: []ChildProfile{
			{Name: "George", Weight: 5.4, Age: 5},
			{Name: "Olivia", Weight: 12.2, Age: 12},
		},
	},
}

// ExpectedSampleRawDataJSON variable contains expected raw result of configuration parsing by JSON encoder.
//
//nolint:gochecknoglobals,gomnd // it's ok for sample test data to be global and use raw numbers.
var ExpectedSampleRawDataJSON = map[string]interface{}{
	"name": "Peter",
	"profile": map[string]interface{}{
		"sex":     "m",
		"age":     32.,
		"married": true,
		"children": []interface{}{
			map[string]interface{}{"name": "George", "weight": 5.4, "age": 5.},
			map[string]interface{}{"name": "Olivia", "weight": 12.2, "age": 12.},
		},
	},
}

// ExpectedSampleRawDataYAML variable contains expected raw result of configuration parsing by YAML encoder.
//
//nolint:gochecknoglobals,gomnd // it's ok for sample test data to be global and use raw numbers.
var ExpectedSampleRawDataYAML = map[string]interface{}{
	"name": "Peter",
	"profile": map[string]interface{}{
		"sex":     "m",
		"age":     32,
		"married": true,
		"children": []interface{}{
			map[string]interface{}{"name": "George", "weight": 5.4, "age": 5},
			map[string]interface{}{"name": "Olivia", "weight": 12.2, "age": 12},
		},
	},
}

// ExpectedSampleRawDataTOML variable contains expected raw result of configuration parsing by TOML encoder.
//
//nolint:gochecknoglobals,gomnd // it's ok for sample test data to be global and use raw numbers.
var ExpectedSampleRawDataTOML = map[string]interface{}{
	"name": "Peter",
	"profile": map[string]interface{}{
		"sex":     "m",
		"age":     int64(32),
		"married": true,
		"children": []map[string]interface{}{
			{"name": "George", "weight": 5.4, "age": int64(5)},
			{"name": "Olivia", "weight": 12.2, "age": int64(12)},
		},
	},
}

// TestSampleConfig function can be used to test provided configuration against expected raw data.
//
//nolint:forcetypeassert // we know what is inside expectedRawData structure. if not - test will fail anyway.
func TestSampleConfig(t *testing.T, cfg *config.Config, expectedRawData map[string]interface{}) {
	t.Helper()

	require.Equal(t, expectedRawData, cfg.GetRaw())

	rawDataByEmptyKey, err := cfg.GetRawByKey("")
	require.NoError(t, err)
	require.Equal(t, interface{}(expectedRawData), rawDataByEmptyKey)

	rawChildren, err := cfg.GetRawByKey("profile.children")
	require.NoError(t, err)
	require.Equal(t, expectedRawData["profile"].(map[string]interface{})["children"], rawChildren)

	rawChildren, err = cfg.GetRawByKey(".profile.children")
	require.NoError(t, err)
	require.Equal(t, expectedRawData["profile"].(map[string]interface{})["children"], rawChildren)

	var data SampleConfigType

	require.NoError(t, cfg.Get(&data))
	require.Equal(t, ExpectedSampleConfig, data)

	var children []ChildProfile

	require.NoError(t, cfg.GetByKey("profile.children", &children))
	require.Equal(t, ExpectedSampleConfig.Profile.Children, children)

	require.NoError(t, cfg.GetByKey(".profile.children", &children))
	require.Equal(t, ExpectedSampleConfig.Profile.Children, children)
}
