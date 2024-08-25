package tests

import (
    "encoding/binary"
    "fmt"
)

const (
    usernameLen      = 32
    serverAddressLen = 32
    headerLen        = usernameLen + serverAddressLen + 1 + 4 + 4 // Total header length
)

func parsePaymentResponse(response []byte) (string, error) {
    if len(response) < headerLen {
        return "", fmt.Errorf("unexpected response length")
    }

    // Extract username and server address
    username := string(response[:usernameLen])
    serverAddress := string(response[usernameLen : usernameLen+serverAddressLen])

    // Extract payment direction, amount, and nonce
    paymentDirection := response[usernameLen+serverAddressLen]
    amount := binary.BigEndian.Uint32(response[usernameLen+serverAddressLen+1 : usernameLen+serverAddressLen+5])
    nonce := binary.BigEndian.Uint32(response[usernameLen+serverAddressLen+5 : usernameLen+serverAddressLen+9])

    return fmt.Sprintf("Username: %s, Server Address: %s, Payment Direction: %d, Amount: %d, Nonce: %d",
        username, serverAddress, paymentDirection, amount, nonce), nil
}
