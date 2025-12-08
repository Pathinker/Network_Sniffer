package utils

import "fmt"

func MacToString(b []byte) string {
    return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x",
        b[0], b[1], b[2], b[3], b[4], b[5])
}
