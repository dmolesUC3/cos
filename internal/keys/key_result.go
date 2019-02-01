package keys

import (
	"fmt"
	"strings"
)

type KeyResult struct {
	List  KeyList
	Index int
	Key   string
	Error error
}

func (f *KeyResult) Success() bool {
	return f.Error == nil
}

func (f *KeyResult) Pretty() string {
	if f.Success() {
		return fmt.Sprintf("%#v (%d of %d from %v) succeeded",
			f.Key,
			1+f.Index,
			f.List.Count(),
			f.List.Name(),
		)
	}

	return fmt.Sprintf("%#v (%d of %d from %v) failed: %v",
		f.Key,
		1+f.Index,
		f.List.Count(),
		f.List.Name(),
		strings.Replace(f.Error.Error(), "\n", "\\n", -1),
	)
}
