package common

import (
	"github.com/astaxie/beego/logs"
	"strings"
	"time"
)

const DATE_FORMAT = "2006-01-02 15:04:05"

func GetCurTime() string {
	return time.Now().Format(DATE_FORMAT)
}

func Catchs() {
	if err := recover(); err != nil {
		logs.Error("The program is abnormal, err: ", err)
	}
}

//TrimString Remove the \n \r \t spaces in the string
func TrimString(str string) string {
	str = strings.Replace(str, " ", "", -1)
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	return str
}

//TrimStringNR Remove the \n \r in the string
func TrimStringNR(str string) string {
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\r", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	return str
}

//TimeStrToInt parse time string to unix nano
func TimeStrToInt(ts, layout string) int64 {
	if ts == "" {
		return 0
	}
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	t, err := time.ParseInLocation(layout, ts, time.Local)
	if err != nil {
		logs.Error(err)
		return 0
	}
	return t.Unix()
}

func TimeToLocal(times, layout string) string {
	if times == "" {
		return ""
	}
	if layout == "" {
		layout = "2006-01-02 15:04:05"
	}
	t, err := time.ParseInLocation(layout, times, time.Local)
	if err != nil {
		logs.Error(err)
		return ""
	}
	return t.Format(DATE_FORMAT)
}