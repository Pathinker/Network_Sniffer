package protocols

import "bytes"

type HTTP struct {
	Method string
}

func ParseHTTP(data []byte) (HTTP, bool) {
	methods := [][]byte{
		[]byte("GET "),
		[]byte("POST "),
		[]byte("PUT "),
		[]byte("DELETE "),
	}

	for _, m := range methods {
		if bytes.HasPrefix(data, m) {
			return HTTP{Method: string(m[:len(m)-1])}, true
		}
	}

	return HTTP{}, false
}
