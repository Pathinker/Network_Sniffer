package protocols

import "encoding/binary"

type UDP struct {
	SrcPort uint16
	DstPort uint16
	Length  uint16
}

func ParseUDP(data []byte) (UDP, bool) {
	if len(data) < 8 {
		return UDP{}, false
	}

	u := UDP{
		SrcPort: binary.BigEndian.Uint16(data[0:2]),
		DstPort: binary.BigEndian.Uint16(data[2:4]),
		Length:  binary.BigEndian.Uint16(data[4:6]),
	}

	return u, true
}
