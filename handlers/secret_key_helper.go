package handlers

import (
	"ripple/database"
)

// LoadClientSecretKey loads the client's secret key using the embedded Datagram.
func (dh DatagramHelper) LoadClientSecretKey() ([]byte, error) {
	return database.LoadSecretKey(dh.Username)
}

// LoadServerSecretKey loads the server's secret key using the embedded Datagram.
func (dh DatagramHelper) LoadServerSecretKey() ([]byte, error) {
	return database.LoadPeerSecretKey(dh.Username, dh.PeerServerAddress, dh.PeerUsername)
}

// LoadServerSecretKeyOut loads the server's outbound secret key using the embedded Datagram and the provided peer server address.
func (dh DatagramHelper) LoadServerSecretKeyOut(peerServerAddress string) ([]byte, error) {
	return database.LoadPeerSecretKey(dh.PeerUsername, peerServerAddress, dh.Username)
}
