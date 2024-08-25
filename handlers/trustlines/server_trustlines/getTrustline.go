package server_trustlines

import (
    "log"
    "encoding/binary"

    "ripple/comm"
    "ripple/types"
    "ripple/handlers"
    "ripple/database/db_trustlines"
    "ripple/commands"
)

// GetTrustline handles the request to get the current trustline amount from another server
func GetTrustline(session types.Session) {
    datagram := session.Datagram

    trustline, err := db_trustlines.GetTrustlineOutWithDatagram(session.Datagram)
    if err != nil {
        log.Printf("Error getting trustline for user %s in GetTrustline: %v", session.Datagram.Username, err)
        return
    }

    arguments := types.Uint32ToBytes(trustline)

    // Prepare, sign, and send the datagram using the helper function from the handlers package
    if err := handlers.PrepareAndSendDatagramWithDatagram(datagram, commands.ServerTrustlines_SetTrustline, arguments); err != nil {
        log.Printf("Failed to prepare and send ServerTrustlines_SetTrustline command from %s to peer %s at server %s: %v", datagram.Username, datagram.PeerUsername, datagram.PeerServerAddress, err)
        return
    }

    log.Printf("ServerTrustlines_GetTrustline operation completed successfully for user %s.", datagram.Username)
}
