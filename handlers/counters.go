package datagram

import (
	"ripple/database"
)

// GetCounter retrieves the counter value using the embedded Datagram to determine the directory.
func (dh DatagramHelper) GetCounter() (uint32, error) {
	return database.GetCounter(dh.Username)
}

// SetCounter sets the counter value using the embedded Datagram.
func (dh DatagramHelper) SetCounter() error {
	return database.SetCounter(dh.Username, dh.Counter)
}

// GetCounterIn retrieves the counter_in value using the embedded Datagram to determine the directory.
func (dh DatagramHelper) GetCounterIn() (uint32, error) {
	return database.GetCounterIn(dh.Username, dh.PeerServerAddress, dh.PeerUsername)
}

// SetCounterIn sets the counter_in value using the embedded Datagram.
func (dh DatagramHelper) SetCounterIn() error {
	return database.SetCounterIn(dh.Username, dh.PeerServerAddress, dh.PeerUsername, dh.Counter)
}

// GetCounterOut retrieves the counter_out value using the embedded Datagram to determine the directory.
func (dh DatagramHelper) GetCounterOut() (uint32, error) {
	return database.GetCounterOut(dh.Username, dh.PeerServerAddress, dh.PeerUsername)
}

// SetCounterOut sets the counter_out value using the embedded Datagram.
func (dh DatagramHelper) SetCounterOut(value uint32) error {
	return database.SetCounterOut(dh.Username, dh.PeerServerAddress, dh.PeerUsername, value)
}

// ValidateAndIncrementClientCounter checks if the datagram's counter is valid for client connections and increments it if valid.
func (dh DatagramHelper) ValidateAndIncrementClientCounter() error {
	prevCounter, err := dh.GetCounter()
	if err != nil {
		return fmt.Errorf("error retrieving counter: %v", err)
	}
	if dh.Counter <= prevCounter {
		return fmt.Errorf("replay detected or old datagram: Counter %d is not greater than the last seen counter %d", dh.Counter, prevCounter)
	}
	if err := dh.SetCounter(); err != nil {
		return fmt.Errorf("failed to set counter: %v", err)
	}
	return nil
}

// ValidateAndIncrementServerCounter checks if the datagram's counter is valid for server connections and increments it if valid.
func (dh DatagramHelper) ValidateAndIncrementServerCounter() error {
	prevCounter, err := dh.GetCounterIn()
	if err != nil {
		return fmt.Errorf("error retrieving in-counter: %v", err)
	}
	if dh.Counter <= prevCounter {
		return fmt.Errorf("replay detected or old datagram: Counter %d is not greater than the last seen in-counter %d", dh.Counter, prevCounter)
	}
	if err := dh.SetCounterIn(); err != nil {
		return fmt.Errorf("failed to set in-counter: %v", err)
	}
	return nil
}
