package db_trustlines

import "ripple/types"

// GetTrustlineOutWithDatagram retrieves the outbound trustline using fields from datagram
func GetTrustlineOutWithDatagram(dg *types.Datagram) (uint32, error) {
	return GetTrustlineOut(dg.Username, dg.PeerServerAddress, dg.PeerUsername)
}

// GetTrustlineInWithDatagram retrieves the inbound trustline using fields from datagram
func GetTrustlineInWithDatagram(dg *types.Datagram) (uint32, error) {
	return GetTrustlineIn(dg.Username, dg.PeerServerAddress, dg.PeerUsername)
}

// SetTrustlineOutWithDatagram sets the outbound trustline amount using fields from datagram
func SetTrustlineOutWithDatagram(dg *types.Datagram, value uint32) error {
	return SetTrustlineOut(dg.Username, dg.PeerServerAddress, dg.PeerUsername, value)
}

// SetTrustlineInWithDatagram sets the inbound trustline amount using fields from datagram
func SetTrustlineInWithDatagram(dg *types.Datagram, value uint32) error {
	return SetTrustlineIn(dg.Username, dg.PeerServerAddress, dg.PeerUsername, value)
}

// GetTrustline retrieves the trustline (either incoming or outgoing) based on the inOrOut parameter.
func GetTrustline(username, peerServerAddress, peerUsername string, inOrOut byte) (uint32, error) {
    if inOrOut == 0 { // Assume 0 means incoming trustline
        return GetTrustlineIn(username, peerServerAddress, peerUsername)
    } else { // Assume 1 means outgoing trustline
        return GetTrustlineOut(username, peerServerAddress, peerUsername)
    }
}

// GetCreditline retrieves the creditline (either incoming or outgoing) based on the inOrOut parameter.
func GetCreditline(username, peerServerAddress, peerUsername string, inOrOut byte) (uint32, error) {
    // if inOrOut == 0 { // Assume 0 means incoming trustline
    //     return GetCreditlineIn(username, peerServerAddress, peerUsername)
    // } else { // Assume 1 means outgoing trustline
    //     return GetCreditlineOut(username, peerServerAddress, peerUsername)
    // }
	return 0, nil
}
