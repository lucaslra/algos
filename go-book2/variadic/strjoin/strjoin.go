package strjoin

import "strings"

func Join(sep string, str ...string) string {
	return strings.Join(str, sep)
}
