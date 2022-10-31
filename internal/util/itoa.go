package util

import "strconv"

func ToString64(in int64) string {
	return strconv.FormatInt(in, 10)
}
