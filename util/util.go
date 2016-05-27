package util

import "time"

func NowMS() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
