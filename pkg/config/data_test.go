package config_test

type sampleConfigType struct {
	Name    string
	Profile userProfile
}

type userProfile struct {
	Sex      string
	Age      int
	Married  bool
	Children []childProfile
}

type childProfile struct {
	Name   string
	Weight float32
	Age    int8
}

// nolint: gochecknoglobals // this variable is actually read-only so it's ok to use it.
var sampleConfigDataJSON = []byte(`{
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

// nolint: gochecknoglobals // this variable is actually read-only so it's ok to use it.
var sampleConfigDataYAML = []byte(`
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

// nolint: gochecknoglobals // this variable is actually read-only so it's ok to use it.
var sampleConfigDataTOML = []byte(`
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

// nolint: gochecknoglobals // this variable is actually read-only so it's ok to use it.
var sampleConfigDataWrongJSON = []byte(`{
  "name": "Peter",
`)

// nolint: gochecknoglobals // this variable is actually read-only so it's ok to use it.
var sampleConfigDataWrongYAML = []byte(`
name: - Peter
`)

// nolint: gochecknoglobals // this variable is actually read-only so it's ok to use it.
var sampleConfigDataWrongTOML = []byte(`
name = Peter
`)

// nolint: gochecknoglobals // this variable is actually read-only so it's ok to use it.
var expectedSampleConfig = sampleConfigType{
	Name: "Peter",
	Profile: userProfile{
		Sex:     "m",
		Age:     32,
		Married: true,
		Children: []childProfile{
			{Name: "George", Weight: 5.4, Age: 5},
			{Name: "Olivia", Weight: 12.2, Age: 12},
		},
	},
}

// nolint: gochecknoglobals // this variable is actually read-only so it's ok to use it.
var expectedSampleRawDataJSON = map[string]interface{}{
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

// nolint: gochecknoglobals // this variable is actually read-only so it's ok to use it.
var expectedSampleRawDataYAML = map[string]interface{}{
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

// nolint: gochecknoglobals // this variable is actually read-only so it's ok to use it.
var expectedSampleRawDataTOML = map[string]interface{}{
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
