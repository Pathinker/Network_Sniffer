package protocols

import (
	"net"
)

type IPv4 struct {
	SrcIP    string
	DstIP    string
	Protocol uint8
}

func ParseIPv4(data []byte) (IPv4, []byte, bool) {
	if len(data) < 20 {
		return IPv4{}, nil, false
	}

	ihl := (data[0] & 0x0F) * 4
	if len(data) < int(ihl) {
		return IPv4{}, nil, false
	}

	ip := IPv4{
		SrcIP:    net.IP(data[12:16]).String(),
		DstIP:    net.IP(data[16:20]).String(),
		Protocol: data[9],
	}

	return ip, data[ihl:], true
}
