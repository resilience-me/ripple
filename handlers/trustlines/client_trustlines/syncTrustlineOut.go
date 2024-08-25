package client_trustlines

import (
    "log"

    "ripple/commands"
    "ripple/comm"
    "ripple/database/db_trustlines"
    "ripple/handlers"
    "ripple/types"
)

// SyncTrustlineOut handles the client request to sync the outbound trustline to the peer server.
func SyncTrustlineOut(session types.Session) {
    datagram := session.Datagram

    // Read the trustline value
    trustline, err := db_trustlines.GetTrustlineOutWithDatagram(datagram)
    if err != nil {
        log.Printf("Error getting trustline for user %s in SyncTrustlineOut: %v", datagram.Username, err)
        comm.SendErrorResponse(session.Addr, "Failed to retrieve trustline.")
        return
    }

    arguments := types.Uint32ToBytes(trustline)
    
    // Prepare, sign, and send the datagram using the helper function from the handlers package
    if err := handlers.PrepareAndSendDatagramWithDatagram(datagram, commands.ServerTrustlines_SetTrustline, arguments); err != nil {
        log.Printf("Failed to prepare and send SetTrustline command from %s to peer %s at server %s: %v", datagram.Username, datagram.PeerUsername, datagram.PeerServerAddress, err)
        comm.SendErrorResponse(session.Addr, "Failed to send ServerTrustlines_SetTrustline command.")
        return
    }

    // Send success response to the client
    if err := comm.SendSuccessResponse(session.Addr, []byte("Outbound trustline sync request processed successfully.")); err != nil {
        log.Printf("Failed to send success response in SyncTrustlineOut for user %s: %v", datagram.Username, err)
        return
    }

    log.Printf("SyncTrustline command processed successfully for user %s to peer %s.", datagram.Username, datagram.PeerUsername)
}
