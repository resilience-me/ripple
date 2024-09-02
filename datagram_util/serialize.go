package datagram_util

import (
  "encoding/binary"
  "ripple/types"
)

// Serialize converts a Datagram struct to a byte slice.
func (dg *Datagram) Serialize() ([]byte, error) {
    // Create the byte slice
    data := make([]byte, 389)
    data[0] = dg.Command // First byte is the Command

    // Copy Usernames and Server Address
    copy(data[1:], dg.Username)
    copy(data[33:], dg.PeerUsername)
    copy(data[65:], dg.PeerServerAddress)

    // Copy Arguments
    copy(data[97:], dg.Arguments[:])

    // Write the Counter
    binary.BigEndian.PutUint32(data[353:], dg.Counter)

    // Copy the Signature (if needed for future serialization)
    copy(data[357:], dg.Signature[:])

    return data, nil
}

// DeserializeDatagram converts a byte slice to a Datagram struct.
func DeserializeDatagram(buf []byte) *Datagram {
    // Assuming buf is already confirmed to be of the correct length
    datagram := &Datagram{
        Command:           buf[0],
        Username:          types.BytesToString(buf[1:33]),
        PeerUsername:      types.BytesToString(buf[33:65]),
        PeerServerAddress: types.BytesToString(buf[65:97]),
        Arguments:         [256]byte{},
        Counter:           binary.BigEndian.Uint32(buf[353:357]),
        Signature:         [32]byte{},
    }

    // Copy data into fixed-size arrays for Arguments and Signature
    copy(datagram.Arguments[:], buf[97:353])
    copy(datagram.Signature[:], buf[357:389])

    return datagram
}
