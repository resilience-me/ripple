package db_trustlines

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
