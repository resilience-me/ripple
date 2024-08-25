package datagram

import (
	"ripple/database"
)

// GetCounter retrieves the counter value using the embedded Datagram to determine the directory.
func (dh DatagramHelper) GetCounter() (uint32, error) {
	return database.GetCounter(dh.Username)
}

// SetCounter sets the counter value using the embedded Datagram.
func (dh DatagramHelper) SetCounter() error {
	return database.SetCounter(dh.Username, dh.Counter)
}

// GetCounterIn retrieves the counter_in value using the embedded Datagram to determine the directory.
func (dh DatagramHelper) GetCounterIn() (uint32, error) {
	return database.GetCounterIn(dh.Username, dh.PeerServerAddress, dh.PeerUsername)
}

// SetCounterIn sets the counter_in value using the embedded Datagram.
func (dh DatagramHelper) SetCounterIn() error {
	return database.SetCounterIn(dh.Username, dh.PeerServerAddress, dh.PeerUsername, dh.Counter)
}

// GetCounterOut retrieves the counter_out value using the embedded Datagram to determine the directory.
func (dh DatagramHelper) GetCounterOut() (uint32, error) {
	return database.GetCounterOut(dh.Username, dh.PeerServerAddress, dh.PeerUsername)
}

// SetCounterOut sets the counter_out value using the embedded Datagram.
func (dh DatagramHelper) SetCounterOut(value uint32) error {
	return database.SetCounterOut(dh.Username, dh.PeerServerAddress, dh.PeerUsername, value)
}
