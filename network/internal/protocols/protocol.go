package protocols

// IP Protocol numbers (IANA)
const (
	ProtoICMP = 1
	ProtoTCP  = 6
	ProtoUDP  = 17
	ProtoIPv6 = 41
)

// EtherTypes
const (
	EtherTypeIPv4 = 0x0800
	EtherTypeARP  = 0x0806
	EtherTypeVLAN = 0x8100
	EtherTypeIPv6 = 0x86DD
)

// Generic decoded packet (optional, for later)
type PacketInfo struct {
	L2 interface{}
	L3 interface{}
	L4 interface{}
}
