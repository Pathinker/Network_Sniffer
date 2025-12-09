package decoder

func (d *Decoder) colorForProto(p string) string {
	switch p {
	case "TCP":
		return "\033[32m" // Verde
	case "UDP":
		return "\033[34m" // Azul
	case "ICMP":
		return "\033[33m" // Amarillo
	case "DNS":
		return "\033[35m" // Magenta
	case "HTTP":
		return "\033[92m" // Verde claro
	case "SSH":
		return "\033[93m" // Amarillo claro
	case "SMTP":
		return "\033[95m" // Rosa
	case "IMAP":
		return "\033[36m" // Cian
	case "POP3":
		return "\033[90m" // Gris
	case "ARP":
		return "\033[96m" // Cian claro
	default:
		return "\033[37m" // Blanco
	}
}

func (d *Decoder) colorizeProto(proto string, s string) string {
	reset := "\033[0m"
	return d.colorForProto(proto) + s + reset
}
