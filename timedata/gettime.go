package timedata

import "time"

func GetTime() int {
	return time.Now().Day()
}
