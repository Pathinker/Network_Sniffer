package decoder

import (
	"fmt"
	"network_sniffer/internal/protocols"
)

type Decoder struct{}

func New() *Decoder {
	return &Decoder{}
}

func (d *Decoder) Decode(pkt []byte) (string, error) {
	eth, rest, ok := protocols.ParseEthernet(pkt)
	if !ok {
		return "", fmt.Errorf("invalid ethernet frame")
	}

	// EtherType 0x0800 = IPv4
	if eth.EtherType != 0x0800 {
		return "", fmt.Errorf("not IPv4")
	}

	ip, rest, ok := protocols.ParseIPv4(rest)
	if !ok {
		return "", fmt.Errorf("invalid ipv4 packet")
	}

	// Protocol 6 = TCP
	if ip.Protocol != 6 {
		return "", fmt.Errorf("not TCP")
	}

	tcp, ok := protocols.ParseTCP(rest)
	if !ok {
		return "", fmt.Errorf("invalid TCP segment")
	}

	out := fmt.Sprintf(
		"TCP %s:%d → %s:%d",
		ip.SrcIP, tcp.SrcPort,
		ip.DstIP, tcp.DstPort,
	)

	return out, nil
}
