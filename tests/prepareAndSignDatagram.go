package tests

import (
    "ripple/auth"
    "ripple/database"
    "ripple/types"
)

// prepareAndSignDatagram prepares a datagram based on the provided parameters and signs it.
func prepareAndSignDatagram(username, peerUsername, peerServerAddress string, command byte, counter uint32, arguments []byte) ([]byte, error) {
    dg := &types.Datagram{
        Command:           command,
        Username:          username,
        PeerUsername:      peerUsername,
        PeerServerAddress: peerServerAddress,
        Counter:           counter,
    }

    // Copy the Arguments into the Datagram
    copy(dg.Arguments[:], arguments)

    // Load the client's secret key for signing
    secretKey, err := database.LoadSecretKey(username)
    if err != nil {
        return nil, err
    }

    // Serialize the datagram excluding the signature
    data, err := types.SerializeDatagram(dg)
    if err != nil {
        return nil, err
    }

    // Generate the signature
    signature := auth.GenerateSignature(data, secretKey)

    // Copy the signature directly into the serialized datagram
    copy(data[len(data)-32:], signature)

    return data, nil
}
