package comm

import (
	"fmt"
	"net"
	"ripple/config"
	"ripple/udpr"
)

// sendWithAddress sends data to a specified UDP address with retry logic.
// It handles the creation and closure of the UDP connection internally.
func sendWithAddress(addr *net.UDPAddr, data []byte, maxRetries int) error {
	// Create a UDP connection with an ephemeral local port
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		return fmt.Errorf("failed to create UDP connection: %w", err)
	}
	defer conn.Close()

	// Use the udpr.SendWithRetry to send the data
	if err := udpr.SendWithRetry(conn, data, maxRetries); err != nil {
		return fmt.Errorf("error sending data: %w", err)
	}

	return nil
}

// sendWithResolvedAddress resolves the address, creates a new UDP connection, and sends data with retries.
func sendWithResolvedAddress(address string, data []byte, maxRetries int) error {
	// Resolve the destination address to a UDP address
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("%s:%d", address, config.Port))
	if err != nil {
		return fmt.Errorf("failed to resolve address '%s:%d': %w", address, config.Port, err)
	}
	// Call sendWithAddress function with the resolved address
	return sendWithAddress(addr, data, maxRetries)
}

// Wrapper for udpr.SendAck
func SendAck(conn *net.UDPConn, addr *net.UDPAddr, idBytes []byte) error {
	return udpr.SendAck(conn, addr, idBytes)
}
