# GoTime
Time and date formatting in Go can be confusing to read - choose **02** instead of **03** and you get a month not an hour!

Formatting verbs from Python are a lot easier to distinguish between.  This small package allows you to use Python verbs instead of Go verbs.

## Usage
Using the library is as simple as importing then using the one exported function `FormatDateTime`:
```go
// time1 is 12:01:01 on 1st Jan 2018, GMT
time1 := time.Date(2018, time.January, 01, 12, 01, 01, 0, time.UTC)
fmt.Println(FormatDateTime(time1, "%Y/%m/%d %H:%M:%S"))
// Output: 2018/01/01 12:01:01

// time2 is 9:15:00 on 18th July 2021, GMT
time2 := time.Date(2021, time.July, 18, 9, 15, 0, 0, time.UTC)
fmt.Println(FormatDateTime(time2, "It's a %A during the month of %B"))
// Output: It's a Sunday during the month of July
```
