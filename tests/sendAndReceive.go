package tests

import (
    "crypto/rand"
    "fmt"
    "net"
)

// sendAndReceive sends a datagram to the server, waits for an acknowledgment, and returns the server's response.
func sendAndReceive(serverAddress string, data []byte) ([]byte, error) {
    // Generate a random 4-byte identifier
    identifier := make([]byte, 4)
    if _, err := rand.Read(identifier); err != nil {
        return nil, fmt.Errorf("failed to generate identifier: %w", err)
    }

    // Prepend the identifier to the data
    dataWithID := append(identifier, data...)

    // Listen on an ephemeral port
    conn, err := net.ListenUDP("udp", &net.UDPAddr{
        IP:   net.ParseIP("127.0.0.1"),
        Port: 0, // 0 lets the OS choose a free port
    })
    if err != nil {
        return nil, fmt.Errorf("failed to create UDP connection: %w", err)
    }
    defer conn.Close()

    // Send the datagram to the server
    serverUDPAddr, err := net.ResolveUDPAddr("udp", serverAddress)
    if err != nil {
        return nil, fmt.Errorf("failed to resolve server address: %w", err)
    }

    _, err = conn.WriteToUDP(dataWithID, serverUDPAddr)
    if err != nil {
        return nil, fmt.Errorf("failed to send datagram: %w", err)
    }

    fmt.Println("Datagram sent, waiting for acknowledgment...")

    // Wait for the acknowledgment
    buffer := make([]byte, 4)
    n, _, err := conn.ReadFromUDP(buffer)
    if err != nil {
        return nil, fmt.Errorf("failed to receive acknowledgment: %w", err)
    }

    ack := buffer[:n]
    fmt.Printf("Received acknowledgment: %x\n", ack)

    // Wait for the server's response
    buffer = make([]byte, 1024)
    n, _, err = conn.ReadFromUDP(buffer)
    if err != nil {
        return nil, fmt.Errorf("failed to receive server response: %w", err)
    }

    response := buffer[:n]
    fmt.Printf("Received response: %x\n", response)

    return response, nil
}
