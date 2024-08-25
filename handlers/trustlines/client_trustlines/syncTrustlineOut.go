package client_trustlines

import (
    "log"

    "ripple/commands"
    "ripple/comm"
    "ripple/handlers"
    "ripple/types"
)

// SyncTrustlineOut handles the client request to sync the outbound trustline to the peer server.
func SyncTrustlineOut(session types.Session) {
    // Instantiate a DatagramHelper using the NewDatagramHelper function
    dh := handlers.NewDatagramHelper(session.Datagram)

    // Read the trustline value using DatagramHelper
    trustline, err := dh.GetTrustlineOut()
    if err != nil {
        log.Printf("Error getting trustline for user %s in SyncTrustlineOut: %v", dh.Username, err)
        comm.SendErrorResponse(session.Addr, "Failed to retrieve trustline.")
        return
    }

    arguments := types.Uint32ToBytes(trustline)

    // Prepare, sign, and send the datagram using the DatagramHelper
    if err := dh.PrepareAndSendDatagram(commands.ServerTrustlines_SetTrustline, arguments); err != nil {
        log.Printf("Failed to prepare and send SetTrustline command from %s to peer %s at server %s: %v", dh.Username, dh.PeerUsername, dh.PeerServerAddress, err)
        comm.SendErrorResponse(session.Addr, "Failed to send ServerTrustlines_SetTrustline command.")
        return
    }

    // Send success response to the client
    if err := comm.SendSuccessResponse(session.Addr, []byte("Outbound trustline sync request processed successfully.")); err != nil {
        log.Printf("Failed to send success response in SyncTrustlineOut for user %s: %v", dh.Username, err)
        return
    }

    log.Printf("SyncTrustline command processed successfully for user %s to peer %s.", dh.Username, dh.PeerUsername)
}
