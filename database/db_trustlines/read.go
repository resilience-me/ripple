package db_trustlines

import (
	"ripple/types"
	"ripple/database"
)

// GetTrustlineOut retrieves the outbound trustline
func GetTrustlineOut(username, peerServerAddress, peerUsername string) (uint32, error) {
	trustlineDir := database.GetTrustlineDir(username, peerServerAddress, peerUsername)
	return database.GetUint32FromFile(trustlineDir, "trustline_out.txt")
}

// GetTrustlineIn retrieves the inbound trustline
func GetTrustlineIn(username, peerServerAddress, peerUsername string) (uint32, error) {
	trustlineDir := database.GetTrustlineDir(username, peerServerAddress, peerUsername)
	return database.GetUint32FromFile(trustlineDir, "trustline_in.txt")
}
