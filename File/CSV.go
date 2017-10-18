package File

import (
	"bytes"
)

func SliceStringToCsvString(strings []string) string {
	var buffer bytes.Buffer

	var size = len(strings)
	for i := 0; i < size; i++ {
		buffer.WriteString(strings[i])
		if i < size - 1 {
			buffer.WriteString("\r\n")
		}
	}

	var byteArray = buffer.Bytes()
	s := string(byteArray[:])
	return s
}
