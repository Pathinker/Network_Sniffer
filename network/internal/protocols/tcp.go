package protocols

import "encoding/binary"

type TCP struct {
	SrcPort uint16
	DstPort uint16
}

func ParseTCP(data []byte) (TCP, bool) {
	if len(data) < 20 {
		return TCP{}, false
	}

	tcp := TCP{
		SrcPort: binary.BigEndian.Uint16(data[0:2]),
		DstPort: binary.BigEndian.Uint16(data[2:4]),
	}

	return tcp, true
}
