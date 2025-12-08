package protocols

import (
	"encoding/binary"
	"network_sniffer/internal/utils"
)

type Ethernet struct {
	DstMAC    string
	SrcMAC    string
	EtherType uint16
}

func ParseEthernet(data []byte) (Ethernet, []byte, bool) {
	if len(data) < 14 {
		return Ethernet{}, nil, false
	}

	eth := Ethernet{
		DstMAC:    utils.FormatMAC(data[0:6]),
		SrcMAC:    utils.FormatMAC(data[6:12]),
		EtherType: binary.BigEndian.Uint16(data[12:14]),
	}

	return eth, data[14:], true
}
