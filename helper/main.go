package helper

import (
	"fmt"
	"os"
)

func NetworkAddress(port int) string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return "0.0.0.0:" + envPort
	}
	return "0.0.0.0:" + fmt.Sprintf("%d", port)
}
