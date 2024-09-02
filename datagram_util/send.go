package datagram_util

import (
    "fmt"
    "ripple/comm"
)

// SignAndSend signs the datagram and sends it over the network with the standard importance (5 retries).
func (dg *Datagram) SignAndSend(peerServerAddress string) error {
    // Sign the datagram
    serializedData, err := dg.Sign(peerServerAddress)
    if err != nil {
        return fmt.Errorf("failed to create signed datagram: %w", err)
    }

    // Use the default Send function for standard importance
    if err := comm.Send(peerServerAddress, serializedData); err != nil {
        return fmt.Errorf("failed to send datagram: %w", err)
    }

    return nil // Successfully signed and sent
}

// SignAndSendPriority signs the datagram and sends it over the network with priority importance (12 retries).
func (dg *Datagram) SignAndSendPriority(peerServerAddress string) error {
    // Sign the datagram
    serializedData, err := dg.Sign(peerServerAddress)
    if err != nil {
        return fmt.Errorf("failed to create signed datagram: %w", err)
    }

    // Use the SendPriority function for high importance
    if err := comm.SendPriority(peerServerAddress, serializedData); err != nil {
        return fmt.Errorf("failed to send datagram with priority: %w", err)
    }

    return nil // Successfully signed and sent
}
