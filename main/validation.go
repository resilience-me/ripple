package main

import (
	"errors"
	"fmt"
	"ripple/types"
	"ripple/database"
)

var (
	// Predefined error for signature verification failure
	ErrSignatureVerificationFailed = errors.New("signature verification failed")
)

// validatePeerExists checks for the existence of user and peer directories
// It returns an error message string (empty if successful) and an error object for detailed information if an error occurs.
func validatePeerExists(dg *types.Datagram) (string, error) {
	exists, err := database.CheckPeerExists(dg.Username, dg.PeerServerAddress, dg.PeerUsername)
	if err != nil {
		return "Error checking peer existence", fmt.Errorf("error checking peer existence for server '%s' and user '%s': %v", dg.PeerServerAddress, dg.PeerUsername, err)
	} else if !exists {
		return "Peer account does not exist", fmt.Errorf("peer directory does not exist for server '%s' and user '%s'", dg.PeerServerAddress, dg.PeerUsername)
	}

	return "", nil // No error, directories exist
}

// validateAndIncrementClientCounter checks if the datagram's counter is valid by comparing it to the last known counter for client connections.
// If valid, it sets the counter to the value in the datagram to prevent replay attacks.
func validateAndIncrementClientCounter(datagram *types.Datagram) error {
	prevCounter, err := database.GetCounter(datagram.Username)
	if err != nil {
		return fmt.Errorf("error retrieving counter: %v", err)
	}
	if datagram.Counter <= prevCounter {
		return fmt.Errorf("replay detected or old datagram: Counter %d is not greater than the last seen counter %d", datagram.Counter, prevCounter)
	}
	if err := database.SetCounter(datagram.Username, datagram.Counter); err != nil {
		return fmt.Errorf("failed to set counter: %v", err)
	}
	return nil
}

// validateAndIncrementServerCounter checks if the datagram's counter is valid by comparing it to the last known counter for server connections.
// If valid, it sets the counter to the value in the datagram to prevent replay attacks.
func validateAndIncrementServerCounter(datagram *types.Datagram) error {
	prevCounter, err := database.GetCounterIn(datagram.Username, datagram.PeerServerAddress, datagram.PeerUsername)
	if err != nil {
		return fmt.Errorf("error retrieving in-counter: %v", err)
	}
	if datagram.Counter <= prevCounter {
		return fmt.Errorf("replay detected or old datagram: Counter %d is not greater than the last seen in-counter %d", datagram.Counter, prevCounter)
	}
	if err := database.SetCounterIn(datagram.Username, datagram.PeerServerAddress, datagram.PeerUsername, datagram.Counter); err != nil {
		return fmt.Errorf("failed to set in-counter: %v", err)
	}
	return nil
}

// validateClientDatagram validates the client datagram and checks the counter
func validateClientDatagram(buf []byte, dg *types.Datagram) error {
	secretKey, err := loadClientSecretKey(dg)
	if err != nil {
		return fmt.Errorf("loading client secret key failed: %w", err)
	}

	if !verifySignature(buf, secretKey) {
		return ErrSignatureVerificationFailed
	}

	// Validate the counter
	if err := validateAndIncrementClientCounter(dg); err != nil {
		return fmt.Errorf("counter validation failed: %w", err)
	}

	return nil
}

// validateServerDatagram validates the server datagram and checks the counter
func validateServerDatagram(buf []byte, dg *types.Datagram) error {
	secretKey, err := loadServerSecretKey(dg)
	if err != nil {
		return fmt.Errorf("loading server secret key failed: %w", err)
	}

	if !verifySignature(buf, secretKey) {
		return ErrSignatureVerificationFailed
	}

	// Validate the counter
	if err := validateAndIncrementServerCounter(dg); err != nil {
		return fmt.Errorf("counter validation failed: %w", err)
	}

	return nil
}

// ValidateDatagram validates a datagram based on whether it's for a client or server session.
func ValidateDatagram(buf []byte, dg *types.Datagram) error {
	if dg.Command&0x80 == 0 { // Client session if MSB is 0
		return validateClientDatagram(buf, dg)
	} else { // Server session if MSB is 1
		return validateServerDatagram(buf, dg)
	}
}
