package database

// GetCounter retrieves the counter value using the username to determine the directory.
func GetCounter(username string) (uint32, error) {
	accountDir := GetAccountDir(username)
	return GetUint32FromFile(accountDir, "counter.txt")
}

// SetCounter sets the counter value using the username to determine the directory.
func SetCounter(username string, value uint32) error {
	accountDir := GetAccountDir(username)
	return WriteUint32ToFile(accountDir, "counter.txt", value)
}

// GetCounterIn retrieves the counter_in value using the user and peer identifiers to determine the directory.
func GetCounterIn(username, peerServerAddress, peerUsername string) (uint32, error) {
	peerDir := GetPeerDir(username, peerServerAddress, peerUsername)
	return GetUint32FromFile(peerDir, "counter_in.txt")
}

// SetCounterIn sets the counter_in value using the user and peer identifiers to determine the directory.
func SetCounterIn(username, peerServerAddress, peerUsername string, value uint32) error {
	peerDir := GetPeerDir(username, peerServerAddress, peerUsername)
	return WriteUint32ToFile(peerDir, "counter_in.txt", value)
}

// GetCounterOut retrieves the counter_out value using the user and peer identifiers to determine the directory.
func GetCounterOut(username, peerServerAddress, peerUsername string) (uint32, error) {
	peerDir := GetPeerDir(username, peerServerAddress, peerUsername)
	return GetUint32FromFile(peerDir, "counter_out.txt")
}

// SetCounterOut sets the counter_out value using the user and peer identifiers to determine the directory.
func SetCounterOut(username, peerServerAddress, peerUsername string, value uint32) error {
	peerDir := GetPeerDir(username, peerServerAddress, peerUsername)
	return WriteUint32ToFile(peerDir, "counter_out.txt", value)
}
