package client_trustlines

import (
    "encoding/binary"
    "log"

    "ripple/commands"
    "ripple/comm"
    "ripple/database/db_trustlines"
    "ripple/handlers"
    "ripple/types"
)

// SyncTrustlineIn handles the client request to sync the inbound trustline from the peer server.
func SyncTrustlineIn(session types.Session) {
    datagram := session.Datagram

    // Prepare the datagram
    dgOut, err := handlers.PrepareDatagramResponse(datagram)
    if err != nil {
        log.Printf("Error preparing datagram for user %s: %v", datagram.Username, err)
        comm.SendErrorResponse(session.Addr, "Error preparing datagram.")
        return
    }

    dgOut.Command = commands.ServerTrustlines_GetTrustline

    // Send the GetTrustline command to the peer server
    if err := comm.SignAndSendDatagram(dgOut, datagram.PeerServerAddress); err != nil {
        log.Printf("Failed to send GetTrustline command for user %s to peer %s: %v", datagram.Username, datagram.PeerUsername, err)
        comm.SendErrorResponse(session.Addr, "Failed to send GetTrustline command.")
        return
    }

    // Send success response to the client
    if err := comm.SendSuccessResponse(session.Addr, []byte("Trustline sync request sent successfully.")); err != nil {
        log.Printf("Failed to send success response to user %s: %v", datagram.Username, err)
        return
    }

    log.Printf("GetTrustline command sent successfully for user %s to peer %s.", datagram.Username, datagram.PeerUsername)
}
