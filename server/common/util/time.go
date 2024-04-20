package util

import (
	"time"
)

const layout = "2006-01-02 15:04:05"

func StringToTimestamp(str string) (int64, error) {
	t, err := time.Parse(layout, str)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}
