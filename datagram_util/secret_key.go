package datagram_util

import (
    "ripple/database"
)

// LoadServerSecretKeyOut loads the server's outgoing secret key for the peer server address.
func (dg *Datagram) LoadServerSecretKeyOut(peerServerAddress string) ([]byte, error) {
    return database.LoadPeerSecretKey(dg.PeerUsername, peerServerAddress, dg.Username)
}

// LoadClientSecretKey loads the client's secret key.
func (dg *Datagram) LoadClientSecretKey() ([]byte, error) {
    return database.LoadSecretKey(dg.Username)
}

// LoadServerSecretKey loads the server's secret key for communication with the client.
func (dg *Datagram) LoadServerSecretKey() ([]byte, error) {
    return database.LoadPeerSecretKey(dg.Username, dg.PeerServerAddress, dg.PeerUsername)
}
