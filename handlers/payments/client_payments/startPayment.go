package client_payments

import (
    "log"
    "ripple/comm"
    "ripple/handlers/payments/payment_operations"
    "ripple/pathfinding"
    "ripple/types"
)

// StartPayment handles the command to start a payment by initiating the pathfinding process.
func StartPayment(session types.Session) {
    username := session.Datagram.Username

    // Extract payment details
    identifier, amount, inOrOut, err := GetPaymentDetails(username)
    if err != nil {
        if err := comm.SendErrorResponse(session.Addr, "Payment not registered or missing payment details."); err != nil {
            log.Printf("Failed to send error response for user %s: %v", username, err)
        }
        return
    }

    // Initiate pathfinding using StartFindPath from payment_operations
    payment_operations.StartFindPath(username, paymentDetails.Identifier, path.Amount, paymentDetails.InOrOut)

    log.Printf("Payment started for user %s.", username)

    // Respond with success
    if err := comm.SendSuccessResponse(session.Addr, []byte("Payment started successfully.")); err != nil {
        log.Printf("Failed to send success response for user %s: %v", username, err)
        return
    }
    log.Printf("Sent success response to client for user %s.", username)
}
