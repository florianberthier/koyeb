package utils

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/exp/rand"
)

func GetRandomPort(minPort, maxPort int) (int, error) {
	if minPort > maxPort {
		return 0, fmt.Errorf("invalid port range: minPort > maxPort")
	}

	rand.Seed(uint64(time.Now().UnixNano()))

	for attempts := 0; attempts < 10; attempts++ {
		port := rand.Intn(maxPort-minPort+1) + minPort
		if IsPortAvailable(port) {
			return port, nil
		}
	}

	return 0, fmt.Errorf("could not find an available port after multiple attempts")
}

func IsPortAvailable(port int) bool {
	addr := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		return false
	}
	_ = ln.Close()
	return true
}
