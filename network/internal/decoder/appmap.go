package decoder

import (
	"fmt"
	"network_sniffer/internal/protocols"
)

func (d *Decoder) prettyAppDetect(
	l4 string,
	srcIP, dstIP string,
	srcPort, dstPort uint16,
	payload []byte,
) (string, error) {

	// -------- DNS --------
	if srcPort == 53 || dstPort == 53 {
		if msg, ok := protocols.ParseDNS(payload); ok {
			return d.colorize(fmt.Sprintf(
				"[DNS] %s:%d → %s:%d id=%d",
				srcIP, srcPort, dstIP, dstPort, msg.ID,
			)), nil
		}
	}

	// -------- HTTP --------
	if http, ok := protocols.ParseHTTP(payload); ok {
		return d.colorize(fmt.Sprintf(
			"[HTTP %s] %s:%d → %s:%d",
			http.Method, srcIP, srcPort, dstIP, dstPort,
		)), nil
	}

	// -------- SSH --------
	if _, ok := protocols.ParseSSH(payload); ok {
		return d.colorize(fmt.Sprintf(
			"[SSH] %s:%d → %s:%d",
			srcIP, srcPort, dstIP, dstPort,
		)), nil
	}

	// -------- SMTP --------
	if _, ok := protocols.ParseSMTP(payload); ok {
		return d.colorize(fmt.Sprintf(
			"[SMTP] %s:%d → %s:%d",
			srcIP, srcPort, dstIP, dstPort,
		)), nil
	}

	// -------- IMAP --------
	if _, ok := protocols.ParseIMAP(payload); ok {
		return d.colorize(fmt.Sprintf(
			"[IMAP] %s:%d → %s:%d",
			srcIP, srcPort, dstIP, dstPort,
		)), nil
	}

	// -------- POP3 --------
	if _, ok := protocols.ParsePOP3(payload); ok {
		return d.colorize(fmt.Sprintf(
			"[POP3] %s:%d → %s:%d",
			srcIP, srcPort, dstIP, dstPort,
		)), nil
	}

	// -------- fallback --------
	return d.colorize(fmt.Sprintf(
		"[%s] %s:%d → %s:%d",
		l4, srcIP, srcPort, dstIP, dstPort,
	)), nil
}
