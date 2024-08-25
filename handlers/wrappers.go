package handlers

import "ripple/types"

// PrepareDatagramWithDatagram calls PrepareDatagram with fields from an incoming datagram
func PrepareDatagramWithDatagram(datagram *types.Datagram) (*types.Datagram, error) {
    return PrepareDatagramWithoutCommand(datagram.Username, datagram.PeerServerAddress, datagram.PeerUsername)
}

// PrepareAndSendDatagramWithDatagram calls PrepareAndSendDatagram with with fields from an incoming datagram
func PrepareAndSendDatagramWithDatagram(datagram *types.Datagram, command byte, arguments []byte) error {
    return PrepareAndSendDatagram(command, datagram.Username, datagram.PeerServerAddress, datagram.PeerUsername, arguments)
}
