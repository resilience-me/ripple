package handler_util

import (
    "time"
    "ripple/types"
    "ripple/database/db_trustlines"
)

// DatagramHelper provides utility methods for working with a Datagram instance
type DatagramHelper struct {
    *types.Datagram // Embedding Datagram directly
}

// NewDatagramHelper creates a new DatagramHelper instance
func NewDatagramHelper(datagram *types.Datagram) DatagramHelper {
    return DatagramHelper{Datagram: datagram}
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
    now := time.Now().Unix()
    return db_trustlines.SetTimestamp(dh.Username, dh.PeerServerAddress, dh.PeerUsername, now)
}

// PrepareDatagramWithoutCommand prepares a new datagram using fields from the embedded Datagram
func (dh DatagramHelper) PrepareDatagramWithoutCommand() (*types.Datagram, error) {
    return PrepareDatagramWithoutCommand(dh.Username, dh.PeerServerAddress, dh.PeerUsername)
}

// PrepareAndSendDatagram prepares and sends a new datagram using fields from the embedded Datagram
func (dh DatagramHelper) PrepareAndSendDatagram(command byte, arguments []byte) error {
    return PrepareAndSendDatagram(command, dh.Username, dh.PeerServerAddress, dh.PeerUsername, arguments)
}
