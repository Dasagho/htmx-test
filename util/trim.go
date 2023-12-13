package util

import "strings"

func TrimUpTo(trim, toTrim string) string {
	index := strings.Index(toTrim, trim)
	if index == -1 {
		// "views" no se encuentra en la cadena.
		return toTrim
	}
	return toTrim[index:]
}
