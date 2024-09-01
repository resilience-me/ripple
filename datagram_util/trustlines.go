package datagram_util

import (
    "ripple/db_trustlines" // Import the db_trustlines package
    "time"
)

// GetTrustlineOut retrieves the outbound trustline using the embedded Datagram.
func (dg *Datagram) GetTrustlineOut() (uint32, error) {
    return db_trustlines.GetTrustlineOut(dg.Username, dg.PeerServerAddress, dg.PeerUsername)
}

// GetTrustlineIn retrieves the inbound trustline using the embedded Datagram.
func (dg *Datagram) GetTrustlineIn() (uint32, error) {
    return db_trustlines.GetTrustlineIn(dg.Username, dg.PeerServerAddress, dg.PeerUsername)
}

// SetTrustlineOut sets the outbound trustline amount using the embedded Datagram.
func (dg *Datagram) SetTrustlineOut(value uint32) error {
    return db_trustlines.SetTrustlineOut(dg.Username, dg.PeerServerAddress, dg.PeerUsername, value)
}

// SetTrustlineIn sets the inbound trustline amount using the embedded Datagram.
func (dg *Datagram) SetTrustlineIn(value uint32) error {
    return db_trustlines.SetTrustlineIn(dg.Username, dg.PeerServerAddress, dg.PeerUsername, value)
}

// SetTimestamp sets the current Unix timestamp for the Datagram.
func (dg *Datagram) SetTimestamp() error {
    now := time.Now().Unix()
    return db_trustlines.SetTimestamp(dg.Username, dg.PeerServerAddress, dg.PeerUsername, now)
}
