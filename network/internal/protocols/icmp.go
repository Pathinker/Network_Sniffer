package protocols

type ICMP struct {
	Type uint8
	Code uint8
}

func ParseICMP(data []byte) (ICMP, bool) {
	if len(data) < 4 {
		return ICMP{}, false
	}

	return ICMP{
		Type: data[0],
		Code: data[1],
	}, true
}
