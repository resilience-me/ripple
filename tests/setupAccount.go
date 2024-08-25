package tests

import (
    "fmt"
    "os"
    "ripple/database"
    "ripple/database/db_trustlines"
)

// SetupAccount initializes the account and peer directories with default values.
func SetupAccount(username, peerUsername, peerServerAddress, secretKey string) error {
    // Use GetTrustlineDir to obtain the directory path
    trustlineDir := database.GetTrustlineDir(username, peerServerAddress, peerUsername)

    // Create the directories if they don't exist
    if err := os.MkdirAll(trustlineDir, 0755); err != nil {
        return fmt.Errorf("failed to create directories: %w", err)
    }

    // Initialize the files with default values using the database package
    if err := database.SetCounter(username, 0); err != nil {
        return fmt.Errorf("failed to set counter: %w", err)
    }
    if err := database.SetCounterIn(username, peerServerAddress, peerUsername, 0); err != nil {
        return fmt.Errorf("failed to set counter_in: %w", err)
    }
    if err := database.SetCounterOut(username, peerServerAddress, peerUsername, 0); err != nil {
        return fmt.Errorf("failed to set counter_out: %w", err)
    }
    if err := db_trustlines.SetTrustlineOut(username, peerServerAddress, peerUsername, 0); err != nil {
        return fmt.Errorf("failed to set trustline_out: %w", err)
    }

    // Call database.WriteFile with the account directory and secret key filename
    if err := database.WriteFile(database.GetAccountDir(username), "secretkey.txt", []byte(secretKey)); err != nil {
        return fmt.Errorf("failed to write secret key: %w", err)
    }

    fmt.Println("Account and peer directories initialized successfully with secret key.")
    return nil
}
