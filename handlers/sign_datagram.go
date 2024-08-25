package handlers

import (
    "fmt"
    "ripple/types"
)

func loadServerSecretKeyOut(dg *types.Datagram, peerServerAddress string) ([]byte, error) {
    return database.LoadPeerSecretKey(dg.PeerUsername, peerServerAddress, dg.Username)
}

// generateSignature generates a SHA-256 hash for the given datagram using the provided key.
func generateSignature(datagram []byte, secretKey []byte) []byte {

    datagramWithoutSignatureField := datagram[:len(datagram)-32]

    // Concatenate data and secret
    preimage := append(datagramWithoutSignatureField, secretKey...)

    // Compute the SHA-256 hash
    hash := sha256.Sum256(preimage)

    // Return the hash as a byte slice
    return hash[:]
}

// SignDatagram creates a signed datagram by serializing it and adding a signature.
// It requires the session to load the secret key for signature generation.
func SignDatagram(dg *types.Datagram, peerServerAddress string) ([]byte, error) {
    // Serialize the datagram without the signature field
    serializedData, err := types.SerializeDatagram(dg)
    if err != nil {
        return nil, fmt.Errorf("failed to serialize datagram: %w", err)
    }

    // Load the secret key for signature generation
    secretKey, err := loadServerSecretKeyOut(dg, peerServerAddress)
    if err != nil {
        return nil, fmt.Errorf("failed to load server secret key: %w", err)
    }

    // Generate signature for the serialized data
    signature := generateSignature(serializedData, secretKey)

    // Update the datagram's signature field with the generated signature
    copy(dg.Signature[:], []byte(signature)) // Ensure we copy the signature into the byte array

    // Return the serialized data including the signature
    return serializedData, nil
}
