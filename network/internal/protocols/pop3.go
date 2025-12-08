package protocols

import "bytes"

type POP3 struct {
	Banner string
}

func ParsePOP3(data []byte) (POP3, bool) {
	if bytes.HasPrefix(data, []byte("+OK")) {
		return POP3{Banner: string(data)}, true
	}
	return POP3{}, false
}
