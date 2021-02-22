package flag_type

import (
	"fmt"
	"github.com/confetti-framework/errors"
	"strconv"
	"strings"
)

type IntList []int

func (s *IntList) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *IntList) Set(value string) error {
	//goland:noinspection GoPreferNilSlice
	result := []int{}
	for _, part := range strings.Split(value, ",") {
		v, err := strconv.ParseInt(part, 0, 0)
		if err != nil {
			return errors.New("unable to cast %#v of type %T to int", part, part)
		}

		result = append(result, int(v))
	}
	*s = result
	return nil
}

// Get returns the value of type which must be
// the same type as defined in a field of a command.
func (s *IntList) Get() interface{} {
	return []int(*s)
}

