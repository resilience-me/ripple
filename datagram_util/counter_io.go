package datagram_util

import "ripple/database"

// GetCounter retrieves the client's counter.
func (dg *Datagram) GetCounter() (uint32, error) {
    return database.GetCounter(dg.Username)
}

// SetCounter sets the client's counter.
func (dg *Datagram) SetCounter() error {
    return database.SetCounter(dg.Username, dg.Counter)
}

// GetCounterIn retrieves the incoming counter for a server connection.
func (dg *Datagram) GetCounterIn() (uint32, error) {
    return database.GetCounterIn(dg.Username, dg.PeerServerAddress, dg.PeerUsername)
}

// SetCounterIn sets the incoming counter for a server connection.
func (dg *Datagram) SetCounterIn() error {
    return database.SetCounterIn(dg.Username, dg.PeerServerAddress, dg.PeerUsername, dg.Counter)
}

// GetCounterOut retrieves the outgoing counter for a server connection.
func (dg *Datagram) GetCounterOut() (uint32, error) {
    return database.GetCounterOut(dg.Username, dg.PeerServerAddress, dg.PeerUsername)
}

// SetCounterOut sets the outgoing counter for a server connection.
func (dg *Datagram) SetCounterOut(value uint32) error {
    return database.SetCounterOut(dg.Username, dg.PeerServerAddress, dg.PeerUsername, value)
}
