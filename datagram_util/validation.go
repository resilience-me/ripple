package datagram_util

import (
	"errors"
	"fmt"
	"ripple/database"
	"ripple/auth"
)

var (
	// Predefined error for signature verification failure
	ErrSignatureVerificationFailed = errors.New("signature verification failed")
)

// ValidatePeerExists checks for the existence of user and peer directories.
// It returns an error message string (empty if successful) and an error object for detailed information if an error occurs.
func (dg *Datagram) ValidatePeerExists() (string, error) {
	exists, err := database.CheckPeerExists(dg.Username, dg.PeerServerAddress, dg.PeerUsername)
	if err != nil {
		return "Error checking peer existence", fmt.Errorf("error checking peer existence for server '%s' and user '%s': %v", dg.PeerServerAddress, dg.PeerUsername, err)
	} else if !exists {
		return "Peer account does not exist", fmt.Errorf("peer directory does not exist for server '%s' and user '%s'", dg.PeerServerAddress, dg.PeerUsername)
	}

	return "", nil // No error, directories exist
}

// ValidateClientDatagram validates the client datagram and checks the counter.
func (dg *Datagram) ValidateClientDatagram(buf []byte) error {
	secretKey, err := dg.loadClientSecretKey()
	if err != nil {
		return fmt.Errorf("loading client secret key failed: %w", err)
	}

	if !auth.verifySignature(buf, secretKey) {
		return ErrSignatureVerificationFailed
	}

	// Validate the counter
	if err := dg.ValidateAndIncrementClientCounter(); err != nil {
		return fmt.Errorf("counter validation failed: %w", err)
	}

	return nil
}

// ValidateServerDatagram validates the server datagram and checks the counter.
func (dg *Datagram) ValidateServerDatagram(buf []byte) error {
	secretKey, err := dg.loadServerSecretKey()
	if err != nil {
		return fmt.Errorf("loading server secret key failed: %w", err)
	}

	if !auth.verifySignature(buf, secretKey) {
		return ErrSignatureVerificationFailed
	}

	// Validate the counter
	if err := dg.ValidateAndIncrementServerCounter(); err != nil {
		return fmt.Errorf("counter validation failed: %w", err)
	}

	return nil
}

// ValidateDatagram validates a datagram based on whether it's for a client or server session.
func (dg *Datagram) ValidateDatagram(buf []byte) error {
	if dg.Command&0x80 == 0 { // Client session if MSB is 0
		return dg.ValidateClientDatagram(buf)
	} else { // Server session if MSB is 1
		return dg.ValidateServerDatagram(buf)
	}
}
