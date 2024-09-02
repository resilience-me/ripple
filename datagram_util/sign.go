package datagram_util

import (
    "fmt"
    "ripple/auth"
    "ripple/types"
)

// Sign creates a signed datagram by serializing it and adding a signature.
func (dg *Datagram) Sign(peerServerAddress string) ([]byte, error) {
    // Serialize the datagram
    serializedData, err := dg.Serialize()
    if err != nil {
        return nil, fmt.Errorf("failed to serialize datagram: %w", err)
    }

    // Load the secret key
    secretKey, err := dg.LoadServerSecretKeyOut(peerServerAddress)
    if err != nil {
        return nil, fmt.Errorf("failed to load server secret key: %w", err)
    }

    // Generate signature for the serialized data
    signature := auth.GenerateSignature(serializedData, secretKey)

    // Update the serialized datagram's signature field with the generated signature
    copy(serializedData[357:389], signature) // Ensure we copy the signature into the byte array

    // Return the serialized data including the signature
    return serializedData, nil
}
