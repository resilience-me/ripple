package db_pathfinding

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"ripple/config"
)

// GetServerDirs retrieves a list of server address directories for a given username.
func GetServerDirs(username string) ([]string, error) {
	baseDir := filepath.Join(config.GetDataDir(), "accounts", username, "peers")

	// Read all server address directories in the peers directory
	serverDirs, err := ioutil.ReadDir(baseDir)
	if err != nil {
		return nil, fmt.Errorf("unable to read directory %s: %v", baseDir, err)
	}

	// Extract and return directory names
	var serverAddresses []string
	for _, serverDir := range serverDirs {
		if serverDir.IsDir() {
			serverAddresses = append(serverAddresses, serverDir.Name())
		}
	}

	return serverAddresses, nil
}

// GetPeerDirs retrieves a list of peer directories under a specific server address for a given username.
func GetPeerDirs(username, serverAddress string) ([]string, error) {
	serverPath := filepath.Join(config.GetDataDir(), "accounts", username, "peers", serverAddress)

	// Read all peer directories under the current server address
	peerDirs, err := ioutil.ReadDir(serverPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read directory %s: %v", serverPath, err)
	}

	// Extract and return directory names
	var peerUsernames []string
	for _, peerDir := range peerDirs {
		if peerDir.IsDir() {
			peerUsernames = append(peerUsernames, peerDir.Name())
		}
	}

	return peerUsernames, nil
}
