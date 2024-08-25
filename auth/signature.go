package auth

import (
    "crypto/sha256"
    "bytes"
    "ripple/database"
    "ripple/types"
)

func loadClientSecretKey(dg *types.Datagram) ([]byte, error) {
    return database.LoadSecretKey(dg.Username)
}

func loadServerSecretKey(dg *types.Datagram) ([]byte, error) {
    return database.LoadPeerSecretKey(dg.Username, dg.PeerServerAddress, dg.PeerUsername)
}

func loadServerSecretKeyOut(dg *types.Datagram, peerServerAddress string) ([]byte, error) {
    return database.LoadPeerSecretKey(dg.PeerUsername, peerServerAddress, dg.Username)
}

// GenerateSignature generates a SHA-256 hash for the given datagram using the provided key.
func GenerateSignature(data []byte, secretKey []byte) []byte {
    // Remove the signature from the datagram
    dataWithoutSignature := data[:len(data)-32]

    // Concatenate data and secret
    preimage := append(dataWithoutSignature, secretKey...)

    // Compute the SHA-256 hash
    hash := sha256.Sum256(preimage)

    // Return the hash as a byte slice
    return hash[:]
}

// VerifySignature checks the integrity of the received buffer
func VerifySignature(data []byte, key []byte) bool {
    // The signature is the last 32 bytes of the buffer
    signature := data[len(data)-32:]

    // Generate the expected signature using the GenerateSignature method
    expectedSignature := generateSignature(data, key)

    // Compare the computed hash with the signature directly using bytes.Equal
    return bytes.Equal(signature, expectedSignature)
}
