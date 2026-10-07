package outport

import "time"

type Clock interface {
	Now() time.Time
}
