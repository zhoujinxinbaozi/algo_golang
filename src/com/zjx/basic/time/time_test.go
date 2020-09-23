package time

import (
	"fmt"
	"testing"
	"time"
)

const (
	// seconds per day 一天的秒数
	daySeconds = 86400
)

// GetDiffDayByUnix 比较两个时间戳之前相差的天数
// @param lessUnix 较小的时间戳
// @param moreUnix 较大的时间戳
// @return 两个时间相差的天数
func DiffDayByUnix(lessUnix, moreUnix int64) int32 {
	var result int32
	if lessUnix >= moreUnix {
		return result
	}

	var difference int64
	difference = moreUnix - lessUnix

	if difference%daySeconds == 0 {
		return int32(difference / daySeconds)
	}

	return int32(difference/daySeconds + 1)
}

func TestDiffDayByUnix(t *testing.T) {
	startTime := "2016-09-10 13:00:00"
	endTime := "2016-09-11 13:01:00"

	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
	lessTime := t1.Unix()
	moreTime := t2.Unix()

	fmt.Println(DiffDayByUnix(lessTime, moreTime))
}