package go_utils

import (
	"strconv"
)

func Int2String(i int) string {
	return strconv.Itoa(i)
}

func String2Int(str string) (int, error) {
	return strconv.Atoi(str)
}
