package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "ripple/config"
    "ripple/database"
)

// loadServerAddress reads the server address from the configuration file using database.ReadFile.
func loadServerAddress() error {
    addressBytes, err := database.ReadFile(config.GetDataDir(), "server_address.txt")
    if err != nil {
        return fmt.Errorf("failed to load server address: %w", err)
    }

    config.SetServerAddress(string(addressBytes))
    log.Printf("Loaded server address: %s", config.GetServerAddress()) // Log that the address was loaded
    return nil
}

// setupLogger initializes the logging configuration.
func setupLogger() error {
    // Construct the full path to the log file
    logFilePath := filepath.Join(config.GetDataDir(), "ripple.log")

    logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
    if err != nil {
        return fmt.Errorf("failed to open log file: %w", err)
    }

    log.SetOutput(logFile)
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    return nil
}

// initConfig initializes the configuration by setting up the logger and loading the server address.
func initConfig() error {
    if err := setupLogger(); err != nil {
        return fmt.Errorf("initializing logger: %w", err)
    }
    log.Println("Logger setup completed, initializing configuration...")

    if err := loadServerAddress(); err != nil {
        return fmt.Errorf("initializing configuration by loading server address: %w", err)
    }

    log.Println("Configuration initialized successfully.")
    return nil
}
