package auth

import (
	"fmt"
	"ripple/database"
	"ripple/types"
)

// GetAndIncrementCounterOut retrieves the current counter_out, increments it, and updates the database.
// It returns the new counter value after it has been incremented.
func GetAndIncrementCounterOut(username, peerServerAddress, peerUsername string) (uint32, error) {
    // Retrieve the current value of counter_out from the database.
    counterOut, err := database.GetCounterOut(username, peerServerAddress, peerUsername)
    if err != nil {
        return 0, err  // Return error if unable to fetch the counter.
    }

    // Increment the counter.
    counterOut++  // Increment the counter value by 1

    // Update the database with the new counter value.
    if err := database.SetCounterOut(username, peerServerAddress, peerUsername, counterOut); err != nil {
        return 0, err  // Return error if unable to update the counter.
    }

    // Return the updated counter value.
    return counterOut, nil
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
