package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var year int
var month int
var fs *flag.FlagSet
var showVersion, showUsage bool
var xr map[int]string // 星期到人名的映射
var slice []string

func init() {
	slice = append(slice, "朱开郑")
	slice = append(slice, "单会花")
	slice = append(slice, "李伟山")
	slice = append(slice, "吴桂花")
	slice = append(slice, "姜林杰")
	slice = append(slice, "李设桂")
	slice = append(slice, "周伟")
	slice = append(slice, "干卫众")
	xr = map[int]string{
		-1: "朱开郑",     //星期六
		0:  "李伟山",     //星期日
		1:  "吴桂花",     //星期一
		2:  "姜林杰",     //星期二
		3:  "杨文生,李设桂", //星期三
		4:  "周伟,单会花",  //星期四
		5:  "干卫众",     //星期五
	}
	//初始化解析参数
	fs = flag.NewFlagSet("calendar", flag.ExitOnError)
	fs.BoolVar(&showUsage, "help", false, "Show this message")
	fs.BoolVar(&showUsage, "h", false, "Show this message")
	fs.BoolVar(&showVersion, "version", false, "Print version information.")
	fs.BoolVar(&showVersion, "v", false, "Print version information.")
	fs.IntVar(&year, "year", 2020, "Input year.")
	fs.IntVar(&year, "y", 2020, "Input year.")
	fs.IntVar(&month, "month", 2, "Input month.")
	fs.IntVar(&month, "m", 2, "Input month.")
}

func perpetualCalendar(year int, month int) {
	var monthDays int = 0
	var totalDyas int = 0
	var isLeapYear bool = false
	fmt.Println("query date is:", year, month)

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		isLeapYear = true
		fmt.Println(year, "is leap year")
	} else {
		isLeapYear = false
		fmt.Println(year, "is not lear year")
	}

	//计算距离1900年的天数
	for i := 1900; i < year; i++ {
		if i%400 == 0 || (i%4 == 0 && i%100 != 0) {
			totalDyas += 366
		} else {
			totalDyas += 365
		}
	}

	//计算截至到当月1号前的总天数
	for j := 1; j <= month; j++ {
		switch j {
		case 1, 3, 5, 7, 8, 10, 12:
			monthDays = 31
			break
		case 2:
			if isLeapYear {
				monthDays = 29
			} else {
				monthDays = 28
			}
			break
		case 4, 6, 9, 11:
			monthDays = 30
			break
		default:
			fmt.Println("input month is error.")
		}

		if j != month {
			totalDyas += monthDays
		}
	}

	//计算当月1号是星期几
	var weekDay int = 0
	weekDay = 1 + totalDyas%7
	if weekDay == 7 {
		weekDay = 0
	}

	fmt.Println("weekDay is: ", weekDay)

	//展示日期
	fmt.Println("星期日\t" + "星期一\t" + "星期二\t" + "星期三\t" + "星期四\t" + "星期五\t" + "星期六\t")
	for k := 0; k < weekDay; k++ {
		fmt.Printf("\t")
	}
	target := 0
	for m := 1; m <= monthDays; m++ {
		fmt.Printf("%d ", m)
		//fmt.Printf("%d", (weekDay+m)%7-1)
		// 当前日期和休息日不相同
		for {
			split := strings.Split(xr[(weekDay+m)%7-1], ",")
			if !check(split, slice, target) {
				if target == len(slice)-1 {
					fmt.Printf("%s,t=%d\t", slice[target], target)
					target = 0
					break
				} else {
					fmt.Printf("%s, t=%d\t", slice[target], target)
					target++
					target = target % len(slice)
					break
				}
			} else {
				//fmt.Printf("%s,%d,%d,%s", slice[target], target, (weekDay+m)%7-1, xr[(weekDay+m)%7-1])
				change(&slice, target)
			}
		}

		if (weekDay+m)%7 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func change(slice *[]string, target int) {
	if target == len(*slice)-1 {
		tmp := (*slice)[0]
		(*slice)[0] = (*slice)[target]
		(*slice)[target] = tmp
	} else {
		tmp := (*slice)[target]
		(*slice)[target] = (*slice)[target+1]
		(*slice)[target+1] = tmp
	}

}

func check(split []string, slice []string, target int) bool {
	if len(split) == 1 {
		if split[0] == slice[target] {
			return true
		}
	} else {
		if split[0] == slice[target] || split[1] == slice[target] {
			return true
		}
	}
	return false
}

func main() {
	//开启main 函数的解析功能
	fs.Parse(os.Args[1:])
	if showUsage {
		fs.PrintDefaults()
		os.Exit(0)
	}
	if showVersion {
		fmt.Println("version:v1.0")
		os.Exit(0)
	}
	perpetualCalendar(year, month)
}
