package easytime

import (
	"time"
)

var YmdHislayout = "2006-01-02 15:04:05"
var DateY = "2006"
var DateM = "01"
var DateD = "02"
var DateH = "15"
var DateI = "04"
var DateS = "05"
var DefaultDateTimeZone = "Asia/Shanghai"

/*
	date 时间
	layouts 时间格式
	dateTimeZone 时区
	return 根据格式返回对应时区的时间戳
*/
func DateToUnix(date string, layouts string, dateTimeZone string) int64 {

	if layouts == "" {
		layouts = YmdHislayout
	}

	if dateTimeZone == "" {
		dateTimeZone = DefaultDateTimeZone
	}

	loc, _ := time.LoadLocation(dateTimeZone)
	tmp, _ := time.ParseInLocation(layouts, date, loc)
	unix := tmp.Unix()

	return unix

}

/*
	unix 时间戳
	layouts 时间格式
	dateTimeZone 时区
	return 根据时间戳返回对应格式的时间
*/
func UnixToDate(unix int64, layouts string, dateTimeZone string) string {

	if layouts == "" {

		layouts = YmdHislayout

	}

	if dateTimeZone == "" {
		dateTimeZone = DefaultDateTimeZone
	}

	loc, _ := time.LoadLocation(dateTimeZone)

	date := time.Unix(unix, 0).In(loc).Format(layouts)

	return date

}

/*
	layouts 时间格式
	dateTimeZone 时区
	return 返回对应格式的当前时区的时间
*/

func ToDate(layouts string, dateTimeZone string) (string, error) {

	if layouts == "" {
		layouts = YmdHislayout
	}

	if dateTimeZone == "" {
		dateTimeZone = DefaultDateTimeZone
	}

	local, err := time.LoadLocation(dateTimeZone)

	if err != nil {
		return "", err

	}

	return time.Now().In(local).Format(layouts), nil
}

/*
	dateTimeZone 时区
	return 返回对应时区当前时间戳
*/
func ToUnix(dateTimeZone string) int64 {

	if dateTimeZone == "" {
		dateTimeZone = DefaultDateTimeZone
	}

	loc, _ := time.LoadLocation(dateTimeZone)

	return time.Now().In(loc).Unix()

}

//字符串日期格式加减
func AddStringTime(dateTime,xTime,layouts string) (time.Time,error){
	if layouts == "" {
		layouts = YmdHislayout
	}
	t,err := time.Parse(DefaultDateTimeZone,dateTime)
	if err != nil {
		return time.Time{},err
	}
	m, _ := time.ParseDuration(xTime)
	loc, _ := time.LoadLocation(DefaultDateTimeZone)
	return t.Add(m).In(loc),nil
}

//时间戳格式加减
func AddUnixTime(unixTime int64,xTime int64,layouts string) time.Time {
	if layouts == "" {
		layouts = YmdHislayout
	}
	if unixTime == 0 {
		unixTime = ToUnix("")
	}
	unixTime += xTime
	loc, _ := time.LoadLocation(DefaultDateTimeZone)
	return time.Unix(unixTime, 0).In(loc)
}

//当前日期加减
func AddNowTime(xTime,dateTimeZone string) time.Time{
	if dateTimeZone == "" {
		dateTimeZone = DefaultDateTimeZone
	}
	now := time.Now()
	m, _ := time.ParseDuration(xTime)
	loc, _ := time.LoadLocation(dateTimeZone)
	return now.Add(m).In(loc)
}

//当天最大时间
func MaxDateToString(xTime time.Time,layouts string,dateTimeZone string) string{
	if layouts == "" {
		layouts = YmdHislayout
	}
	if dateTimeZone == "" {
		dateTimeZone = DefaultDateTimeZone
	}
	loc, _ := time.LoadLocation(dateTimeZone)
	endTime := time.Date(xTime.Year(),xTime.Month(), xTime.Day(), 23, 59, 59, 0, loc).Format(layouts)
	return endTime
}

func MaxDateToUnix(xTime time.Time,dateTimeZone string) int64 {
	if dateTimeZone == "" {
		dateTimeZone = DefaultDateTimeZone
	}
	loc, _ := time.LoadLocation(dateTimeZone)
	endTime := time.Date(xTime.Year(),xTime.Month(), xTime.Day(), 23, 59, 59, 0, loc).Unix()
	return endTime
}

//当天最小时间
func MinDateToString(xTime time.Time,layouts string,dateTimeZone string) string {
	if layouts == "" {
		layouts = YmdHislayout
	}
	if dateTimeZone == "" {
		dateTimeZone = DefaultDateTimeZone
	}
	loc, _ := time.LoadLocation(dateTimeZone)
	startTime := time.Date(xTime.Year(),xTime.Month(), xTime.Day(), 0, 0, 0, 0, loc).Format(layouts)
	return startTime
}

func MinDateToUnix(xTime time.Time,dateTimeZone string) int64 {
	if dateTimeZone == "" {
		dateTimeZone = DefaultDateTimeZone
	}
	loc, _ := time.LoadLocation(dateTimeZone)
	startTime := time.Date(xTime.Year(),xTime.Month(), xTime.Day(), 0, 0, 0, 0, loc).Unix()
	return startTime
}