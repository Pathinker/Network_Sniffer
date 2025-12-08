package capture

import (
    "net"
)

type Capture struct {
    Iface string
}

func NewCapture(iface string) *Capture {
    return &Capture{Iface: iface}
}

func (c *Capture) Listen(callback func([]byte)) error {
    conn, err := net.ListenPacket("ip4:ethernet", c.Iface)
    if err != nil {
        return err
    }
    defer conn.Close()

    buf := make([]byte, 65535)

    for {
        n, _, err := conn.ReadFrom(buf)
        if err != nil {
            return err
        }
        callback(buf[:n])
    }
}
