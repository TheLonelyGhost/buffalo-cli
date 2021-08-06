//go:generate mapgen -name "plug" -zero "plug{}" -go-type "plug" -pkg "" -a "nil" -b "nil" -c "nil" -bb "nil" -destination "plugcmds"
// Code generated by github.com/gobuffalo/mapgen. DO NOT EDIT.

package plugcmds

import (
	"sort"
	"sync"
)

// plugMap wraps sync.Map and uses the following types:
// key:   string
// value: plug
type plugMap struct {
	data sync.Map
}

// Delete the key from the map
func (m *plugMap) Delete(key string) {
	m.data.Delete(key)
}

// Load the key from the map.
// Returns plug or bool.
// A false return indicates either the key was not found
// or the value is not of type plug
func (m *plugMap) Load(key string) (plug, bool) {
	i, ok := m.data.Load(key)
	if !ok {
		return plug{}, false
	}
	s, ok := i.(plug)
	return s, ok
}

// LoadOrStore will return an existing key or
// store the value if not already in the map
func (m *plugMap) LoadOrStore(key string, value plug) (plug, bool) {
	i, _ := m.data.LoadOrStore(key, value)
	s, ok := i.(plug)
	return s, ok
}

// Range over the plug values in the map
func (m *plugMap) Range(f func(key string, value plug) bool) {
	m.data.Range(func(k, v interface{}) bool {
		key, ok := k.(string)
		if !ok {
			return false
		}
		value, ok := v.(plug)
		if !ok {
			return false
		}
		return f(key, value)
	})
}

// Store a plug in the map
func (m *plugMap) Store(key string, value plug) {
	m.data.Store(key, value)
}

// Keys returns a list of keys in the map
func (m *plugMap) Keys() []string {
	var keys []string
	m.Range(func(key string, value plug) bool {
		keys = append(keys, key)
		return true
	})
	sort.Strings(keys)
	return keys
}
