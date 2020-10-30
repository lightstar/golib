package config

// NewInner function creates new configuration service using data under some key in another already existing
// configuration service.
func NewInner(key string, config *Config) (*Config, error) {
	data, err := config.GetRawByKey(key)
	if err != nil {
		return nil, err
	}

	dataMap, ok := data.(map[string]interface{})
	if !ok {
		return nil, ErrNotMap
	}

	return NewFromRaw(dataMap), nil
}
