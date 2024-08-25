package tests

import (
    "fmt"
    "ripple/database/db_trustlines"
)

// setupTrustline sets up the incoming and outgoing trustlines between two accounts.
func setupTrustline(senderUsername, receiverUsername, serverAddress string, trustlineAmount uint32) error {
    // Set the incoming trustline for the sender from the receiver
    if err := db_trustlines.SetTrustlineIn(senderUsername, serverAddress, receiverUsername, trustlineAmount); err != nil {
        return fmt.Errorf("failed to set trustlineIn for sender: %w", err)
    }

    // Set the outgoing trustline for the receiver to the sender
    if err := db_trustlines.SetTrustlineOut(receiverUsername, serverAddress, senderUsername, trustlineAmount); err != nil {
        return fmt.Errorf("failed to set trustlineOut for receiver: %w", err)
    }

    return nil
}
