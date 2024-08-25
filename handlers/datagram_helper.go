package handlers

import (
    "time"
    "ripple/types"
    "ripple/db_trustlines"
    "ripple/commands"
)

// DatagramHelper provides utility methods for working with a Datagram instance
type DatagramHelper struct {
    *types.Datagram // Embedding Datagram directly
}

// NewDatagramHelper creates a new DatagramHelper instance
func NewDatagramHelper(datagram *types.Datagram) DatagramHelper {
    return DatagramHelper{Datagram: datagram}
}

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

// GetTrustlineOut retrieves the outbound trustline using the embedded Datagram
func (dh DatagramHelper) GetTrustlineOut() (uint32, error) {
    return db_trustlines.GetTrustlineOut(dh.Username, dh.PeerServerAddress, dh.PeerUsername)
}

// GetTrustlineIn retrieves the inbound trustline using the embedded Datagram
func (dh DatagramHelper) GetTrustlineIn() (uint32, error) {
    return db_trustlines.GetTrustlineIn(dh.Username, dh.PeerServerAddress, dh.PeerUsername)
}

// SetTrustlineOut sets the outbound trustline amount using the embedded Datagram
func (dh DatagramHelper) SetTrustlineOut(value uint32) error {
    return db_trustlines.SetTrustlineOut(dh.Username, dh.PeerServerAddress, dh.PeerUsername, value)
}

// SetTrustlineIn sets the inbound trustline amount using the embedded Datagram
func (dh DatagramHelper) SetTrustlineIn(value uint32) error {
    return db_trustlines.SetTrustlineIn(dh.Username, dh.PeerServerAddress, dh.PeerUsername, value)
}

// SetTimestamp sets the current Unix timestamp for the Datagram
func (dh DatagramHelper) SetTimestamp() error {
    return db_trustlines.SetTimestamp(dh.Username, dh.PeerServerAddress, dh.PeerUsername, time.Now().Unix())
}

// PrepareDatagramWithoutCommand prepares a new datagram using fields from the embedded Datagram
func (dh DatagramHelper) PrepareDatagramWithoutCommand() (*types.Datagram, error) {
    return PrepareDatagramWithoutCommand(dh.Username, dh.PeerServerAddress, dh.PeerUsername)
}

// PrepareAndSendDatagram prepares and sends a new datagram using fields from the embedded Datagram
func (dh DatagramHelper) PrepareAndSendDatagram(command byte, arguments []byte) error {
    return PrepareAndSendDatagram(command, dh.Username, dh.PeerServerAddress, dh.PeerUsername, arguments)
}
