package client_payments

import (
    "log"
    "ripple/comm"
    "ripple/pathfinding"
    "ripple/types"
)

// StartPayment handles the command to start a payment by initiating the pathfinding process.
func StartPayment(session types.Session) {
    username := session.Datagram.Username

    // Inline check for the existence of a registered payment
    account := pathfinding.GetPathManager().Find(username)
    if account == nil || account.Payment == nil {
        if err := comm.SendErrorResponse(session.Addr, "Payment not registered or missing payment details."); err != nil {
            log.Printf("Failed to send error response for user %s: %v", username, err)
        }
        return
    }

    // Extract payment details
    paymentDetails := account.Payment

    // Initiate pathfinding using StartFindPath
    err := pathfinding.StartFindPath(paymentDetails)
    if err != nil {
        log.Printf("Pathfinding failed for user %s: %v", username, err)
        if err := comm.SendErrorResponse(session.Addr, "Failed to start payment due to pathfinding error."); err != nil {
            log.Printf("Failed to send error response for user %s: %v", username, err)
        }
        return
    }

    // Respond with success
    if err := comm.SendSuccessResponse(session.Addr, []byte("Payment started successfully.")); err != nil {
        log.Printf("Failed to send success response for user %s: %v", username, err)
    }

    log.Printf("Payment started successfully for user %s.", username)
}
