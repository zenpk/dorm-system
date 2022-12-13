package util

import (
	"github.com/spf13/viper"
	"time"
)

// TimeToString UTC time.Time -> local time string
func TimeToString(t time.Time) string {
	offset := viper.GetInt("datetime.timezone")
	newT := ParseUTCTime(t, offset)
	return newT.Format(viper.GetString("datetime.format.time"))
}

// StringToTime local time string -> UTC time.Time
func StringToTime(str string) (time.Time, error) {
	offset := viper.GetInt("datetime.timezone")
	var t time.Time
	var err error
	t, err = time.Parse(viper.GetString("datetime.format.time"), str)
	if err != nil {
		// if time format not working, try date format
		t, err = time.Parse(viper.GetString("datetime.format.date"), str)
		if err != nil { // still not working, try all possible formats
			t, err = time.Parse("2006-01-02T15:04:05", str) // ISO local
			if err != nil {
				t, err = time.Parse("2006-01-02T15:04:05Z", str) // ISO global
				if err != nil {
					t, err = time.Parse("2006-01-02T15:04:05-0700", str) // ISO region
					if err != nil {
						return time.Time{}, err
					}
				}
			}
		}
	}
	newT := FormatUTCTime(t, offset)
	return newT, nil
}

// GetUTCTime UTC format time now
func GetUTCTime() time.Time {
	return time.Time.UTC(time.Now())
}

// ParseUTCTime UTC -> local time
func ParseUTCTime(t time.Time, offset int) time.Time {
	newT := t.Add(time.Hour * time.Duration(offset))
	return newT
}

// FormatUTCTime local time -> UTC
func FormatUTCTime(t time.Time, offset int) time.Time {
	newT := t.Add(time.Hour * time.Duration(-offset))
	return newT
}
