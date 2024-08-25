package tests

import (
    "fmt"
    "os"
    "ripple/database"
)

// SetupAccount initializes the account directory with counters and a secret key.
func SetupAccount(username, secretKey string) error {
    accountDir := database.GetAccountDir(username)

    // Create the account directory if it doesn't exist
    if err := os.MkdirAll(accountDir, 0755); err != nil {
        return fmt.Errorf("failed to create account directory: %w", err)
    }

    // Initialize the counters with default values
    if err := database.SetCounter(username, 0); err != nil {
        return fmt.Errorf("failed to set counter: %w", err)
    }

    // Set the account secret key
    if err := database.WriteFile(accountDir, "secretkey.txt", []byte(secretKey)); err != nil {
        return fmt.Errorf("failed to write secret key: %w", err)
    }

    fmt.Printf("Account %s initialized successfully with secret key.\n", username)
    return nil
}
