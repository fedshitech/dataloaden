//go:generate dataloaden -keys int github.com/vektah/dataloaden/example.User

package intkeys

import (
	"strconv"
	"time"

	"github.com/vektah/dataloaden/example"
)

func NewLoader() *UserLoader {
	return &UserLoader{
		wait:     2 * time.Millisecond,
		maxBatch: 100,
		fetch: func(keys []int) ([]*example.User, []error) {
			users := make([]*example.User, len(keys))
			errors := make([]error, len(keys))

			for i, key := range keys {
				users[i] = &example.User{ID: strconv.Itoa(key), Name: "user " + strconv.Itoa(key)}
			}
			return users, errors
		},
	}
}
