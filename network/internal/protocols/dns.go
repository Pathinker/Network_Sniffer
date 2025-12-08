package protocols

import "encoding/binary"

type DNS struct {
	ID      uint16
	QR      uint8
	OpCode  uint8
}

func ParseDNS(data []byte) (DNS, bool) {
	if len(data) < 12 {
		return DNS{}, false
	}

	flags := binary.BigEndian.Uint16(data[2:4])

	return DNS{
		ID:     binary.BigEndian.Uint16(data[0:2]),
		QR:     uint8((flags >> 15) & 0x01),
		OpCode: uint8((flags >> 11) & 0x0F),
	}, true
}
