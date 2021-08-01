package convert

import (
	"strconv"
	"strings"
)

func JoinInt64(a []int64, sep string) string {
	if len(a) == 0 {
		return ""
	}

	s := make([]string, len(a))
	for idx, val := range a {
		s[idx] = strconv.FormatInt(val, 10)
	}
	return strings.Join(s, sep)
}
