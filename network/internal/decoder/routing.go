package decoder

var EthertypeMap = map[uint16]string{
    0x0800: "IPv4",
    0x86DD: "IPv6",
    0x0806: "ARP",
    0x8100: "VLAN",
}

var IPv4ProtocolMap = map[uint8]string{
    1:  "ICMP",
    2:  "IGMP",
    6:  "TCP",
    17: "UDP",
}

var IPv6NextHeaderMap = map[uint8]string{
    6:  "TCP",
    17: "UDP",
    58: "ICMPv6",
}
