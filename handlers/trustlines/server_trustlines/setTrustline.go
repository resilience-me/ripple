package server_trustlines

import (
    "log"
    "time"
    "ripple/types"
    "ripple/database/db_trustlines"
)

// SetTrustline handles setting or updating a trustline from another server's perspective.
func SetTrustline(session types.Session) {
    datagram := session.Datagram

    // Retrieve the trustline amount from the Datagram
    trustlineAmount := types.BytesToUint32(datagram.Arguments[:4])

    // Update the trustline and timestamp
    if err := db_trustlines.SetTrustlineInWithDatagram(datagram, trustlineAmount); err != nil {
        log.Printf("Error writing trustline to file for user %s: %v", datagram.Username, err)
        return
    }

    if err := db_trustlines.SetTimestamp(datagram, time.Now().Unix()); err != nil {
        log.Printf("Error writing timestamp to file for user %s: %v", datagram.Username, err)
        return
    }

    log.Printf("Trustline synchronization timestamp updated successfully for user %s.", datagram.Username)
}
