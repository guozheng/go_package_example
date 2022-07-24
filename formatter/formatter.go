package print

import (
	"bytes"
	"fmt"
	"strings"
)

func Format(num int) string {
	return fmt.Sprintf("The number is %d", num)
}

func Space(length int, parts ...string) string {
	if len(parts) == 0 {
		return ""
	} else if len(parts) == 1 {
		return parts[0]
	}
	totalLen := 0
	for _, v := range parts {
		totalLen += len(v)
	}
	minSpaces := len(parts) - 1
	if length < totalLen+minSpaces {
		length = totalLen + minSpaces
	}
	spares := length - totalLen
	per, extra := spares/minSpaces, spares%minSpaces
	var out bytes.Buffer
	out.WriteString(parts[0])
	for _, v := range parts[1:] {
		out.WriteString(strings.Repeat(" ", per))
		if extra > 0 {
			out.WriteString(" ")
			extra--
		}
		out.WriteString(v)
	}
	return out.String()
}
