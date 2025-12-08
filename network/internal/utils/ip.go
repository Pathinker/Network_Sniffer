package utils

import "net"

func IPv4ToString(b []byte) string {
    return net.IPv4(b[0], b[1], b[2], b[3]).String()
}
