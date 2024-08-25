package db_trustlines

import (
	"ripple/database"
)

// SetTrustlineOut sets the outbound trustline amount.
func SetTrustlineOut(username, peerServerAddress, peerUsername string, value uint32) error {
	trustlineDir := database.GetTrustlineDir(username, peerServerAddress, peerUsername)
	return database.WriteUint32ToFile(trustlineDir, "trustline_out.txt", value)
}

// SetTrustlineOut sets the inbound trustline amount.
func SetTrustlineIn(username, peerServerAddress, peerUsername string, value uint32) error {
	trustlineDir := database.GetTrustlineDir(username, peerServerAddress, peerUsername)
	return database.WriteUint32ToFile(trustlineDir, "trustline_in.txt", value)
}

// SetTimestamp sets the sync timestamp.
func SetTimestamp(username, peerServerAddress, peerUsername string, timestamp int64) error {
	trustlineDir := database.GetTrustlineDir(username, peerServerAddress, peerUsername)
	return database.WriteTimeToFile(trustlineDir, "timestamp.txt", timestamp)
}
