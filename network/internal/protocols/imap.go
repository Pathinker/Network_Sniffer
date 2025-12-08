package protocols

import "bytes"

type IMAP struct {
	Banner string
}

func ParseIMAP(data []byte) (IMAP, bool) {
	if bytes.HasPrefix(data, []byte("* OK")) {
		return IMAP{Banner: string(data)}, true
	}
	return IMAP{}, false
}
