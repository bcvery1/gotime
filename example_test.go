package gotime

import (
	"fmt"
	"time"
)

func ExampleFormatDateTime() {
	// time1 is 12:01:01 on 1st Jan 2018, GMT
	time1 := time.Date(2018, time.January, 01, 12, 01, 01, 0, time.UTC)
	fmt.Println(FormatDateTime(time1, "%Y/%m/%d %H:%M:%S"))

	// time2 is 9:15:00 on 18th July 2021, GMT
	time2 := time.Date(2021, time.July, 18, 9, 15, 0, 0, time.UTC)
	fmt.Println(FormatDateTime(time2, "It's a %A during the month of %B"))

	// Output:
	// 2018/01/01 12:01:01
	// It's a Sunday during the month of July
}
