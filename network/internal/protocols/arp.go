package protocols

import (
	"encoding/binary"
	"net"
	"network_sniffer/internal/utils"
)

type ARP struct {
	Operation uint16
	SenderIP  string
	TargetIP  string
	SenderMAC string
	TargetMAC string
}

func ParseARP(data []byte) (ARP, bool) {
	if len(data) < 28 {
		return ARP{}, false
	}

	a := ARP{
		Operation: binary.BigEndian.Uint16(data[6:8]),
		SenderIP:  net.IP(data[14:18]).String(),
		TargetIP:  net.IP(data[24:28]).String(),
		SenderMAC: utils.FormatMAC(data[8:14]),
		TargetMAC: utils.FormatMAC(data[18:24]),
	}

	return a, true
}
