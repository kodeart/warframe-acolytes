package internal

import (
	"strconv"
	"time"
)

const (
	mss = int64(1000)
	nss = int64(1000000)
)

// IntMillisToTime converts milliseconds to time.Time instant
func IntMillisToTime(ms int64) time.Time {
	return time.Unix(ms/mss, (ms%mss)*nss)
}

// MillisToTime converts string-milliseconds to time.TIme instant
func StrMillisToTime(ms string) (time.Time, error) {
	m, err := strconv.ParseInt(ms, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return IntMillisToTime(m), nil
}
