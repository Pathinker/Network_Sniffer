package protocols

import "encoding/binary"

type VLAN struct {
	TCI       uint16
	EtherType uint16
}

func ParseVLAN(data []byte) (VLAN, []byte, bool) {
	if len(data) < 4 {
		return VLAN{}, nil, false
	}

	v := VLAN{
		TCI:       binary.BigEndian.Uint16(data[0:2]),
		EtherType: binary.BigEndian.Uint16(data[2:4]),
	}

	return v, data[4:], true
}
