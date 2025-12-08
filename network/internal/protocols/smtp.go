package protocols

import "bytes"

type SMTP struct {
	Banner string
}

func ParseSMTP(data []byte) (SMTP, bool) {
	if bytes.HasPrefix(data, []byte("220")) {
		return SMTP{Banner: string(data)}, true
	}
	return SMTP{}, false
}
