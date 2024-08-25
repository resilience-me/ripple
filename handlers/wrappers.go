package handlers

import "ripple/types"

// Method on Datagram to prepare a new datagram using fields from the Datagram instance
func (dg *types.Datagram) PrepareDatagram() (*types.Datagram, error) {
    return PrepareDatagramWithoutCommand(dg.Username, dg.PeerServerAddress, dg.PeerUsername)
}

// Method on Datagram to prepare and send a new datagram using fields from the Datagram instance
func (dg *types.Datagram) PrepareAndSendDatagram(command byte, arguments []byte) error {
    return PrepareAndSendDatagram(command, dg.Username, dg.PeerServerAddress, dg.PeerUsername, arguments)
}
