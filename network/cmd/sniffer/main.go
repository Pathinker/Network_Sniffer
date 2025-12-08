package main

import (
    "fmt"
    "network-sniffer/internal/capture"
    "network-sniffer/internal/decoder"
)

func main() {
    cap := capture.NewCapture("eth0")
    dec := decoder.NewDecoder()

    cap.Listen(func(packet []byte) {
        decoded, _ := dec.Decode(packet)
        fmt.Println(decoded)
    })
}
