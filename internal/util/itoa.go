package util

import "strconv"

func IntToString(in int64) string {
	return strconv.FormatInt(in, 10)
}

func UintToString(in uint64) string {
	return strconv.FormatUint(in, 10)
}
