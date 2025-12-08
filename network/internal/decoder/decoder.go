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
		return "", fmt.Errorf("invalid ethernet")
	}

	// VLAN support
	if eth.EtherType == protocols.EtherTypeVLAN {
		vlan, rest2, ok := protocols.ParseVLAN(rest)
		if !ok {
			return "", fmt.Errorf("bad vlan")
		}
		eth.EtherType = vlan.EtherType
		rest = rest2
	}

	// -------- IPv4 --------
	if eth.EtherType == protocols.EtherTypeIPv4 {
		return d.decodeIPv4(rest)
	}

	// -------- IPv6 --------
	if eth.EtherType == protocols.EtherTypeIPv6 {
		return d.decodeIPv6(rest)
	}

	// -------- ARP --------
	if eth.EtherType == protocols.EtherTypeARP {
		arp, ok := protocols.ParseARP(rest)
		if ok {
			return fmt.Sprintf(
				"ARP %s (%s) → %s (%s)",
				arp.SenderIP, arp.SenderMAC,
				arp.TargetIP, arp.TargetMAC,
			), nil
		}
	}

	return "", fmt.Errorf("unknown ethertype")
}

func (d *Decoder) decodeIPv4(data []byte) (string, error) {
	ip, rest, ok := protocols.ParseIPv4(data)
	if !ok {
		return "", fmt.Errorf("invalid ipv4")
	}

	return d.decodeL4(ip.Protocol, rest, ip.SrcIP, ip.DstIP)
}

func (d *Decoder) decodeIPv6(data []byte) (string, error) {
	ip, rest, ok := protocols.ParseIPv6(data)
	if !ok {
		return "", fmt.Errorf("invalid ipv6")
	}

	return d.decodeL4(ip.NextHdr, rest, ip.SrcIP, ip.DstIP)
}

func (d *Decoder) decodeL4(proto uint8, payload []byte, srcIP, dstIP string) (string, error) {
	switch proto {

	case protocols.ProtoTCP:
		tcp, ok := protocols.ParseTCP(payload)
		if !ok {
			return "", fmt.Errorf("bad tcp")
		}
		return d.prettyAppDetect("TCP", srcIP, dstIP, tcp.SrcPort, tcp.DstPort, payload[20:])

	case protocols.ProtoUDP:
		udp, ok := protocols.ParseUDP(payload)
		if !ok {
			return "", fmt.Errorf("bad udp")
		}
		return d.prettyAppDetect("UDP", srcIP, dstIP, udp.SrcPort, udp.DstPort, payload[8:])

	case protocols.ProtoICMP:
		icmp, ok := protocols.ParseICMP(payload)
		if !ok {
			return "", fmt.Errorf("bad icmp")
		}
		return d.colorize(fmt.Sprintf("ICMP %s → %s type=%d code=%d",
			srcIP, dstIP, icmp.Type, icmp.Code)), nil
	}

	return "", fmt.Errorf("unknown L4")
}
