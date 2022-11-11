package util

import (
	"github.com/spf13/viper"
	"time"
)

// TimeToString time.Time -> string
func TimeToString(t time.Time) string {
	return t.Format(viper.GetString("time.format"))
}

// GetUTCTime UTC format time now
func GetUTCTime() time.Time {
	return time.Time.UTC(time.Now())
}

// ParseUTCTime UTC -> local time
func ParseUTCTime(t *time.Time, offset int) time.Time {
	newT := t.Add(time.Hour * time.Duration(offset))
	return newT
}
