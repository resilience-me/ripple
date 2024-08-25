package tests

import (
    "encoding/binary"
    "fmt"
    "log"
    "ripple/commands"
    "ripple/config"
    "ripple/database"
)

func TestNewPayments() {
    // Set the parameters for the test
    senderUsername := "testsender"
    receiverUsername := "testreceiver"
    senderServerAddress := "127.0.0.1"
    receiverServerAddress := "127.0.0.1"
    paymentAmount := uint32(500)
    senderCounter := uint32(1)
    receiverCounter := uint32(1)

    // Setup accounts with initial balances
    if err := SetupAccount(senderUsername, receiverUsername, senderServerAddress, "sendersecretkey1234567890abcdef"); err != nil {
        log.Fatalf("Failed to set up sender account: %v", err)
    }
    if err := SetupAccount(receiverUsername, senderUsername, receiverServerAddress, "receiversecretkey1234567890abcdef"); err != nil {
        log.Fatalf("Failed to set up receiver account: %v", err)
    }

    // Prepare the arguments (the amount to be paid)
    arguments := make([]byte, 4)
    binary.BigEndian.PutUint32(arguments[:4], paymentAmount)

    // Hardcoded server address using the config port
    serverAddress := fmt.Sprintf("127.0.0.1:%d", config.Port)

    // **1. Sender Initiates Payment**

    // Prepare and sign the payment datagram for the sender
    data, err := prepareAndSignDatagram(senderUsername, receiverUsername, receiverServerAddress, commands.ClientPayments_NewPaymentOut, senderCounter, arguments)
    if err != nil {
        log.Fatalf("Failed to prepare payment datagram for sender: %v", err)
    }

    // Send the datagram and receive the response
    response, err := sendAndReceive(serverAddress, data)
    if err != nil {
        log.Fatalf("Failed to send and receive for sender: %v", err)
    }

    log.Printf("Sender response: %s", string(response))

    // **2. Receiver Initiates Payment**

    // Prepare and sign the payment datagram for the receiver
    data, err = prepareAndSignDatagram(receiverUsername, senderUsername, senderServerAddress, commands.ClientPayments_NewPaymentIn, receiverCounter, arguments)
    if err != nil {
        log.Fatalf("Failed to prepare payment datagram for receiver: %v", err)
    }

    // Send the datagram and receive the response
    response, err = sendAndReceive(serverAddress, data)
    if err != nil {
        log.Fatalf("Failed to send and receive for receiver: %v", err)
    }

    log.Printf("Receiver response: %s", string(response))

    // **3. Verify Payment Registration Using GetPayment**

    // Prepare and sign the GetPayment datagram for the sender
    data, err = prepareAndSignDatagram(senderUsername, receiverUsername, receiverServerAddress, commands.ClientPayments_GetPaymentOut, senderCounter+1, arguments)
    if err != nil {
        log.Fatalf("Failed to prepare GetPayment datagram for sender: %v", err)
    }

    // Send the GetPayment datagram and receive the response
    response, err = sendAndReceive(serverAddress, data)
    if err != nil {
        log.Fatalf("Failed to send and receive GetPayment for sender: %v", err)
    }

    log.Printf("GetPayment response for sender: %s", string(response))

    // Prepare and sign the GetPayment datagram for the receiver
    data, err = prepareAndSignDatagram(receiverUsername, senderUsername, senderServerAddress, commands.ClientPayments_GetPaymentIn, receiverCounter+1, arguments)
    if err != nil {
        log.Fatalf("Failed to prepare GetPayment datagram for receiver: %v", err)
    }

    // Send the GetPayment datagram and receive the response
    response, err = sendAndReceive(serverAddress, data)
    if err != nil {
        log.Fatalf("Failed to send and receive GetPayment for receiver: %v", err)
    }

    log.Printf("GetPayment response for receiver: %s", string(response))

    // The test passes if both the sender and receiver successfully initiate the payment
    log.Println("Test passed: both sender and receiver initiated the payment successfully.")
}
