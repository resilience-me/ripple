package client_trustlines

import (
    "log"

    "ripple/commands"
    "ripple/comm"
    "ripple/handler_util"
    "ripple/types"
)

// SyncTrustlineIn handles the client request to sync the inbound trustline from the peer server.
func SyncTrustlineIn(session types.Session) {
    // Instantiate a DatagramHelper using the NewDatagramHelper function
    dh := handler_util.NewDatagramHelper(session.Datagram)

    // Prepare, sign, and send the datagram using the DatagramHelper
    if err := dh.PrepareAndSendDatagram(commands.ServerTrustlines_GetTrustline, nil); err != nil {
        log.Printf("Failed to prepare and send GetTrustline command from %s to peer %s at server %s: %v", dh.Username, dh.PeerUsername, dh.PeerServerAddress, err)
        comm.SendErrorResponse(session.Addr, "Failed to send GetTrustline command.")
        return
    }

    // Send success response to the client
    if err := comm.SendSuccessResponse(session.Addr, []byte("Trustline sync request sent successfully.")); err != nil {
        log.Printf("Failed to send success response to user %s: %v", dh.Username, err)
        return
    }

    log.Printf("GetTrustline command sent successfully for user %s to peer %s.", dh.Username, dh.PeerUsername)
}
