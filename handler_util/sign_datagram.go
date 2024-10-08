package handler_util

import (
    "fmt"
    "ripple/auth"
    "ripple/types"
)

// SignDatagram creates a signed datagram by serializing it and adding a signature.
// It requires the session to load the secret key for signature generation.
func SignDatagram(dg *types.Datagram, peerServerAddress string) ([]byte, error) {
    // Serialize the datagram without the signature field
    serializedData, err := types.SerializeDatagram(dg)
    if err != nil {
        return nil, fmt.Errorf("failed to serialize datagram: %w", err)
    }

    // Load the secret key for signature generation
    secretKey, err := auth.LoadServerSecretKeyOut(dg, peerServerAddress)
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
