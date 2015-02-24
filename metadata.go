package fireload

// Metadata represents key-value pairs of metadata associated with a Namespace
type Metadata map[string]interface{}

// Set adds the given key-value pair
func (m Metadata) Set(key string, value interface{}) {
	m[key] = value
}

// Get retrieves the value stored under key.
func (m Metadata) Get(key string) (interface{}, bool) {
	val, ok := m[key]
	return val, ok
}
