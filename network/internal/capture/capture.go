package capture

import (
	"context"
	"fmt"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

type Capture struct {
	handle *pcap.Handle
}

func New() (*Capture, error) {
	iface := "eth0" // change if needed

	handle, err := pcap.OpenLive(iface, 1600, true, pcap.BlockForever)
	if err != nil {
		return nil, err
	}

	return &Capture{
		handle: handle,
	}, nil
}

func (c *Capture) Start(ctx context.Context) (<-chan []byte, error) {
	packetSource := gopacket.NewPacketSource(c.handle, c.handle.LinkType())
	out := make(chan []byte)

	go func() {
		defer close(out)

		packets := packetSource.Packets()

		for {
			select {
			case <-ctx.Done():
				return

			case packet, ok := <-packets:
				if !ok {
					return
				}

				if packet == nil {
					continue
				}

				out <- packet.Data()
			}
		}
	}()

	return out, nil
}

func (c *Capture) Close() {
	fmt.Println("Capture closed")
	c.handle.Close()
}
