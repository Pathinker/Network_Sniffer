package protocols

import "bytes"

type SMB struct{}

func ParseSMB(data []byte) (SMB, bool) {
	if bytes.Contains(data, []byte("SMB")) {
		return SMB{}, true
	}
	return SMB{}, false
}
