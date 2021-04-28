/**
 *  @Author: JinxinZhou
 *  @Date  : 2021/4/15 12:09 下午
 */

package lock

import "sync/atomic"

type TimeLocker struct {
	begin int64 // begin time (unix second)
	end   int64 // end time = begin + interval
}

// NewTimeLocker create a new TimeLocker
func NewTimeLocker() *TimeLocker {
	return &TimeLocker{}
}

// LockAfter will locked from begin to begin+interval (seconds)
func (tl *TimeLocker) LockAfter(begin, interval int64) {
	if swap := atomic.CompareAndSwapInt64(&tl.begin, tl.begin, begin); swap {
		atomic.StoreInt64(&tl.end, begin+interval)
	}
}

// IsLocked lock after interval seconds. t is time point to check lock status
func (tl *TimeLocker) IsLocked(t int64) bool {
	return atomic.LoadInt64(&tl.end) > t
}
