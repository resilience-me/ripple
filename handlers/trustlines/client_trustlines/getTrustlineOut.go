package client_trustlines

import (
    "log"

    "ripple/comm"
    "ripple/handlers"
    "ripple/types"
)

// GetTrustlineOut handles fetching the outbound trustline information
func GetTrustlineOut(session types.Session) {
    // Instantiate a DatagramHelper using the NewDatagramHelper function
    dh := handlers.NewDatagramHelper(session.Datagram)

    // Fetch the outbound trustline using DatagramHelper
    trustline, err := dh.GetTrustlineOut()
    if err != nil {
        log.Printf("Error reading outbound trustline for user %s: %v", dh.Username, err)
        comm.SendErrorResponse(session.Addr, "Error reading outbound trustline.")
        return
    }

    // Prepare success response using the main utility function
    responseData := types.Uint32ToBytes(trustline)

    // Send the success response back to the client
    if err := comm.SendSuccessResponse(session.Addr, responseData); err != nil {
        log.Printf("Error sending success response to user %s: %v", dh.Username, err)
        return
    }

    log.Printf("Outbound trustline sent successfully to user %s.", dh.Username) // Log successful operation
}
