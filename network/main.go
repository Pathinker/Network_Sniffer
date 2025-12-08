package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"network_sniffer/internal/capture"
	"network_sniffer/internal/decoder"
)

func main() {
	fmt.Println("🕸 Simple Go Network Sniffer")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig
		fmt.Println("\nStopping...")
		cancel()
	}()

	cap, err := capture.New()
	if err != nil {
		log.Fatal(err)
	}
	defer cap.Close()

	pkts, err := cap.Start(ctx)
	if err != nil {
		log.Fatal(err)
	}

	dec := decoder.New()

	for {
		select {
		case <-ctx.Done():
			return

		case p, ok := <-pkts:
			if !ok {
				return
			}

			info, err := dec.Decode(p)
			if err == nil {
				fmt.Println(info)
			}
		}
	}
}
