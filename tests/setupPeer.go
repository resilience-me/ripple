package tests

import (
    "fmt"
    "os"
    "ripple/database"
    "ripple/database/db_trustlines"
)

// SetupPeer initializes the peer relationship between two accounts, including trustlines and peer-specific secret keys.
func SetupPeer(username, peerUsername, peerServerAddress, secretKey string) error {
    trustlineDir := database.GetTrustlineDir(username, peerServerAddress, peerUsername)

    // Create the directories if they don't exist
    if err := os.MkdirAll(trustlineDir, 0755); err != nil {
        return fmt.Errorf("failed to create trustline directory: %w", err)
    }

    // Initialize the counters for peer communication
    if err := database.SetCounterIn(username, peerServerAddress, peerUsername, 0); err != nil {
        return fmt.Errorf("failed to set counter_in: %w", err)
    }

    if err := database.SetCounterOut(username, peerServerAddress, peerUsername, 0); err != nil {
        return fmt.Errorf("failed to set counter_out: %w", err)
    }

    // Initialize the trustlines
    if err := db_trustlines.SetTrustlineOut(username, peerServerAddress, peerUsername, 0); err != nil {
        return fmt.Errorf("failed to set trustline_out: %w", err)
    }

    if err := db_trustlines.SetTrustlineIn(username, peerServerAddress, peerUsername, 0); err != nil {
        return fmt.Errorf("failed to set trustline_in: %w", err)
    }

    // Set the server secret key for communication between peers
    peerDir := database.GetPeerDir(username, peerServerAddress, peerUsername)

    if err := database.WriteFile(peerDir, "secretkey.txt", []byte(secretKey)); err != nil {
        return fmt.Errorf("failed to write peer secret key: %w", err)
    }

    fmt.Printf("Peer relationship between %s and %s initialized successfully.\n", username, peerUsername)
    return nil
}
