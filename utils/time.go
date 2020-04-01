package utils

import (
	"strconv"
	"time"
)

func TimeFormat(t time.Time, formatType int64) (formatTime string) {

	if formatType == 0 {
		formatTime = t.Format("2006-01-02 15:04:05")
	} else if formatType == 1 {
		formatTime = t.Format("2006.01.02")
	} else {
		formatTime = t.Format("01月02日")
	}
	return formatTime
}

func TimeSpan(t time.Time) string {

	nowTime := time.Now().Unix()
	var timeStr string

	tTime := t.Unix()

	resTime := nowTime - tTime
	if resTime < 60 {

		timeStr = strconv.FormatInt(resTime,10) + "秒前"

	} else if resTime > 0 && resTime < 3600 {

		timeStr = strconv.FormatInt(resTime/60,10) + "分钟前"

	} else if resTime >= 3600 && resTime < 86400 {

		timeStr = strconv.FormatInt(resTime/3600, 10) + "小时前"

	} else if resTime >= 86400 && resTime < 604800 {

		timeStr = strconv.FormatInt(resTime/86400,10) + "天前"

	} else {

		timeStr = t.Format("01月02日")

	}

	return timeStr

}

func BetweenDays(startTime, endTime string) int64 {

	startT, _ := time.Parse("2006-01-02 15:04:05", startTime)
	endT, _ := time.Parse("2006-01-02 15:04:05", endTime)
	duringDays := int64((startT.Sub(endT).Hours()) / 24)

	return duringDays
}
