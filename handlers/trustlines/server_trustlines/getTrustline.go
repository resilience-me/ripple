package server_trustlines

import (
    "log"

    "ripple/types"
    "ripple/handler_util"
    "ripple/commands"
)

// GetTrustline handles the request to get the current trustline amount from another server
func GetTrustline(session types.Session) {
    // Instantiate a DatagramHelper using the NewDatagramHelper function
    dh := handler_util.NewDatagramHelper(session.Datagram)

    // Retrieve the outbound trustline using DatagramHelper
    trustline, err := dh.GetTrustlineOut()
    if err != nil {
        log.Printf("Error getting trustline for user %s in GetTrustline: %v", dh.Username, err)
        return
    }

    arguments := types.Uint32ToBytes(trustline)

    // Prepare, sign, and send the datagram using the DatagramHelper
    if err := dh.PrepareAndSendDatagram(commands.ServerTrustlines_SetTrustline, arguments); err != nil {
        log.Printf("Failed to prepare and send ServerTrustlines_SetTrustline command from %s to peer %s at server %s: %v", dh.Username, dh.PeerUsername, dh.PeerServerAddress, err)
        return
    }

    log.Printf("ServerTrustlines_GetTrustline operation completed successfully for user %s.", dh.Username)
}
