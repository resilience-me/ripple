package handlers

import (
    "fmt"
    "ripple/types"
)

// generateSignature generates a SHA-256 hash for the given data using the provided key.
func generateSignature(data []byte, secretKey []byte) []byte {
    // Concatenate data and secret
    preimage := append(data, secretKey...)

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

    dataWithoutSignature := serializedData[:len(serializedData)-32]

    // Generate signature for the serialized data
    signature := generateSignature(dataWithoutSignature, secretKey)

    // Update the datagram's signature field with the generated signature
    copy(dg.Signature[:], []byte(signature)) // Ensure we copy the signature into the byte array

    // Return the serialized data including the signature
    return serializedData, nil
}
