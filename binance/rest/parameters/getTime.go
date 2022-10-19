package parameters

import "time"

func getTimestamp() int64 {
	return time.Now().UTC().Unix() * 1000
}
