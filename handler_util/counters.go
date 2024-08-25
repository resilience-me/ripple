package handler_util

import (
	"fmt"
	"ripple/database"
)

// GetAndIncrementCounterOut retrieves the current counter_out, increments it, and updates the database.
// It returns the counter value before it was incremented.
func GetAndIncrementCounterOut(username, peerServerAddress, peerUsername string) (uint32, error) {
    // Retrieve the current value of counter_out from the database.
    counterOut, err := database.GetCounterOut(username, peerServerAddress, peerUsername)
    if err != nil {
        return 0, err  // Return error if unable to fetch the counter.
    }

    // Increment the counter and update it in the database within the same function call.
    if err := database.SetCounterOut(username, peerServerAddress, peerUsername, counterOut+1); err != nil {
        return 0, err  // Return error if unable to update the counter.
    }

    // Return the original counter value that was fetched.
    return counterOut, nil
}
