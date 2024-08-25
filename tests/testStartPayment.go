package tests

import (
    "encoding/binary"
    "fmt"
    "log"
    "ripple/commands"
    "ripple/config"
)

func TestStartPayment() {
    // Set the parameters for the test
    senderUsername := "testsender"
    receiverUsername := "testreceiver"
    senderServerAddress := "127.0.0.1"
    receiverServerAddress := "127.0.0.1"
    paymentAmount := uint32(500)
    senderCounter := uint32(1)
    receiverCounter := uint32(1)
    sharedSecretKey := "sharedsecretkey1234567890abcdef" // Shared secret key for the peers

    // Setup accounts with initial balances and secret keys
    if err := SetupAccount(senderUsername, "sendersecretkey1234567890abcdef"); err != nil {
        log.Fatalf("Failed to set up sender account: %v", err)
    }
    if err := SetupAccount(receiverUsername, "receiversecretkey1234567890abcdef"); err != nil {
        log.Fatalf("Failed to set up receiver account: %v", err)
    }

    // Set up trustlines between sender and receiver
    if err := setupPeersAndTrustlines(senderUsername, receiverUsername, senderServerAddress, receiverServerAddress, sharedSecretKey, paymentAmount); err != nil {
        log.Fatalf("Failed to set up peers and trustlines: %v", err)
    }

    // Hardcoded server address using the config port
    serverAddress := fmt.Sprintf("127.0.0.1:%d", config.Port)
    log.Printf("Using server address: %s", serverAddress)

    // Initiate payment for both sender and receiver
    if err := initiatePayment(senderUsername, receiverUsername, senderServerAddress, receiverServerAddress, serverAddress, paymentAmount, senderCounter, receiverCounter); err != nil {
        log.Fatalf("Failed to initiate payment: %v", err)
    }

    // **3. Start the Payment (ClientPayments_StartPayment)**

    // Prepare the payment initiation arguments (the amount to be paid)
    arguments := make([]byte, 4)
    binary.BigEndian.PutUint32(arguments[:4], paymentAmount)

    data, err := prepareAndSignDatagram(senderUsername, receiverUsername, receiverServerAddress, commands.ClientPayments_StartPayment, senderCounter+2, arguments)
    if err != nil {
        log.Fatalf("Failed to prepare StartPayment datagram for sender: %v", err)
    }

    response, err := sendAndReceive(serverAddress, data)
    if err != nil {
        log.Fatalf("Failed to send and receive StartPayment for sender: %v", err)
    }

    log.Printf("Sender StartPayment response: %s", string(response))

    data, err = prepareAndSignDatagram(receiverUsername, senderUsername, senderServerAddress, commands.ClientPayments_StartPayment, receiverCounter+2, arguments)
    if err != nil {
        log.Fatalf("Failed to prepare StartPayment datagram for receiver: %v", err)
    }

    response, err = sendAndReceive(serverAddress, data)
    if err != nil {
        log.Fatalf("Failed to send and receive StartPayment for receiver: %v", err)
    }

    log.Printf("Receiver StartPayment response: %s", string(response))

    log.Println("Test passed: both sender and receiver started the payment successfully.")
}
