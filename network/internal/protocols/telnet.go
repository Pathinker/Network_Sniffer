package protocols

type Telnet struct{}

func ParseTelnet(data []byte) (Telnet, bool) {
	if len(data) > 0 {
		return Telnet{}, true
	}
	return Telnet{}, false
}
