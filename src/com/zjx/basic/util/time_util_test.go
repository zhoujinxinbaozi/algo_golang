/**
 *  @Author: JinxinZhou
 *  @Date  : 2020/9/22 4:22 下午
 */

package util

import (
	"fmt"
	"testing"
	"time"
)

// TestGetDiffDayByUnix
func TestDiffDayByUnix(t *testing.T) {

	startTime := "2016-09-10 13:00:00"
	endTime := "2016-09-11 13:01:00"

	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
	t2, _ := time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
	tt, _ := time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
	ttt, _ := time.Parse("2006-01-02 15:04:05", startTime)
	fmt.Println(tt.Unix())
	fmt.Println(ttt.Unix())


	lessTime := t1.Unix()
	moreTime := t2.Unix()

	fmt.Println(DiffDayByUnix(lessTime, moreTime))
}

// TestCalTimeConsuming
func TestCalTimeConsuming(t *testing.T) {
	nowTime := time.Now()
	time.Sleep(time.Duration(2) * time.Millisecond)
	timeConsuming := CalTimeConsuming(nowTime)
	fmt.Println(timeConsuming)
}

// TestTimeCost
func TestCalTimeCost(t *testing.T) {

	tcFunc := CalTimeCost()
	defer func() {
		fmt.Println("qqq",tcFunc())
	}()

	time.Sleep(time.Duration(1) * time.Second)

}
