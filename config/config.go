package config

import (
    "os"
    "path/filepath"
    "time"
)

const (
    Port             = 2012
)

const (
    PathFindingTimeout = 5 * time.Minute
    CommitTimeout      = 10 * time.Minute
)

var (
    dataDir       = filepath.Join(os.Getenv("HOME"), "ripple")
    serverAddress string
)

// GetDataDir returns the datadir as a string.
func GetDataDir() string {
    return dataDir
}

// GetServerAddress returns the server address as a string.
func GetServerAddress() string {
    return serverAddress
}

// SetServerAddress sets the server address.
func SetServerAddress(address string) {
    serverAddress = address
}
