package handlers

import (
	"fmt"
	"ripple/auth"
)

// LoadClientSecretKey, ValidateAndIncrementClientCounter, etc. methods...

// SignDatagram creates a signed datagram by serializing it and adding a signature.
func (dh DatagramHelper) SignDatagram(peerServerAddress string) ([]byte, error) {
	// Serialize the datagram without the signature field
	serializedData, err := types.SerializeDatagram(dh.Datagram)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize datagram: %w", err)
	}

	// Load the secret key for signature generation using DatagramHelper's method
	secretKey, err := dh.LoadServerSecretKeyOut(peerServerAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to load server secret key: %w", err)
	}

	// Generate signature for the serialized data
	signature := auth.GenerateSignature(serializedData[:357], secretKey)

	// Update the datagram's signature field with the generated signature
	copy(dh.Signature[:], []byte(signature)) // Ensure we copy the signature into the byte array

	// Return the serialized data including the signature
	return serializedData, nil
}
