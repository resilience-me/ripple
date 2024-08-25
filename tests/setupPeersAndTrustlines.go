package tests

import (
    "fmt"
    "ripple/database/db_trustlines"
)

// setupPeersAndTrustlines sets up peer relationships and initializes trustlines between a sender and a receiver.
func setupPeersAndTrustlines(sender, receiver, senderServerAddress, receiverServerAddress, sharedSecretKey string, trustlineAmount uint32) error {
    // Set up peer relationship for the sender
    if err := setupPeer(sender, receiver, receiverServerAddress, sharedSecretKey); err != nil {
        return fmt.Errorf("failed to set up peer relationship for %s: %w", sender, err)
    }
    if err := db_trustlines.SetTrustlineIn(sender, receiverServerAddress, receiver, trustlineAmount); err != nil {
        return fmt.Errorf("failed to set trustline_in for %s: %w", sender, err)
    }

    // Set up peer relationship for the receiver
    if err := setupPeer(receiver, sender, senderServerAddress, sharedSecretKey); err != nil {
        return fmt.Errorf("failed to set up peer relationship for %s: %w", receiver, err)
    }
    if err := db_trustlines.SetTrustlineOut(receiver, senderServerAddress, sender, trustlineAmount); err != nil {
        return fmt.Errorf("failed to set trustline_out for %s: %w", receiver, err)
    }

    fmt.Printf("Peer relationship and trustlines between %s (sender) and %s (receiver) initialized successfully.\n", sender, receiver)
    return nil
}
