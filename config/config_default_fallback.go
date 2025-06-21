package config

// GetStringDefault implementation using "GetString" method with manual input default value
func (c *DefaultConfig) GetStringDefault(name, val string) (result string) {
	result, ok := c.GetString(name)
	if !ok {
		result = val
	}
	return
}

// GetIntDefault implementation using "GetInt" method with manual input default value
func (c *DefaultConfig) GetIntDefault(name string, val int) (result int) {
	result, ok := c.GetInt(name)
	if !ok {
		result = val
	}
	return
}

// GetBoolDefault implementation using "GetBool" method with manual input default value
func (c *DefaultConfig) GetBoolDefault(name string, val bool) (result bool) {
	result, ok := c.GetBool(name)
	if !ok {
		result = val
	}
	return
}

// GetFloatDefault implementation using "GetFloat" method with manual input default value
func (c *DefaultConfig) GetFloatDefault(name string, val float64) (result float64) {
	result, ok := c.GetFloat(name)
	if !ok {
		result = val
	}
	return
}
