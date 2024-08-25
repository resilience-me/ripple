package client_trustlines

import (
    "encoding/binary"
    "log"

    "ripple/commands"
    "ripple/comm"
    "ripple/handlers"
    "ripple/types"
)

// SyncTrustlineIn handles the client request to sync the inbound trustline from the peer server.
func SyncTrustlineIn(session types.Session) {
    datagram := session.Datagram

    // Prepare, sign, and send the datagram using the helper function from the handlers package
    if err := handlers.PrepareAndSendDatagram(commands.ServerTrustlines_GetTrustline, datagram.Username, datagram.PeerServerAddress, datagram.PeerUsername, nil); err != nil {
        log.Printf("Failed to prepare and send GetTrustline command from %s to peer %s at server %s: %v", datagram.Username, datagram.PeerUsername, datagram.PeerServerAddress, err)
        return
    }

    // Send success response to the client
    if err := comm.SendSuccessResponse(session.Addr, []byte("Trustline sync request sent successfully.")); err != nil {
        log.Printf("Failed to send success response to user %s: %v", datagram.Username, err)
        return
    }

    log.Printf("GetTrustline command sent successfully for user %s to peer %s.", datagram.Username, datagram.PeerUsername)
}
