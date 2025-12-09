package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"network_sniffer/internal/capture"
	"network_sniffer/internal/decoder"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
)

type PacketInfo struct {
	Timestamp time.Time
	Summary   string
	Raw       []byte
}

var history []PacketInfo

func main() {
	for {
		printMenu()

		reader := bufio.NewReader(os.Stdin)
		opt, _ := reader.ReadString('\n')

		switch opt[0] {
		case '1':
			runCapture()
		case '2':
			showHistory()
		case '3':
			fmt.Println(Yellow + "Saliendo..." + Reset)
			return
		default:
			fmt.Println(Red + "Opción inválida" + Reset)
		}
	}
}

func printMenu() {
	fmt.Println(Cyan + Bold + "\n===============================")
	fmt.Println("   Simple Go Network Sniffer")
	fmt.Println("===============================" + Reset)

	fmt.Println(Green + "1)" + Reset + " Iniciar captura de paquetes")
	fmt.Println(Green + "2)" + Reset + " Ver historial de paquetes capturados")
	fmt.Println(Green + "3)" + Reset + " Salir")

	fmt.Print(Blue + "\nSelecciona una opción: " + Reset)
}

func runCapture() {
	fmt.Println(Green + "\nIniciando captura. Presiona Ctrl+C para detener..." + Reset)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sig
		fmt.Println(Yellow + "\nDeteniendo captura..." + Reset)
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
			fmt.Println(Yellow + "Captura detenida." + Reset)
			return

		case pkt, ok := <-pkts:
			if !ok {
				fmt.Println(Red + "Canal cerrado." + Reset)
				return
			}

			info, err := dec.Decode(pkt)
			if err == nil {
				fmt.Println(info)

				history = append(history, PacketInfo{
					Timestamp: time.Now(),
					Summary:   info,
					Raw:       pkt,
				})
			}
		}
	}
}

func showHistory() {
	if len(history) == 0 {
		fmt.Println(Yellow + "\nNo hay paquetes capturados aún." + Reset)
		return
	}

	fmt.Println(Bold + "\nPaquetes almacenados:" + Reset)

	for i, p := range history {
		fmt.Printf("%s[%d]%s %s - %s\n",
			Cyan, i, Reset,
			p.Timestamp.Format("15:04:05"),
			p.Summary,
		)
	}

	fmt.Print(Blue + "\nSelecciona un índice para ver detalles (o Enter para salir): " + Reset)

	var index int
	_, err := fmt.Scanln(&index)
	if err != nil {
		fmt.Println(Yellow + "Regresando al menú..." + Reset)
		return
	}

	if index < 0 || index >= len(history) {
		fmt.Println(Red + "Índice inválido." + Reset)
		return
	}

	fmt.Println(Bold + "\nDetalles del paquete:" + Reset)
	fmt.Printf("Hora: %s\n", history[index].Timestamp)
	fmt.Println("Resumen:", history[index].Summary)

	fmt.Println(Bold + "\nHexdump:" + Reset)
	fmt.Println(formatHex(history[index].Raw))
}

func formatHex(data []byte) string {
	out := ""
	for i := 0; i < len(data); i += 16 {
		line := fmt.Sprintf("%04x  ", i)
		for j := 0; j < 16 && i+j < len(data); j++ {
			line += fmt.Sprintf("%02x ", data[i+j])
		}
		out += line + "\n"
	}
	return out
}
