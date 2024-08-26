package config

import (
    "fmt"
    "log"
    "os"
    "io/ioutil"
    "path/filepath"
    "time"
    "strings"
)

// loadServerAddress reads the server address from the configuration file.
func loadServerAddress() error {
    addressPath := filepath.Join(datadir, "server_address.txt")
    addressBytes, err := ioutil.ReadFile(addressPath)
    if err != nil {
        return fmt.Errorf("error loading server address from %s: %w", addressPath, err)
    }
    
    // Remove trailing newlines that might be added by text editors
    serverAddress = strings.TrimSpace(string(addressBytes))
    
    log.Printf("Loaded server address: %s", serverAddress) // Log that the address was loaded
    return nil
}

// setupLogger initializes the logging configuration.
func setupLogger() error {
    // Construct the full path to the log file
    logFilePath := filepath.Join(datadir, "ripple.log")

    logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
    if err != nil {
        return fmt.Errorf("failed to open log file: %w", err)
    }

    log.SetOutput(logFile)
    log.SetFlags(log.LstdFlags | log.Lshortfile)
    return nil
}

// InitConfig initializes the configuration by setting up the logger and loading the server address.
func InitConfig() error {
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
