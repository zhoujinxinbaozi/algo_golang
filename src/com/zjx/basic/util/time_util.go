/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/9/22 3:17 下午
 */

package util

import (
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

// CalTimeConsuming 计算耗时
// @param startTime 函数起始时间
// @return 耗时
func CalTimeConsuming(startTime time.Time) time.Duration {
	timeConsuming := time.Since(startTime)
	return timeConsuming
}

// CalTimeCost 计算耗时
// @param nil
// @return 耗时
func CalTimeCost() func() time.Duration {
	start := time.Now()
	return func() time.Duration{
		tc:=time.Since(start)
		return tc
	}
}
