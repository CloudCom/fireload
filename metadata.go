package fireload

// Metadata .
type Metadata map[string]interface{}

// Set .
func (m Metadata) Set(key string, value interface{}) {
	m[key] = value
}

// Get .
func (m Metadata) Get(key string) (interface{}, bool) {
	val, ok := m[key]
	return val, ok
}
