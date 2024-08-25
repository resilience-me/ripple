package handlers

import (
	"ripple/database"
)

// LoadServerSecretKeyOut loads the server's outbound secret key using the embedded Datagram and the provided peer server address.
func (dh DatagramHelper) LoadServerSecretKeyOut(peerServerAddress string) ([]byte, error) {
	return database.LoadPeerSecretKey(dh.PeerUsername, peerServerAddress, dh.Username)
}
