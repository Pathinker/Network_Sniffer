package protocols

type ICMPv6 struct {
	Type uint8
	Code uint8
}

func ParseICMPv6(data []byte) (ICMPv6, bool) {
	if len(data) < 4 {
		return ICMPv6{}, false
	}

	return ICMPv6{
		Type: data[0],
		Code: data[1],
	}, true
}
