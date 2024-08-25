package config

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
    "time"
    "ripple/database"
)

const (
    Port = 2012
)

// PathFindingTimeout is a global constant that defines the timeout duration for pathfinding operations
const PathFindingTimeout = 5 * time.Minute

// CommitTimeout is a global constant that defines the timeout duration for commits during payment
const CommitTimeout = 10 * time.Minute

var datadir = filepath.Join(os.Getenv("HOME"), "ripple")
var serverAddress string

// GetServerAddress returns the server address as a string
func GetServerAddress() string {
    return serverAddress
}

// GetDataDir returns the datadir as a string
func GetDataDir() string {
    return datadir
}

// loadServerAddress reads the server address from the configuration file.
func loadServerAddress() error {
    // Read the server address from the file
    addressBytes, err := database.ReadFile(datadir, "server_address.txt")
    if err != nil {
        return fmt.Errorf("error loading server address: %w", err)
    }

    // Convert the byte slice to a string
    serverAddress = string(addressBytes)

    // Log that the address was loaded
    log.Printf("Loaded server address: %s", serverAddress)
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
