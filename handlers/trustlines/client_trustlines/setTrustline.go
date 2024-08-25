package client_trustlines

import (
    "encoding/binary"
    "log"

    "ripple/comm"
    "ripple/handlers"
    "ripple/types"
)

// SetTrustline updates the trustline based on the given session.
func SetTrustline(session types.Session) {
    // Instantiate a DatagramHelper using the NewDatagramHelper function
    dh := handlers.NewDatagramHelper(session.Datagram)

    // Retrieve the trustline amount from the Datagram
    trustlineAmount := binary.BigEndian.Uint32(dh.Arguments[:4])

    // Write the new trustline amount using the DatagramHelper
    if err := dh.SetTrustlineOut(trustlineAmount); err != nil {
        log.Printf("Error writing trustline to file for user %s: %v", dh.Username, err)
        comm.SendErrorResponse(session.Addr, "Failed to write trustline.")
        return
    }

    // Log success
    log.Printf("Trustline updated successfully for user %s.", dh.Username)

    // Send success response
    if err := comm.SendSuccessResponse(session.Addr, []byte("Trustline updated successfully.")); err != nil {
        log.Printf("Failed to send success response to user %s: %v", dh.Username, err)
        return
    }

    log.Printf("Sent success response to client for user %s.", dh.Username)
}
