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
            out := fmt.Sprintf("[DNS] %s:%d → %s:%d id=%d",
                srcIP, srcPort, dstIP, dstPort, msg.ID)
            return d.colorizeProto("DNS", out), nil
        }
    }

    // -------- HTTP --------
    if http, ok := protocols.ParseHTTP(payload); ok {
        out := fmt.Sprintf("[HTTP %s] %s:%d → %s:%d",
            http.Method, srcIP, srcPort, dstIP, dstPort)
        return d.colorizeProto("HTTP", out), nil
    }

    // -------- SSH --------
    if _, ok := protocols.ParseSSH(payload); ok {
        out := fmt.Sprintf("[SSH] %s:%d → %s:%d",
            srcIP, srcPort, dstIP, dstPort)
        return d.colorizeProto("SSH", out), nil
    }

    // -------- SMTP --------
    if _, ok := protocols.ParseSMTP(payload); ok {
        out := fmt.Sprintf("[SMTP] %s:%d → %s:%d",
            srcIP, srcPort, dstIP, dstPort)
        return d.colorizeProto("SMTP", out), nil
    }

    // -------- IMAP --------
    if _, ok := protocols.ParseIMAP(payload); ok {
        out := fmt.Sprintf("[IMAP] %s:%d → %s:%d",
            srcIP, srcPort, dstIP, dstPort)
        return d.colorizeProto("IMAP", out), nil
    }

    // -------- POP3 --------
    if _, ok := protocols.ParsePOP3(payload); ok {
        out := fmt.Sprintf("[POP3] %s:%d → %s:%d",
            srcIP, srcPort, dstIP, dstPort)
        return d.colorizeProto("POP3", out), nil
    }

    // -------- fallback --------
    out := fmt.Sprintf("[%s] %s:%d → %s:%d",
        l4, srcIP, srcPort, dstIP, dstPort)

    return d.colorizeProto(l4, out), nil
}
