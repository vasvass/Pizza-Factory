package ovens

import "time"

type Sleeper interface {
	Sleep(duration time.Duration)
}
