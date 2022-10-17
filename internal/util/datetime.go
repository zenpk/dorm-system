package util

import (
	"github.com/spf13/viper"
	"time"
)

// GetUTCTime 获取当前时间对应的 UTC 格式
func GetUTCTime() string {
	return time.Time.UTC(time.Now()).Format(viper.GetString("time.format"))
}

// ParseUTCTime 将 UTC 格式时间转换为当前的时间，offset 为时差 (北京时间 +8)
func ParseUTCTime(timeStr string, offset int) (string, error) {
	timeObj, err := time.Parse(viper.GetString("time.format"), timeStr)
	if err != nil {
		return "", err
	}
	timeObj = timeObj.Add(time.Hour * time.Duration(offset))
	return timeObj.Format(viper.GetString("time.format")), nil
}
