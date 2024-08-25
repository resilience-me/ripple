package tests

import (
    "fmt"
    "log"
    "ripple/commands"
)

func TestStartPayment() {
    senderUsername := "testsender"
    receiverUsername := "testreceiver"
    senderServerAddress := "127.0.0.1"
    receiverServerAddress := "127.0.0.1"
    sharedSecretKey := "sharedsecretkey1234567890abcdef"
    trustlineAmount := uint32(1000)
    paymentAmount := uint32(500)
    senderCounter := uint32(1)
    receiverCounter := uint32(1)

    // Set up the peer relationship and trustlines
    if err := SetupPeersAndTrustlines(senderUsername, receiverUsername, senderServerAddress, receiverServerAddress, sharedSecretKey, trustlineAmount); err != nil {
        log.Fatalf("Failed to set up peers and trustlines: %v", err)
    }

    // Hardcoded server address using the config port
    serverAddress := fmt.Sprintf("127.0.0.1:%d", config.Port)
    log.Printf("Using server address: %s", serverAddress)

    // Initiate the payment process
    if err := initiatePayment(senderUsername, receiverUsername, senderServerAddress, receiverServerAddress, serverAddress, paymentAmount, senderCounter, receiverCounter); err != nil {
        log.Fatalf("Payment initiation failed: %v", err)
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
