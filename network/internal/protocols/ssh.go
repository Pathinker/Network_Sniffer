package protocols

import "bytes"

type SSH struct {
	Version string
}

func ParseSSH(data []byte) (SSH, bool) {
	if bytes.HasPrefix(data, []byte("SSH-")) {
		return SSH{Version: string(data)}, true
	}
	return SSH{}, false
}
