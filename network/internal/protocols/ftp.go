package protocols

import "bytes"

type FTP struct {
	Banner string
}

func ParseFTP(data []byte) (FTP, bool) {
	if bytes.HasPrefix(data, []byte("220")) {
		return FTP{Banner: string(data)}, true
	}
	return FTP{}, false
}
