package datagram_util

import (
    "fmt"
    "ripple/database"
)

// ValidateAndIncrementClientCounter checks and updates the client counter.
func (dg *Datagram) ValidateAndIncrementClientCounter() error {
    prevCounter, err := dg.GetCounter() // Retrieve the previous counter value
    if err != nil {
        return fmt.Errorf("error retrieving counter: %v", err)
    }
    if dg.Counter <= prevCounter {
        return fmt.Errorf("replay detected or old datagram: Counter %d is not greater than the last seen counter %d", dg.Counter, prevCounter)
    }
    if err := dg.SetCounter(); err != nil { // Set the new counter
        return fmt.Errorf("failed to set counter: %v", err)
    }
    return nil
}

// ValidateAndIncrementServerCounter checks and updates the server's incoming counter.
func (dg *Datagram) ValidateAndIncrementServerCounter() error {
    prevCounter, err := dg.GetCounterIn() // Retrieve the previous incoming counter value
    if err != nil {
        return fmt.Errorf("error retrieving in-counter: %v", err)
    }
    if dg.Counter <= prevCounter {
        return fmt.Errorf("replay detected or old datagram: Counter %d is not greater than the last seen in-counter %d", dg.Counter, prevCounter)
    }
    if err := dg.SetCounterIn(); err != nil { // Set the new counter
        return fmt.Errorf("failed to set in-counter: %v", err)
    }
    return nil
}
