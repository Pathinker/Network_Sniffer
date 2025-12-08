package protocols

type DHCP struct {
	Op uint8
}

func ParseDHCP(data []byte) (DHCP, bool) {
	if len(data) < 240 {
		return DHCP{}, false
	}

	return DHCP{
		Op: data[0],
	}, true
}
