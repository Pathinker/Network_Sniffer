package protocols

type Routing struct {
	Type uint8
}

func ParseRouting(data []byte) (Routing, bool) {
	if len(data) < 1 {
		return Routing{}, false
	}

	return Routing{
		Type: data[0],
	}, true
}
