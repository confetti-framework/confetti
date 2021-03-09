package getters

import (
	"fmt"
	"strings"
)

// StringList casts comma separated string to a slice
type StringList []string

// StringList casts a slice to a comma separated string
func (s *StringList) String() string {
	return fmt.Sprintf("%v", *s)
}

// Set casts comma separated string to a slice
func (s *StringList) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}

// Get returns the value of type which must be
// the same type as defined in a field of a command.
func (s *StringList) Get() interface{} {
	return []string(*s)
}
