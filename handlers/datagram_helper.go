package handlers

import (
    "ripple/types"
    "ripple/db_trustlines"
    "ripple/commands"
)

// DatagramHelper provides utility methods for working with a Datagram instance
type DatagramHelper struct {
    Datagram *types.Datagram
}

// NewDatagramHelper creates a new DatagramHelper instance
func NewDatagramHelper(datagram *types.Datagram) DatagramHelper {
    return DatagramHelper{Datagram: datagram}
}

// GetTrustlineOut retrieves the outbound trustline using the wrapped Datagram
func (dh DatagramHelper) GetTrustlineOut() (uint32, error) {
    return db_trustlines.GetTrustlineOut(dh.Datagram.Username, dh.Datagram.PeerServerAddress, dh.Datagram.PeerUsername)
}

// GetTrustlineIn retrieves the inbound trustline using the wrapped Datagram
func (dh DatagramHelper) GetTrustlineIn() (uint32, error) {
    return db_trustlines.GetTrustlineIn(dh.Datagram.Username, dh.Datagram.PeerServerAddress, dh.Datagram.PeerUsername)
}

// SetTrustlineOut sets the outbound trustline amount using the wrapped Datagram
func (dh DatagramHelper) SetTrustlineOut(value uint32) error {
    return db_trustlines.SetTrustlineOut(dh.Datagram.Username, dh.Datagram.PeerServerAddress, dh.Datagram.PeerUsername, value)
}

// SetTrustlineIn sets the inbound trustline amount using the wrapped Datagram
func (dh DatagramHelper) SetTrustlineIn(value uint32) error {
    return db_trustlines.SetTrustlineIn(dh.Datagram.Username, dh.Datagram.PeerServerAddress, dh.Datagram.PeerUsername, value)
}

// PrepareDatagram prepares a new datagram using fields from the wrapped Datagram
func (dh DatagramHelper) PrepareDatagram() (*types.Datagram, error) {
    return PrepareDatagramWithoutCommand(dh.Datagram.Username, dh.Datagram.PeerServerAddress, dh.Datagram.PeerUsername)
}

// PrepareAndSendDatagram prepares and sends a new datagram using fields from the wrapped Datagram
func (dh DatagramHelper) PrepareAndSendDatagram(command byte, arguments []byte) error {
    return PrepareAndSendDatagram(command, dh.Datagram.Username, dh.Datagram.PeerServerAddress, dh.Datagram.PeerUsername, arguments)
}
