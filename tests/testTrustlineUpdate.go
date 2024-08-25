package tests

import (
    "encoding/binary"
    "fmt"
    "log"
    "ripple/config"
)

func TestTrustlineUpdate() {
    // Set the parameters for the test
    username := "testuser"
    peerUsername := "peeruser"
    peerServerAddress := "127.0.0.1"
    command := byte(0) // Assuming 0 is the command for setting a trustline
    counter := uint32(3)
    trustlineAmount := uint32(1000)
    
    // Prepare the arguments
    arguments := make([]byte, 4)
    binary.BigEndian.PutUint32(arguments[:4], trustlineAmount)

    // Prepare and sign the datagram
    data, err := prepareAndSignDatagram(username, peerUsername, peerServerAddress, command, counter, arguments)
    if err != nil {
        log.Fatalf("Failed to prepare datagram: %v", err)
    }

    // Generate the server address using config package
    serverAddress := fmt.Sprintf("127.0.0.1:%d", config.Port)

    // Send the datagram and receive the response
    response, err := sendAndReceive(serverAddress, data)
    if err != nil {
        log.Fatalf("Failed to send and receive: %v", err)
    }

    // Print the response from the server
    log.Printf("Server response: %s", string(response))

    log.Println("Test completed.")
}
