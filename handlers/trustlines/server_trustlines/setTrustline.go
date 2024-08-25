package server_trustlines

import (
    "log"
    "time"
    "ripple/types"
    "ripple/handlers"
    "ripple/database/db_trustlines"
)

// SetTrustline handles setting or updating a trustline from another server's perspective.
func SetTrustline(session types.Session) {
    // Instantiate a DatagramHelper using the NewDatagramHelper function
    dh := handlers.NewDatagramHelper(session.Datagram)

    // Retrieve the trustline amount from the Datagram
    trustlineAmount := types.BytesToUint32(dh.Arguments[:4])

    // Update the trustline and timestamp using DatagramHelper
    if err := dh.SetTrustlineIn(trustlineAmount); err != nil {
        log.Printf("Error writing trustline to file for user %s: %v", dh.Username, err)
        return
    }

    if err := db_trustlines.SetTimestamp(dh.Datagram, time.Now().Unix()); err != nil {
        log.Printf("Error writing timestamp to file for user %s: %v", dh.Username, err)
        return
    }

    log.Printf("Trustline synchronization timestamp updated successfully for user %s.", dh.Username)
}
