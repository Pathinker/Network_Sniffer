package protocols

import "net"

type IPv6 struct {
	SrcIP    string
	DstIP    string
	NextHdr  uint8
}

func ParseIPv6(data []byte) (IPv6, []byte, bool) {
	if len(data) < 40 {
		return IPv6{}, nil, false
	}

	ip := IPv6{
		NextHdr: data[6],
		SrcIP:   net.IP(data[8:24]).String(),
		DstIP:   net.IP(data[24:40]).String(),
	}

	return ip, data[40:], true
}
