package main

import (
    "log"
    "ripple/commands"
    "ripple/datagram_util"
)

// dispatchClientHandler processes a client session by executing the appropriate handler.
func dispatchClientHandler(session *datagram_util.Session) {
    datagram := session.Datagram

    // Execute the client command handler
    handler := clientCommandHandlers[datagram.Command]
    if handler == nil {
        log.Printf("Unknown client command: %d\n", datagram.Command)
        return
    }

    // Log the command being handled
    log.Printf("Running client command handler for: %s\n", commands.GetCommandName(datagram.Command))

    // Call the handler
    handler(session)
}

// dispatchServerHandler processes a server session by executing the appropriate handler.
func dispatchServerHandler(datagram *datagram_util.Datagram) {

    // Execute the server command handler
    handler := serverCommandHandlers[datagram.Command&0x7F] // Mask out the MSB
    if handler == nil {
        log.Printf("Unknown server command: %d\n", datagram.Command)
        return
    }

    // Log the command being handled
    log.Printf("Running server command handler for: %s\n", commands.GetCommandName(datagram.Command))

    // Call the handler
    handler(datagram)
}
