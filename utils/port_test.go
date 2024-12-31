package utils

import (
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRandomPort_Success(t *testing.T) {
	minPort := 3001
	maxPort := 4000

	port, err := GetRandomPort(minPort, maxPort)

	require.NoError(t, err, "unexpected error when getting random port")
	assert.GreaterOrEqual(t, port, minPort, "port should be greater than or equal to minPort")
	assert.LessOrEqual(t, port, maxPort, "port should be less than or equal to maxPort")
	assert.True(t, IsPortAvailable(port), "port should be available")
}

func TestGetRandomPort_InvalidRange(t *testing.T) {
	minPort := 4000
	maxPort := 3001

	port, err := GetRandomPort(minPort, maxPort)

	assert.Error(t, err, "expected error for invalid port range")
	assert.Equal(t, 0, port, "port should be zero on error")
	assert.Contains(t, err.Error(), "invalid port range: minPort > maxPort")
}

func TestGetRandomPort_UnavailablePorts(t *testing.T) {
	minPort := 3001
	maxPort := 3002

	ln1, err := net.Listen("tcp", fmt.Sprintf(":%d", minPort))
	require.NoError(t, err, "unexpected error when listening on port %d", minPort)
	defer ln1.Close()

	ln2, err := net.Listen("tcp", fmt.Sprintf(":%d", maxPort))
	require.NoError(t, err, "unexpected error when listening on port %d", maxPort)
	defer ln2.Close()

	port, err := GetRandomPort(minPort, maxPort)

	assert.Error(t, err, "expected error when no ports are available")
	assert.Equal(t, 0, port, "port should be zero when no ports are available")
	assert.Contains(t, err.Error(), "could not find an available port after multiple attempts")
}

func TestIsPortAvailable(t *testing.T) {
	port := 4000

	assert.True(t, IsPortAvailable(port), "port should be available")

	ln, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	require.NoError(t, err, "unexpected error when listening on port %d", port)
	defer ln.Close()

	assert.False(t, IsPortAvailable(port), "port should not be available")
}
