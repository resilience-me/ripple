package datagram_util

import "ripple/database"

// GetCounter retrieves the client's counter.
func (dh DatagramHelper) GetCounter() (uint32, error) {
    return database.GetCounter(dh.Username)
}

// SetCounter sets the client's counter.
func (dh DatagramHelper) SetCounter() error {
    return database.SetCounter(dh.Username, dh.Counter)
}

// GetCounterIn retrieves the incoming counter for a server connection.
func (dh DatagramHelper) GetCounterIn() (uint32, error) {
    return database.GetCounterIn(dh.Username, dh.PeerServerAddress, dh.PeerUsername)
}

// SetCounterIn sets the incoming counter for a server connection.
func (dh DatagramHelper) SetCounterIn() error {
    return database.SetCounterIn(dh.Username, dh.PeerServerAddress, dh.PeerUsername, dh.Counter)
}

// GetCounterOut retrieves the outgoing counter for a server connection.
func (dh DatagramHelper) GetCounterOut() (uint32, error) {
    return database.GetCounterOut(dh.Username, dh.PeerServerAddress, dh.PeerUsername)
}

// SetCounterOut sets the outgoing counter for a server connection.
func (dh DatagramHelper) SetCounterOut(value uint32) error {
    return database.SetCounterOut(dh.Username, dh.PeerServerAddress, dh.PeerUsername, value)
}
