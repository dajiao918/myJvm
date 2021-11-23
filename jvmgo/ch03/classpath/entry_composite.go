package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(path string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(path, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (c CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range c {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, err
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (c CompositeEntry) String() string {
	strs := make([]string, len(c))
	for i, str := range c {
		strs[i] = str.String()
	}
	return strings.Join(strs, pathListSeparator)
}
