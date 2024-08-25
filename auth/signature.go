package auth

import (
    "crypto/sha256"
    "bytes"
    "ripple/database"
    "ripple/types"
)

func LoadServerSecretKeyOut(dg *types.Datagram, peerServerAddress string) ([]byte, error) {
    return database.LoadPeerSecretKey(dg.PeerUsername, peerServerAddress, dg.Username)
}

func loadClientSecretKey(dg *types.Datagram) ([]byte, error) {
    return database.LoadSecretKey(dg.Username)
}

func loadServerSecretKey(dg *types.Datagram) ([]byte, error) {
    return database.LoadPeerSecretKey(dg.Username, dg.PeerServerAddress, dg.PeerUsername)
}

// GenerateSignature generates a SHA-256 hash for the given datagram using the provided key.
func GenerateSignature(data []byte, secretKey []byte) []byte {
    // Create a temporary slice to hold the concatenated data (buf without the last 32 bytes) and the secret key.
    preimage := make([]byte, len(data)) // Allocate slice with the length of buf

    // Copy the data (excluding the last 32 bytes) into the preimage.
    copy(preimage, buf[:len(data)-32])

    // Append the secret key to the preimage.
    copy(preimage[len(data)-32:], secretKey)

    // Compute the SHA-256 hash of the preimage.
    hash := sha256.Sum256(preimage)

    // Return the hash as a byte slice.
    return hash[:]
}


// verifySignature checks the integrity of the received buffer
func verifySignature(buf []byte, key []byte) bool {
    // The signature is the last 32 bytes of the buffer
    signature := make([]byte, 32)
    copy(signature, buf[len(buf)-32:])

    // Generate the expected signature using the GenerateSignature method
    expectedSignature := GenerateSignature(buf, key)

    // Compare the signature with the expected signature directly using bytes.Equal
    return bytes.Equal(signature, expectedSignature)
}
