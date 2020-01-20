package testDemo

import "strings"

func Split(str string, sep string) []string {
	split := strings.Split(str, sep)
	return split
}
