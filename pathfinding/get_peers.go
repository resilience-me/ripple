package pathfinding

import (
	"ripple/database/db_pathfinding"
)

// GetPeers retrieves a list of all peer accounts for a given username.
func GetPeers(username string) ([]PeerAccount, error) {
	var peers []PeerAccount

	// Get the list of server addresses
	serverAddresses, err := db_pathfinding.GetServerDirs(username)
	if err != nil {
		return nil, err
	}

	// For each server address, get the list of peers
	for _, serverAddress := range serverAddresses {
		peerUsernames, err := db_pathfinding.GetPeerDirs(username, serverAddress)
		if err != nil {
			return nil, err
		}

		// Create PeerAccount structs and add them to the result
		for _, peerUsername := range peerUsernames {
			peer := PeerAccount{
				Username:      peerUsername,
				ServerAddress: serverAddress,
			}
			peers = append(peers, peer)
		}
	}

	return peers, nil
}
