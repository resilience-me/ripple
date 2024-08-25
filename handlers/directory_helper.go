package datagram

import (
	"ripple/database"
)

// CheckPeerExists checks if the peer directory exists using the embedded Datagram.
func (dh DatagramHelper) CheckPeerExists() (bool, error) {
	peerDir := database.GetPeerDir(dh.Username, dh.PeerServerAddress, dh.PeerUsername)
	return database.CheckDirExists(peerDir)
}
