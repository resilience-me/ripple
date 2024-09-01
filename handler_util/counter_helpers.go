package handler_util

// ValidateAndIncrementClientCounter checks and updates the client counter.
func (dh DatagramHelper) ValidateAndIncrementClientCounter() error {
    prevCounter, err := dh.GetCounter() // Retrieve the previous counter value
    if err != nil {
        return fmt.Errorf("error retrieving counter: %v", err)
    }
    if dh.Counter <= prevCounter {
        return fmt.Errorf("replay detected or old datagram: Counter %d is not greater than the last seen counter %d", dh.Counter, prevCounter)
    }
    if err := dh.SetCounter(); err != nil { // Set the new counter
        return fmt.Errorf("failed to set counter: %v", err)
    }
    return nil
}

// ValidateAndIncrementServerCounter checks and updates the server's incoming counter.
func (dh DatagramHelper) ValidateAndIncrementServerCounter() error {
    prevCounter, err := dh.GetCounterIn() // Retrieve the previous incoming counter value
    if err != nil {
        return fmt.Errorf("error retrieving in-counter: %v", err)
    }
    if dh.Counter <= prevCounter {
        return fmt.Errorf("replay detected or old datagram: Counter %d is not greater than the last seen in-counter %d", dh.Counter, prevCounter)
    }
    if err := dh.SetCounterIn(); err != nil { // Set the new counter
        return fmt.Errorf("failed to set in-counter: %v", err)
    }
    return nil
}
