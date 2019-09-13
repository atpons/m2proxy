package util

import "fmt"

var (
	Debug = 0
)

func StringBytes(b []byte) string {
	str := ""
	for _, byte := range b {
		str += fmt.Sprintf("0x%x ", byte)
	}
	return str
}
