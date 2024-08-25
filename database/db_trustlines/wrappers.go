package db_trustlines

import "ripple/types"

// Method on Datagram to get outbound trustline
func (dg *types.Datagram) GetTrustlineOut() (uint32, error) {
	return GetTrustlineOut(dg.Username, dg.PeerServerAddress, dg.PeerUsername)
}

// Method on Datagram to get inbound trustline
func (dg *types.Datagram) GetTrustlineIn() (uint32, error) {
	return GetTrustlineIn(dg.Username, dg.PeerServerAddress, dg.PeerUsername)
}

// Method on Datagram to set outbound trustline
func (dg *types.Datagram) SetTrustlineOut(value uint32) error {
	return SetTrustlineOut(dg.Username, dg.PeerServerAddress, dg.PeerUsername, value)
}

// Method on Datagram to set inbound trustline
func (dg *types.Datagram) SetTrustlineIn(value uint32) error {
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
