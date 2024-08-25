package tests

import (
    "ripple/auth"
    "ripple/database"
    "ripple/types"
)

// prepareAndSignDatagram prepares a datagram based on the test case and signs it.
func prepareAndSignDatagram(tc *TestCase) ([]byte, error) {
    dg := &types.Datagram{
        Command:           tc.Command,
        Username:          tc.Username,
        PeerUsername:      tc.PeerUsername,
        PeerServerAddress: tc.PeerServerAddress,
        Counter:           tc.Counter,
    }

    // Copy the Arguments into the Datagram
    copy(dg.Arguments[:], tc.Arguments)

    // Load the client's secret key for signing
    secretKey, err := database.LoadSecretKey(tc.Username)
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
