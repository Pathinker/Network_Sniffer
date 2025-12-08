package protocols

type IGMP struct {
	Type uint8
	MaxRespTime uint8
}

func ParseIGMP(data []byte) (IGMP, bool) {
	if len(data) < 8 {
		return IGMP{}, false
	}

	return IGMP{
		Type: data[0],
		MaxRespTime: data[1],
	}, true
}
