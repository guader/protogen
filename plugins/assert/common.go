package assert

import (
	"google.golang.org/protobuf/compiler/protogen"
)

func generateCommon(g *protogen.GeneratedFile) {
	g.P(`import (
	"fmt"
	"unicode/utf8"
)

type NumberType interface {
	~int32 | ~int64 | ~uint32 | ~uint64 | ~float32 | ~float64
}

type AssertFunc = func() error

// open: 0 for both close, 1 for left open only, 2 for right open only, 3 for both open.
func AssertNumberRange[T NumberType](name string, v T, min, max *T, open byte) error {
	if min != nil {
		if open&1 == 0 {
			// left close
			if v < *min {
				return fmt.Errorf("%s must be greater or equal than %v, value: %v", name, *min, v)
			}
		} else {
			// left open
			if v <= *min {
				return fmt.Errorf("%s must be greater than %v, value: %v", name, *min, v)
			}
		}
	}
	if max != nil {
		if open&2 == 0 {
			// right close
			if v > *max {
				return fmt.Errorf("%s must be less or equal than %v, value: %v", name, *max, v)
			}
		} else {
			// right open
			if v >= *max {
				return fmt.Errorf("%s must be less than %v, value: %v", name, *max, v)
			}
		}
	}
	return nil
}

func AssertNumberInSlice[T NumberType](name string, v T, vs ...T) error {
	for _, valid := range vs {
		if valid == v {
			return nil
		}
	}
	return fmt.Errorf("%s must be in %v, value: %v", name, vs, v)
}

func AssertNumberInMap[T NumberType](name string, v T, m map[T]struct{}) error {
	if _, ok := m[v]; ok {
		return nil
	}
	vs := make([]T, 0, len(m))
	for valid := range m {
		vs = append(vs, valid)
	}
	return fmt.Errorf("%s must be in %v, value: %v", name, vs, v)
}

// min, max: 0 for unlimited, rune count must be in range [min, max].
func AssertRuneCountRange(name string, v string, min, max int) error {
	n := utf8.RuneCountInString(v)
	if min > 0 && n < min {
		return fmt.Errorf("%s rune count must be greater or equal than %d, count: %d, value: %q", name, min, n, v)
	}
	if max > 0 && n > max {
		return fmt.Errorf("%s rune count must be less or equal than %d, count: %d, value: %q", name, max, n, v)
	}
	return nil
}

// min, max: 0 for unlimited, slice length must be in range [min, max].
func AssertSliceLengthRange[T any](name string, v []T, min, max int) error {
	n := len(v)
	if min > 0 && n < min {
		return fmt.Errorf("%s slice length must be greater or equal than %d, count: %d", name, min, n)
	}
	if max > 0 && n > max {
		return fmt.Errorf("%s slice length must be less or equal than %d, count: %d", name, max, n)
	}
	return nil
}`)
	g.P()
}
