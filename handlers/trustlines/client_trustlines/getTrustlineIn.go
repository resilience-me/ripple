package client_trustlines

import (
    "log"

    "ripple/comm"
    "ripple/handlers"
    "ripple/types"
)

// GetTrustlineIn handles fetching the inbound trustline information
func GetTrustlineIn(session types.Session) {
    // Instantiate a DatagramHelper using the NewDatagramHelper function
    dh := handlers.NewDatagramHelper(session.Datagram)

    // Fetch the inbound trustline using DatagramHelper
    trustline, err := dh.GetTrustlineIn()
    if err != nil {
        log.Printf("Error reading inbound trustline for user %s: %v", dh.Username, err)
        comm.SendErrorResponse(session.Addr, "Error reading inbound trustline.")
        return
    }

    // Prepare success response
    responseData := types.Uint32ToBytes(trustline)

    // Send the success response back to the client
    if err := comm.SendSuccessResponse(session.Addr, responseData); err != nil {
        log.Printf("Error sending success response to user %s: %v", dh.Username, err)
        return
    }

    log.Printf("Inbound trustline sent successfully to user %s.", dh.Username)
}
