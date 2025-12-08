package protocols

import "encoding/binary"

type Ethernet struct {
	DstMAC   string
	SrcMAC   string
	EtherType uint16
}

func ParseEthernet(data []byte) (Ethernet, []byte, bool) {
	if len(data) < 14 {
		return Ethernet{}, nil, false
	}

	eth := Ethernet{
		DstMAC:   formatMAC(data[0:6]),
		SrcMAC:   formatMAC(data[6:12]),
		EtherType: binary.BigEndian.Uint16(data[12:14]),
	}

	return eth, data[14:], true
}

func formatMAC(b []byte) string {
	return formatMAC(b)
}
