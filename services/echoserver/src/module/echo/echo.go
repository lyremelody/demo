package echo

import "time"

func GetNowTime() string {
	return time.Now().String()
}
